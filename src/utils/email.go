package utils

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"github.com/TxZer0/Go_Backend_Blog/src/dto/response"
)

func SendEmail(to string, message []byte) error {
	if message == nil {
		return errors.New(response.Msg[response.InternalError])
	}
	from := os.Getenv("EMAIL")
	password := os.Getenv("APP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	if from == "" || password == "" || smtpHost == "" || smtpPort == "" {
		return errors.New("missing SMTP configuration")
	}

	auth := smtp.PlainAuth("", from, password, smtpHost)
	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message); err != nil {
		return err
	}
	return nil
}

func VerifyAccount(email string) []byte {
	token, err := GenerateEmailToken(email, time.Now().Add(time.Hour*1).Unix())
	if err != nil {
		return nil
	}
	link := fmt.Sprintf("http://%s:%s/api/verify/%s", os.Getenv("HOST"), os.Getenv("PORT"), token)
	subject := "Subject: Verify Your Account\r\n"
	headers := "MIME-Version: 1.0\r\nContent-Type: text/plain; charset=\"utf-8\"\r\n"
	body := fmt.Sprintf("\r\nClick here to verify account: %s", link)
	message := []byte(subject + headers + body)
	return message
}

func ChangePassword(email string) []byte {
	token, err := GenerateEmailToken(email, time.Now().Add(time.Minute*1).Unix())
	if err != nil {
		return nil
	}
	link := fmt.Sprintf("http://%s:%s/api/change/%s", os.Getenv("HOST"), os.Getenv("PORT"), token)
	subject := "Subject: Verify Your Account\r\n"
	headers := "MIME-Version: 1.0\r\nContent-Type: text/plain; charset=\"utf-8\"\r\n"
	body := fmt.Sprintf("\r\nClick here to verify account: %s", link)
	message := []byte(subject + headers + body)
	return message
}
