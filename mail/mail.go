package mail

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func send(recipient string, subject string, bodyType string, body string) error {

	var (
		host     = os.Getenv("SMTP_HOST")
		port, _  = strconv.Atoi(os.Getenv("SMTP_PORT"))
		sender   = os.Getenv("SMTP_SENDER")
		username = os.Getenv("SMTP_USERNAME")
		password = os.Getenv("SMTP_PASSWORD")
	)

	// Email configuration
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody(bodyType, body)

	// Email sending
	d := gomail.NewDialer(host, port, username, password)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	log.Printf("Mail sent to %s successfully.", recipient)
	return nil
}
