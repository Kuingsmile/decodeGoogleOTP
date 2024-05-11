package utils

type Flags struct {
	Input   string
	Csv     string
	Json    string
	Text    string
	QrDir   string
	Url     string
	Version bool
	PrintQR bool
	Debug   bool
	Silent  bool
}

type OutputResult struct {
	Issuer  string `json:"issuer"`
	Name    string `json:"name"`
	Secret  string `json:"secret"`
	Type    string `json:"type"`
	Counter int64  `json:"counter"`
	URL     string `json:"url"`
}
