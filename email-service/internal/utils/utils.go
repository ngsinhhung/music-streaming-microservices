package utils

import (
	"fmt"
	"music-streaming-microservices/email-service/pkg/types"
	"strings"
)

func BuildMessageForEmail(mail types.Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func BuildContentEmailOTPRegistry(to []string, from string, otp string) types.Mail {
	return types.Mail{
		From: types.EmailAddress{
			Address: from,
			Name:    "Music Streaming Service",
		},
		To:      to,
		Subject: "Verify Your From",
		Body:    fmt.Sprintf(`Your OTP is <b>%s</b>. Please use this code to verify your email address.`, otp),
	}
}
