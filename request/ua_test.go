package request

import (
	"fmt"
	"testing"

	"github.com/ua-parser/uap-go/uaparser"
)

var userAgents = []string{
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.1 Safari/605.1.15",
	"Mozilla/5.0 (Windows NT 6.2; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36",
	"Mozilla/5.0 (X11; Linux i686 on x86_64; rv:10.0) Gecko/20100101 Firefox/10.0",
	"Mozilla/5.0 (iPad; CPU OS 5_1 like Mac OS X; en-us) AppleWebKit/534.46 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53",
	"Mozilla/5.0 (Linux; Android 7.0; SM-G892A Build/NRD90M; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/60.0.3112.107 Mobile Safari/537.36",
	"Mozilla/5.0 (iPhone9,4; U; CPU iPhone OS 10_0_1 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) Version/10.0 Mobile/14A403 Safari/602.1",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
}

/*
func TestUASurfer(t *testing.T) {
	//parser := uaparser.NewFromSaved()
	for _, v := range userAgents {
		//client := parser.Parse(v)
		ua := uasurfer.Parse(v)

		fmt.Printf("%+v\n", ua)
	}
}
*/

func TestUAParser(t *testing.T) {
	parser := uaparser.NewFromSaved()

	for _, v := range userAgents {
		ua := parser.Parse(v)

		fmt.Printf("%+v, %+v, %+v\n", ua.Device, ua.UserAgent, ua.Os)

		userAgent := &UserAgent{
			Device:  ua.Device.ToString(),
			Browser: ua.UserAgent.ToString(),
			OS:      ua.Os.ToString(),
		}

		fmt.Printf("%+v\n", userAgent)
	}
}
