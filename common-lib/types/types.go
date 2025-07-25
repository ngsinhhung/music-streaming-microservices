package types

type SendEmailOTPRegistry struct {
	Key string `json:"key" `
}

type SendEmail[T any] struct {
	Type      string `json:"type"`
	Recipient string `json:"recipient"`
	Message   T      `json:"message"`
}

type OTPWithMetadata[T any] struct {
	OTP      int `json:"otp"`
	Metadata T   `json:"metadata"`
}
