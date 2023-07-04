package mail

import (
	"fmt"
	"os"
	"strconv"

	"github.com/FiberApps/core/logger"
	"gopkg.in/gomail.v2"
)

func Send(recipient string, subject string, bodyType string, body string) error {

	var (
		host     = os.Getenv("SMTP_HOST")
		port, _  = strconv.Atoi(os.Getenv("SMTP_PORT"))
		sender   = os.Getenv("SMTP_SENDER")
		username = os.Getenv("SMTP_USERNAME")
		password = os.Getenv("SMTP_PASSWORD")
		log      = logger.NewLogger()
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

	log.Info(fmt.Sprintf("Mail sent to %s successfully.", recipient))
	return nil
}
