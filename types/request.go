package types

import "encoding/json"

const (
	CONFIGURE_FILE_PATH = "~/.curly"
	CONFIGURE_FILE_NAME = "curly.json"
	VERSION             = "0.0.1"
)

var AppConfigString string

type RequestType struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    json.RawMessage   `json:"body,omitempty"`
	Verbose bool              `json:"verbose"`
}

type RequestConfigurations struct {
	TotalTime             bool `json:"total_time"`
	NameLookupTime        bool `json:"namelookup_time"`
	ConnectTime           bool `json:"connect_time"`
	AppConnectTime        bool `json:"appconnect_time"`
	PreTransferTime       bool `json:"pretransfer_time"`
	RedirectTime          bool `json:"redirect_time"`
	StartTransferTime     bool `json:"starttransfer_time"`
	RedirectCount         bool `json:"redirect_count"`
	SizeUpload            bool `json:"size_upload"`
	SizeDownload          bool `json:"size_download"`
	SpeedDownload         bool `json:"speed_download"`
	SpeedUpload           bool `json:"speed_upload"`
	HeaderSize            bool `json:"header_size"`
	RequestSize           bool `json:"request_size"`
	SslVerifyResult       bool `json:"ssl_verify_result"`
	FiletimeConnect       bool `json:"filetime_connect"`
	FiletimePretransfer   bool `json:"filetime_pretransfer"`
	FiletimeStarttransfer bool `json:"filetime_starttransfer"`
	FiletimeAppconnect    bool `json:"filetime_appconnect"`
	RemoteIp              bool `json:"remote_ip"`
	RemotePort            bool `json:"remote_port"`
	LocalIp               bool `json:"local_ip"`
	LocalPort             bool `json:"local_port"`
	RedirectUrl           bool `json:"redirect_url"`
	PrimaryIp             bool `json:"primary_ip"`
	PrimaryPort           bool `json:"primary_port"`
	CurlError             bool `json:"curl_error"`
	HttpCode              bool `json:"http_code"`
	EffectiveUrl          bool `json:"effective_url"`
	ContentType           bool `json:"content_type"`
	NumConnects           bool `json:"num_connects"`
	PrimaryIp6            bool `json:"primary_ip6"`
	LocalIp6              bool `json:"local_ip6"`
}

type CurlyConfig struct {
	AppConfig *RequestConfigurations
}
