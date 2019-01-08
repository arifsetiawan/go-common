package model

// UserDevice ...
type UserDevice struct {
	UserID              string `json:"user_id"`
	UserPhone           string `json:"user_phone"`
	IP                  string `json:"ip"`
	Device              string `json:"device"`
	OS                  string `json:"os"`
	OSMajorVersion      string `json:"os_major_version"`
	Browser             string `json:"browser"`
	BrowserMajorVersion string `json:"browser_major_version"`
	Location            *GeoIP `json:"location"`
}

// GeoIP ...
type GeoIP struct {
	IP            string  `json:"ip"`
	ContinentCode string  `json:"continent_code"`
	ContinentName string  `json:"continent_name"`
	CountryCode   string  `json:"country_code"`
	CountryName   string  `json:"country_name"`
	RegionCode    string  `json:"region_code"`
	RegionName    string  `json:"region_name"`
	City          string  `json:"city"`
	Lat           float32 `json:"latitude"`
	Lon           float32 `json:"longitude"`
}

// MFACodeVerification ...
type MFACodeVerification struct {
	UserID    string `json:"user_id"`
	UserPhone string `json:"user_phone"`
	MFACode   string `json:"mfa_code"`
}
