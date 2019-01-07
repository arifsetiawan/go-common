package request

import (
	"net/http"

	"github.com/ua-parser/uap-go/uaparser"
)

// UserAgent ...
type UserAgent struct {
	Device              string `json:"device"`
	OS                  string `json:"os"`
	OSMajorVersion      string `json:"os_major_version"`
	Browser             string `json:"browser"`
	BrowserMajorVersion string `json:"browser_major_version"`
}

// GetUserAgent is
func GetUserAgent(req *http.Request, parser *uaparser.Parser) *UserAgent {
	uaString := req.Header.Get("User-Agent")
	client := parser.Parse(uaString)

	userAgent := &UserAgent{
		Device:              client.Device.ToString(),
		Browser:             client.UserAgent.Family,
		BrowserMajorVersion: client.UserAgent.Major,
		OS:                  client.Os.Family,
		OSMajorVersion:      client.Os.Major,
	}

	return userAgent
}
