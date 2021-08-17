package l9format

import "time"

const SEVERITY_CRITICAL = "critical"
const SEVERITY_HIGH = "high"
const SEVERITY_MEDIUM = "medium"
const SEVERITY_LOW = "low"
const SEVERITY_INFO = "info"

const STAGE_OPEN = "open"
const STAGE_EXPLORE = "explore"
const STAGE_EXFILTRATE = "exfiltrate"

type L9Event struct {
	EventType        string         `json:"event_type,omitempty"`
	EventSource      string         `json:"event_source,omitempty"`
	EventPipeline    []string       `json:"event_pipeline,omitempty"`
	EventFingerprint string         `json:"event_fingerprint,omitempty"`
	Ip               string         `json:"ip,omitempty"`
	Host             string         `json:"host,omitempty"`
	Reverse          string         `json:"reverse,omitempty"`
	Port             string         `json:"port,omitempty"`
	Mac              string         `json:"mac,omitempty"`
	Vendor           string         `json:"vendor,omitempty"`
	Transports       []string       `json:"transport,omitempty"`
	Protocol         string         `json:"protocol,omitempty"`
	Http             *L9HttpEvent    `json:"http,omitempty"`
	Summary          string         `json:"summary,omitempty"`
	Time             time.Time      `json:"time,omitempty"`
	SSL              *L9SSLEvent     `json:"ssl,omitempty"`
	SSH              *L9SSHEvent     `json:"ssh,omitempty"`
	Service          *L9ServiceEvent `json:"service,omitempty"`
	Leak             *L9LeakEvent    `json:"leak,omitempty"`
	Tags             []string       `json:"tags,omitempty"`
	GeoIp            *GeoLocation    `json:"geoip,omitempty"`
	Network          *Network        `json:"network,omitempty"`
}

type L9HttpEvent struct {
	Root        string            `json:"root,omitempty"`
	Url         string            `json:"url,omitempty"`
	Status      *int               `json:"status,omitempty"`
	Length      *int64             `json:"length,omitempty"`
	Headers     map[string]string `json:"header,omitempty"`
	Title       string            `json:"title,omitempty"`
	FaviconHash string            `json:"favicon_hash,omitempty"`
}

type L9ServiceEvent struct {
	Credentials *ServiceCredentials `json:"credentials,omitempty"`
	Software    *Software           `json:"software,omitempty"`
}

type L9SSHEvent struct {
	Fingerprint string `json:"fingerprint,omitempty"`
	Version     *int    `json:"version,omitempty"`
	Banner      string `json:"banner,omitempty"`
	Motd        string `json:"motd,omitempty"`
}

type L9LeakEvent struct {
	Stage    string         `json:"stage,omitempty"`
	Type     string         `json:"type,omitempty"`
	Severity string         `json:"severity,omitempty"`
	Dataset  *DatasetSummary `json:"dataset,omitempty"`
}

type L9SSLEvent struct {
	Detected    *bool        `json:"detected,omitempty"`
	Enabled     *bool        `json:"enabled,omitempty"`
	JARM        string      `json:"jarm,omitempty"`
	CypherSuite string      `json:"cypher_suite,omitempty"`
	Version     string      `json:"version,omitempty"`
	Certificate *Certificate `json:"certificate,omitempty"`
}

type DatasetSummary struct {
	Rows        *int64    `json:"rows,omitempty"`
	Files       *int64    `json:"files,omitempty"`
	Size        *int64    `json:"size,omitempty"`
	Collections *int64    `json:"collections,omitempty"`
	Infected    *bool     `json:"infected,omitempty"`
	RansomNotes []string `json:"ransom_notes,omitempty"`
}

type Software struct {
	Name            string           `json:"name,omitempty"`
	Version         string           `json:"version,omitempty"`
	OperatingSystem string           `json:"os,omitempty"`
	Modules         []SoftwareModule `json:"modules,omitempty"`
	Fingerprint     string           `json:"fingerprint,omitempty"`
}

type SoftwareModule struct {
	Name        string `json:"name,omitempty"`
	Version     string `json:"version,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
}
type ServiceCredentials struct {
	NoAuth   *bool   `json:"noauth,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Key      string `json:"key,omitempty"`
	Raw      []byte `json:"raw,omitempty"`
}

type Certificate struct {
	CommonName  string    `json:"cn,omitempty"`
	Domains     []string  `json:"domain,omitempty"`
	Fingerprint string    `json:"fingerprint,omitempty"`
	KeyAlgo     string    `json:"key_algo,omitempty"`
	KeySize     *int       `json:"key_size,omitempty"`
	IssuerName  string    `json:"issuer_name,omitempty"`
	NotBefore   time.Time `json:"not_before,omitempty"`
	NotAfter    time.Time `json:"not_after,omitempty"`
	Valid       *bool      `json:"valid,omitempty"`
}

type GeoLocation struct {
	ContinentName  string   `json:"continent_name,omitempty"`
	RegionISOCode  string   `json:"region_iso_code,omitempty"`
	CityName       string   `json:"city_name,omitempty"`
	CountryISOCode string   `json:"country_iso_code,omitempty"`
	CountryName    string   `json:"country_name,omitempty"`
	RegionName     string   `json:"region_name,omitempty"`
	GeoPoint       *GeoPoint `json:"location,omitempty"`
}

type GeoPoint struct {
	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
}

type Network struct {
	OrganisationName string `json:"organization_name,omitempty"`
	ASN              int    `json:"asn,omitempty"`
	NetworkCIDR      string `json:"network,omitempty"`
}
