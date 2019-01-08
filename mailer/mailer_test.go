package mailer

import (
	"testing"

	nats "github.com/nats-io/go-nats"
)

func TestSendEmail(t *testing.T) {
	natsURL := "nats://192.168.99.100:4222"
	nc, err := nats.Connect(natsURL)
	if err != nil {
		t.Error(err.Error())
	}

	defer nc.Close()

	mailer := NewMailerNats(nc, "mailer.send")

	emailData := &EmailData{}
	emailData.To = append(emailData.To, Recipient{
		Name:    "Arif Setiawan",
		Address: "arif.setiawan@notmymail.com",
	})
	emailData.Subject = "This is email from mailer"
	emailData.Body.Type = "html"
	emailData.Body.Value = `
		<p>Dear Arif,</p>
		<p><br></p>
		<p>Verify to this link.</p>
		<p>Thanks</p>
		<p><br></p>
		<p>Best,</p>
	`

	err = mailer.Send(emailData)
	if err != nil {
		t.Error(err.Error())
	}
}
