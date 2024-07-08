package mail

import (
	"fmt"

	"github.com/FiberApps/common-library/kafka"
	"gopkg.in/gomail.v2"
)

type Config struct {
	Host       string
	Port       int
	Username   string
	Password   string
	Sender     string
	SenderName string
}

var mConfig *Config

// Setup SMTP Client
func SetupClient(config Config) {
	mConfig = &config
}

func Send(msg kafka.SendEmailMessage) error {
	if mConfig == nil {
		return fmt.Errorf("mail client isn't initialized yet")
	}

	// Email configuration
	m := gomail.NewMessage()
	from := fmt.Sprintf("%s <%s>", mConfig.SenderName, mConfig.Sender)

	m.SetHeader("From", from)
	m.SetHeader("To", msg.Recipient)
	m.SetHeader("Subject", msg.Subject)
	m.SetBody(msg.BodyType, msg.Body)

	// Email sending
	d := gomail.NewDialer(mConfig.Host, mConfig.Port, mConfig.Username, mConfig.Password)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
