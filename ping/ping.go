package ping

import (
	"time"

	"github.com/parnurzeal/gorequest"
)

// CheckURL is
func CheckURL(url string, timeout time.Duration) (string, []error) {
	request := gorequest.New().Timeout(timeout)
	_, body, errs := request.Get(url).End()
	return body, errs
}
