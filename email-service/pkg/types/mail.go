package types

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress `json:"from"`
	To      []string     `json:"to"`
	Subject string       `json:"subject"`
	Body    string       `json:"body"`
}

type SendEmail struct {
	Type      string
	Recipient string
	Message   interface{}
}

type SendEmailOTPRegistry struct {
	Key string
}
