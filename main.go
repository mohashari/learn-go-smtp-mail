package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const (
	CONFIG_SMTP_HOST     = "smtp.gmail.com"
	CONFIG_SMTP_PORT     = 587
	CONFIG_SENDER_NAME   = "xxxxxxxx"
	CONFIG_AUTH_EMAIL    = "xxxxxx"
	CONFIG_AUTH_PASSWORD = "xxxxxx"
)

func main() {
	to := []string{"xxxxxx@gmail.com", "muklis2@mailinator.com"}
	cc := []string{"ccmuklis@mailinator.com"}
	subject := "test email"
	message := "Hello this is massage from smtp mail golang"

	err := sentMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error)
	}
	log.Print("Email Send")
}

func sentMail(to, cc []string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}
	return nil

}
