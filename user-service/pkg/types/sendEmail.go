package types

type SendEmail struct {
	Type      string      `json:"type"`
	Recipient string      `json:"recipient"`
	Message   interface{} `json:"message"`
}

type SendEmailOTPRegistry struct {
	Key string `json:"key" `
}
