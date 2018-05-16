package mailer

// EmailData is
type EmailData struct {
	To      []Recipient `json:"to"`
	CC      []Recipient `json:"cc"`
	Subject string      `json:"subject"`
	Body    struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"body"`
	Attachment []Attachment `json:"attachment"`
}

// Recipient is
type Recipient struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// Attachment is
type Attachment struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
}
