package mailer

import (
	"testing"

	nats "github.com/nats-io/go-nats"
	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	natsURL := "nats://192.168.99.100:4222"
	nc, err := nats.Connect(natsURL)
	assert.Nil(t, err)
	defer nc.Close()

	mailer := NewMailer(nc, "mailer.send")

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

	err = mailer.Publish(emailData)
	assert.Nil(t, err)
}
