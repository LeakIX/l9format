package l9format

import "time"

type L9Event struct {
	EventType   string `json:"event_type"`
	EventSource      string `json:"event_source"`
	EventPipeline    []string `json:"event_pipeline"`
	Ip          string `json:"ip"`
	Port 		string `json:"port"`
	Transports   []string `json:"transport"`
	Protocol    string `json:"protocol"`
	Http        L9HttpEvent `json:"http"`
	Summary     string `json:"summary"`
	Time		time.Time `json:"time"`
	SSL         L9SSLEvent `json:"ssl"`
	Service     L9ServiceEvent `json:"service"`
}

type L9HttpEvent struct {
	Root   string `json:"root"`
	Url string `json:"url"`
	Status int `json:"status"`
	Length int64 `json:"length"`
	Headers map[string]string `json:"header"`
}

type L9ServiceEvent struct {
	Credentials ServiceCredentials `json:"credentials"`
	Software Software `json:"software"`
}

type L9LeakEvent struct {
	Severity string `json:"severity"`
	Dataset DatasetSummary `json:"dataset"`
}

type L9SSLEvent struct {
	Enabled bool `json:"enabled"`
	JARM string  `json:"jarm"`
	CypherSuite string `json:"cypher_suite"`
	Version string `json:"version"`
	Certificate Certificate `json:"certificate"`
}

type DatasetSummary struct {
	Rows int64 `json:"rows"`
	Files int64 `json:"files"`
	Size int64  `json:"size"`
}

type Software struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	OperatingSystem string            `json:"os"`
	Modules         []SoftwareModule `json:"modules"`
	Fingerprint     string            `json:"fingerprint"`
}

type SoftwareModule struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Fingerprint string `json:"fingerprint"`
}
type ServiceCredentials struct {
	NoAuth   bool   `json:"noauth"`
	Username string `json:"username"`
	Password string `json:"password"`
	Key      string `json:"key"`
	Raw      []byte `json:"raw"`
}

type Certificate struct {
	CommonName string     `json:"cn"`
	Domains    []string   `json:"domain"`
	Fingerprint string    `json:"fingerprint"`
	KeyAlgo	string  `json:"key_algo"`
	KeySize	int  `json:"key_size"`
	IssuerName string `json:"issuer_name"`
	NotBefore time.Time `json:"not_before"`
	NotAfter time.Time `json:"not_after"`
	Valid bool `json:"valid"`
}