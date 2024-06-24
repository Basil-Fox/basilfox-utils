package mail

import (
	"fmt"

	"github.com/FiberApps/common-library/logger"
	"gopkg.in/gomail.v2"
)

type MailClientConfig struct {
	Host       string
	Port       int
	Username   string
	Password   string
	Sender     string
	SenderName string
}

var mConfig *MailClientConfig

// Setup SMTP Client
func SetupMailClient(config MailClientConfig) {
	mConfig = &config
}

func Send(recipient string, subject string, bodyType string, body string) error {

	var log = logger.New()

	if mConfig == nil {
		log.Error("MAIL_CLIENT:: Client isn't initialized yet")
		return fmt.Errorf("mail client isn't initialized yet")
	}

	// Email configuration
	m := gomail.NewMessage()
	from := fmt.Sprintf("%s <%s>", mConfig.SenderName, mConfig.Sender)

	m.SetHeader("From", from)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody(bodyType, body)

	// Email sending
	d := gomail.NewDialer(mConfig.Host, mConfig.Port, mConfig.Username, mConfig.Password)
	if err := d.DialAndSend(m); err != nil {
		log.Error("MAIL_SENT:: Error while sending mail: %v", err)
		return err
	}

	log.Info("MAIL_SENT:: Sender(%s) Recipient(%s) Subject(%s)", mConfig.Sender, recipient, subject)
	return nil
}
