package request

import (
	"net/http"
	"strings"
)

// GetAccessToken is
func GetAccessToken(req *http.Request) string {
	// Acording to https://tools.ietf.org/html/rfc6750 you can pass tokens through:
	// - Form-Encoded Body Parameter. Recomended, more likely to appear. e.g.: Authorization: Bearer mytoken123
	// - URI Query Parameter e.g. access_token=mytoken123

	auth := req.Header.Get("Authorization")
	split := strings.SplitN(auth, " ", 2)
	if len(split) != 2 || !strings.EqualFold(split[0], "bearer") {
		// Nothing in Authorization header, try access_token
		// Empty string returned if there's no such parameter
		err := req.ParseForm()
		if err != nil {
			return ""
		}
		return req.Form.Get("access_token")
	}

	return split[1]
}

// GetAPIKey is
func GetAPIKey(req *http.Request) string {
	// API-Key in header
	apiKey := req.Header.Get("API-Key")
	return apiKey
}

// GetOrigin is
func GetOrigin(req *http.Request) string {
	apiKey := req.Header.Get("Origin")
	return apiKey
}

// GetAndroidPackage is
func GetAndroidPackage(req *http.Request) string {
	apiKey := req.Header.Get("Android-Package")
	return apiKey
}

// GetIOSBundle is
func GetIOSBundle(req *http.Request) string {
	apiKey := req.Header.Get("IOS-Bundle")
	return apiKey
}
