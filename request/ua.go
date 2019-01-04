package request

import (
	"net/http"

	"github.com/ua-parser/uap-go/uaparser"
)

// UserAgent ...
type UserAgent struct {
	Device  string `json:"device"`
	OS      string `json:"os"`
	Browser string `json:"browser"`
}

// GetUserAgent is
func GetUserAgent(req *http.Request, parser *uaparser.Parser) *UserAgent {
	uaString := req.Header.Get("User-Agent")
	client := parser.Parse(uaString)

	userAgent := &UserAgent{
		Device:  client.Device.ToString(),
		Browser: client.UserAgent.ToString(),
		OS:      client.Os.ToString(),
	}

	return userAgent
}
