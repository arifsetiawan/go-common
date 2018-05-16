package mailer

import (
	"encoding/json"

	nats "github.com/nats-io/go-nats"
)

// MailerNats is
type MailerNats struct {
	Connection *nats.Conn
	Subject    string
}

// NewMailerNats is
func NewMailerNats(conn *nats.Conn, subject string) *MailerNats {
	mailer := &MailerNats{}
	mailer.Connection = conn
	mailer.Subject = subject
	return mailer
}

// Send is
func (m *MailerNats) Send(data *EmailData) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = m.Connection.Publish(m.Subject, []byte(body))
	if err != nil {
		return err
	}

	err = m.Connection.Flush()
	if err != nil {
		return err
	}

	return nil
}
