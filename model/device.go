package model

// UserDevice ...
type UserDevice struct {
	IP       string `json:"ip"`
	Device   string `json:"device"`
	OS       string `json:"os"`
	Browser  string `json:"browser"`
	Location *GeoIP `json:"location"`
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
