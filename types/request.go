package types

type RequestType struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Data    *string           `json:"data,omitempty"`
	Verbose bool              `json:"verbose"`
}
