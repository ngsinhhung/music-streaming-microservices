package types

type SendEmailOTPRegistry struct {
	Key string `json:"key" `
}

type SendEmail struct {
	Type      string      `json:"type"`
	Recipient string      `json:"recipient"`
	Message   interface{} `json:"message"`
}

type OTPWithMetadata[T any] struct {
	OTP      int `json:"otp"`
	Metadata T   `json:"metadata"`
}
