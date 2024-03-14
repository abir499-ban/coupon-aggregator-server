package sender

import (
	"fmt"
	"os"

	mailConfig "github.com/p-society/coupon-aggregator/mail/internal"
	"github.com/p-society/coupon-aggregator/mail/internal/models"
)

func SendMail(subject string, content string, to []string) error {
	fmt.Println("called,to  = ", to)
	senderName := os.Getenv("EMAIL_SENDER_NAME")
	senderAddress := os.Getenv("EMAIL_SENDER_ADDRESS")
	senderPassword := os.Getenv("EMAIL_SENDER_PASSWORD")
	fmt.Println("sn = ",senderName)
	sender := mailConfig.NewGmailSender(senderName, senderAddress, senderPassword)
	paramInstance := models.MailingParams{
		Subject: subject,
		Content: content,
		To:      to,
	}

	err := sender.SendEmail(paramInstance.Subject, paramInstance.Content, paramInstance.To, nil, nil, nil)

	return err
}
