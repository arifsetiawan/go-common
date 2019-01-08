package twilio

import (
	"testing"
)

func TestSendSMS(t *testing.T) {

	client := NewClient("", "", "")
	err := client.SendSMS("+6282117374387", "123456")

	if err != nil {
		t.Error(err.Error())
	}
}
