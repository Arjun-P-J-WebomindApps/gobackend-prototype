package integrations

import (
	"fmt"
	"log"
	"net/smtp"
)

type SMTP struct {
	Host     string
	Port     string
	Password string
}

func (config SMTP) SendMail(to string, from string, subject string, content string) {

	msg := []byte(subject + "\n" + content)

	auth := smtp.PlainAuth("", from, config.Password, config.Host)

	err := smtp.SendMail(config.Host+":"+config.Port, auth, from, []string{to}, msg)

	if err != nil {
		log.Println("Error occured while sending otp", err.Error())
		return
	}

	fmt.Println("Mail send successfully")
}
