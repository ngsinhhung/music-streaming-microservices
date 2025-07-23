package handler

import (
	"encoding/json"
	"log"
	"music-streaming-microservices/email-service/global"
	"music-streaming-microservices/email-service/internal/consumer"
	"music-streaming-microservices/email-service/internal/repository"
	"music-streaming-microservices/email-service/internal/utils"
	"music-streaming-microservices/email-service/pkg/types"
	"net/smtp"
)

type IHandler interface {
	SendEmailBySTMP(recipient string, content interface{}) error
	EmailHandler()
}

type handler struct {
	repository.IRepository
}

func (h *handler) SendEmailBySTMP(recipient string, content interface{}) error {
	auth := smtp.PlainAuth("", global.Configs.SMTP.From, global.Configs.SMTP.Password, global.Configs.SMTP.Host)
	err := smtp.SendMail(global.Configs.SMTP.Host+":"+global.Configs.SMTP.Port, auth, global.Configs.SMTP.From, []string{recipient}, []byte(content.(string)))
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) EmailHandler() {
	msgsCh := consumer.Consumer()
	for {
		msgs := <-msgsCh
		for _, msg := range msgs {
			var payload types.SendEmail
			if err := json.Unmarshal(msg.Data, &payload); err != nil {
				log.Printf("Failed to unmarshal message: ", err)
				continue
			}

			msg.Ack()

			switch payload.Type {
			case "verify_otp":
				go func() {
					var message types.SendEmailOTPRegistry
					messageBytes, _ := json.Marshal(payload.Message)
					if err := json.Unmarshal(messageBytes, &message); err != nil {
						log.Printf("Failed to decode message data: ", err)
						return
					}

					otp := h.IRepository.GetOTP(message.Key)
					if otp == "" {
						log.Printf("No OTP found for key: %s", message.Key)
						return
					}

					content := utils.BuildContentEmailOTPRegistry([]string{payload.Recipient}, global.Configs.SMTP.From, otp)
					msg := utils.BuildMessageForEmail(content)
					if err := h.SendEmailBySTMP(payload.Recipient, msg); err != nil {
						log.Printf("Failed to send email to %s: %v", payload.Recipient, err)
					} else {
						log.Printf("Email sent successfully to %s", payload.Recipient)
					}
					return
				}()

			default:
				log.Printf("Unknown email type: %s", payload.Type)
			}
		}
	}
}

func NewHandler(repository repository.IRepository) IHandler {
	return &handler{
		IRepository: repository,
	}
}
