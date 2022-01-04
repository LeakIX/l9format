package l9format

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net"
	"net/http"
	"strings"
	"time"
)

type ServicePluginInterface interface {
	// GetVersion returns plugin version
	GetVersion() (int, int, int)
	// GetProtocols returns the protocol supported by the plugin
	GetProtocols() []string
	// GetName returns the plugin unique name
	GetName() string
	// Run runs the plugin against the remote service
	Run(ctx context.Context, event L9Event, outputChannel chan L9Event,  options map[string]string)
	// Init called once when loading plugins : optional
	Init() error
	// IdentifyHttp Used to check tcpid payloads and identify the software : optional
	IdentifyHttp(event *L9Event, body string, document *goquery.Document) bool
	// IdentifyTcp Used to check tcpid payloads and identify the software : optional
	IdentifyTcp(event *L9Event, bannerBytes []byte, bannerPrintables []string) bool
	// GetReportTitle gets a descriptive title based on event for report title
	GetReportTitle(event *L9Event) string
	// GetReportDescription gets a description based on event for report description. Markdown supported
	GetReportDescription(event *L9Event) string
}

type ServicePluginBase struct {
}
// Optional implementations with defaults :

func (plugin ServicePluginBase) Init() error {
	return nil
}
func (plugin ServicePluginBase) IdentifyHttp(_ *L9Event, _ string, _ *goquery.Document) bool {
	return false
}
func (plugin ServicePluginBase) IdentifyTcp(_ *L9Event, _ []byte, _ []string) bool {
	return false
}
func (plugin ServicePluginBase) GetReportTitle(event *L9Event) string {
	return ""
}
func (plugin ServicePluginBase) GetReportDescription(event *L9Event) string {
	return ""
}

func (plugin ServicePluginBase) GetL9NetworkConnection(event *L9Event) (conn net.Conn, err error) {
	network := "tcp"
	if event.HasTransport("udp") {
		network = "udp"
	}
	addr := net.JoinHostPort(event.Ip, event.Port)
	return plugin.DialContext(nil, network, addr)
}

func (plugin ServicePluginBase) GetNetworkConnection(network string, addr string) (conn net.Conn, err error) {
	return plugin.DialContext(nil, network, addr)
}

func (plugin ServicePluginBase) DialContext(ctx context.Context, network string, addr string) (conn net.Conn, err error) {
	if ctx != nil {
		deadline, hasDeadline := ctx.Deadline()
		if hasDeadline {
			conn, err = net.DialTimeout(network, addr, deadline.Sub(time.Now()))
		} else {
			conn, err = net.DialTimeout(network, addr, 3*time.Second)
		}
	} else {
		conn, err = net.DialTimeout(network, addr, 3*time.Second)
	}
	if tcpConn, isTcp := conn.(*net.TCPConn); isTcp {
		// Will considerably lower TIME_WAIT connections and required fds,
		// since we don't plan to reconnect to the same host:port combo and need TIME_WAIT's window anyway
		// Will lead to out of sequence events if used on the same target host/port and source port starts to collide.
		// TLDR : DO NOT USE ON AN HOST THAT'S NOT DEDICATED TO SCANNING
		_ = tcpConn.SetLinger(0)

	}
	return conn, err
}

func (plugin ServicePluginBase) GetHttpClient(ctx context.Context, ip string, port string) *http.Client {
	if strings.Contains(ip, ":") && !strings.Contains(ip, "[") {
		ip = fmt.Sprintf("[%s]", ip)
	}
	return &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _ string, _ string) (net.Conn, error) {
				addr := ip + ":" + port
				return plugin.DialContext(ctx, "tcp", addr)
			},
			ResponseHeaderTimeout: 2 * time.Second,
			ExpectContinueTimeout: 2 * time.Second,
			DisableKeepAlives:     true,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

type WebPluginInterface interface {
	GetVersion() (int, int, int)
	GetRequests() []WebPluginRequest
	GetName() string
	Run(request WebPluginRequest, response WebPluginResponse, event L9Event, outputChannel chan L9Event, options map[string]string)
	// IdentifyHttp Used to check tcpid payloads and identify the software : optional
	IdentifyHttp(event *L9Event, body string, document *goquery.Document) bool
	// GetReportTitle gets a descriptive title based on event for report title
	GetReportTitle(event *L9Event) string
	// GetReportDescription gets a description based on event for report description. Markdown supported
	GetReportDescription(event *L9Event) string
}

type WebPluginRequest struct {
	Method    string
	Path      string
	Headers   map[string]string
	Body      []byte
	hashCache string
	Tags      []string
}

type WebPluginResponse struct {
	Response *http.Response
	Body     []byte
	Document *goquery.Document
}

func (resp *WebPluginResponse) GetHash() string {
	h := md5.New()
	h.Write([]byte(resp.Response.Status))
	h.Write(resp.Body)
	return string(h.Sum(nil))
}

func (request *WebPluginRequest) GetHash() string {
	if len(request.hashCache) > 0 {
		return request.hashCache
	}
	h := md5.New()
	h.Write([]byte(request.Method))
	h.Write([]byte(request.Path))
	for headerName, headerValue := range request.Headers {
		h.Write([]byte(headerName + headerValue))
	}
	h.Write(request.Body)
	request.hashCache = string(h.Sum(nil))
	return request.hashCache
}

func (request *WebPluginRequest) Equal(testRequest WebPluginRequest) bool {
	return request.GetHash() == testRequest.GetHash()
}

func (request *WebPluginRequest) EqualAny(testRequests []WebPluginRequest) bool {
	for _, testRequest := range testRequests {
		if request.GetHash() == testRequest.GetHash() {
			return true
		}
	}
	return false
}

func (request *WebPluginRequest) HasTag(tag string) bool {
	for _, eventTag := range request.Tags {
		if eventTag == tag {
			return true
		}
	}
	return false
}

func (request *WebPluginRequest) HasAnyTags(tags []string) bool {
	for _, tag := range tags {
		if request.HasTag(tag) {
			return true
		}
	}
	return false
}

func (request *WebPluginRequest) AddTags(tags []string) {
	for _, tag := range tags {
		request.AddTag(tag)
	}
}


func (request *WebPluginRequest) AddTag(tag string) {
	if !request.HasTag(tag) {
		request.Tags = append(request.Tags, tag)
	}
}