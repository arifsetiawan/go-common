package request

import (
	"fmt"
	"testing"
)

/*
http://api.ipstack.com/IP?access_key=API_KEY

{
  "ip": "x.x.x.x",
  "type": "ipv4",
  "continent_code": "AS",
  "continent_name": "Asia",
  "country_code": "ID",
  "country_name": "Indonesia",
  "region_code": "JK",
  "region_name": "Jakarta",
  "city": "Jakarta",
  "zip": null,
  "latitude": -6.1744,
  "longitude": 106.8294,
  "location": {
    "geoname_id": 1642911,
    "capital": "Jakarta",
    "languages": [
      {
        "code": "id",
        "name": "Indonesian",
        "native": "Bahasa Indonesia"
      }
    ],
    "country_flag": "http://assets.ipstack.com/flags/id.svg",
    "country_flag_emoji": "🇮🇩",
    "country_flag_emoji_unicode": "U+1F1EE U+1F1E9",
    "calling_code": "62",
    "is_eu": false
  }
}

*/

func TestGeoIP(t *testing.T) {
	ip := "x.x.x.x"
	apiurl := "http://api.ipstack.com"
	key := "two"

	geo, err := GetLocation(ip, apiurl, key)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", geo)
}
