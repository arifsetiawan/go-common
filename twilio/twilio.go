package twilio

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Client ...
type Client struct {
	AccountSID string
	AuthToken  string
	FromNumber string
	APIURL     string
}

// NewClient is
func NewClient(sid string, token string, from string) *Client {
	if len(sid) == 0 && len(token) == 0 && len(from) == 0 {
		log.Fatal("Invalid twilio SID and Token. Set SID and Token with correct value")
	}

	return &Client{
		AccountSID: sid,
		AuthToken:  token,
		FromNumber: from,
		APIURL:     "https://api.twilio.com/2010-04-01/Accounts/" + sid + "/Messages.json",
	}
}

// SendSMS ...
func (t *Client) SendSMS(to string, otpCode string) error {

	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", t.FromNumber)
	msgData.Set("Body", "Your verification code is "+otpCode)
	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", t.APIURL, &msgDataReader)
	req.SetBasicAuth(t.AccountSID, t.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		log.Println(data)
		if err == nil {
			fmt.Printf("SMS sent with sid: %s", data["sid"])
		}
		// map[body:Sent from your Twilio trial account - Your verification code is 123456 num_segments:1 uri:/2010-04-01/Accounts/AC766d209f2f30557727a1c1c9a666ae81/Messages/SM600cc79913344e62a2cd7accf5c306da.json to:+6282117374387 date_updated:Mon, 17 Sep 2018 23:18:24 +0000 date_sent:<nil> account_sid:AC766d209f2f30557727a1c1c9a666ae81 messaging_service_sid:<nil> direction:outbound-api api_version:2010-04-01 price:<nil> date_created:Mon, 17 Sep 2018 23:18:24 +0000 error_message:<nil> subresource_uris:map[media:/2010-04-01/Accounts/AC766d209f2f30557727a1c1c9a666ae81/Messages/SM600cc79913344e62a2cd7accf5c306da/Media.json] error_code:<nil> from:+18507249873 status:queued num_media:0 price_unit:USD sid:SM600cc79913344e62a2cd7accf5c306da]

	} else {
		// map[code:21211 message:The 'To' number 082117374387 is not a valid phone number. more_info:https://www.twilio.com/docs/errors/21211 status:400]
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Printf("SMS sent with sid: %s", data["sid"])
		}
		log.Println(data)
		return fmt.Errorf("Twilio status code is: %d", resp.StatusCode)
	}

	return nil
}
