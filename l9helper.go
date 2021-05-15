package l9format

import (
	"errors"
	"fmt"
	"hash/fnv"
	"strings"
)

func (event *L9Event) UpdateFingerprint() error {
	hasher := fnv.New128()
	n, err := hasher.Write([]byte(event.EventType))
	if err != nil || n != len(event.EventType) {
		return errors.New("event hashing error")
	}
	n, err = hasher.Write([]byte(event.EventSource))
	if err != nil || n != len(event.EventSource) {
		return errors.New("event hashing error")
	}
	n, err = hasher.Write([]byte(event.Summary))
	if err != nil || n != len(event.Summary) {
		return errors.New("event hashing error")
	}
	event.EventFingerprint = fmt.Sprintf("%x", hasher.Sum([]byte{}))
	return nil
}

func (event *L9Event) HasTag(tag string) bool {
	for _, eventTag := range event.Tags {
		if eventTag == tag {
			return true
		}
	}
	return false
}

func (event *L9Event) AddTag(tag string) {
	if !event.HasTag(tag) {
		event.Tags = append(event.Tags, tag)
	}
}

func (event *L9Event) RemoveTransport(transportCheck string) {
	transports := event.Transports
	event.Transports = []string{}
	for _, transport := range transports {
		if transport != transportCheck {
			event.Transports = append(event.Transports, transport)
		}
	}
}

func (event *L9Event) HasTransport(transport string) bool {
	for _, check := range event.Transports {
		if check == transport {
			return true
		}
	}
	return false
}

func (event *L9Event) HasSource(source string) bool {
	for _, check := range event.EventPipeline {
		if check == source {
			return true
		}
	}
	return false
}

func (event *L9Event) AddSource(source string) {
	event.EventPipeline = append(event.EventPipeline, source)
	event.EventSource = source
}

func (event *L9Event) MatchServicePlugin(plugin ServicePluginInterface) bool {
	for _, eventProtocol := range plugin.GetProtocols() {
		if eventProtocol == event.Protocol {
			return true
		}
	}
	return false
}

func (event *L9Event) Url() string {
	var host string
	var scheme string
	var path string
	host = event.Host
	if len(host) < 1 {
		host = event.Ip
		if strings.Contains(event.Ip, ":") && !strings.Contains(event.Ip, "[") {
			host = "[" + event.Ip + "]"
		}
	}
	if event.HasTransport("http") {
		if event.HasTransport("tls") {
			if event.Port != "443" {
				host += ":" + event.Port
			}
			scheme = "https"
		} else {
			if event.Port != "80" {
				host += ":" + event.Port
			}
			scheme = "http"
		}
	}
	if len(event.Http.Url) > 1 {
		path = event.Http.Url
	} else if len(event.Http.Root) > 1 {
		path = event.Http.Root
	}
	return scheme + "://" + host + path
}
