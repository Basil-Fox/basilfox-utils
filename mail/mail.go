package mail

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Basil-Fox/basilfox-utils/kafka"
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

var (
	mConfig *Config
	mDialer *gomail.Dialer
	once    sync.Once
)

// Setup SMTP Client (Only Once)
func SetupClient(config Config) {
	once.Do(func() {
		mConfig = &config
		mDialer = gomail.NewDialer(mConfig.Host, mConfig.Port, mConfig.Username, mConfig.Password)
	})
}

// Send an Email
func Send(msg kafka.SendEmailMessage) error {
	if mConfig == nil || mDialer == nil {
		return errors.New("mail client isn't initialized")
	}

	// Construct Email
	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", mConfig.SenderName, mConfig.Sender))
	m.SetHeader("To", msg.Recipient)
	m.SetHeader("Subject", msg.Subject)
	m.SetBody(msg.BodyType, msg.Body)

	// Send Email
	return mDialer.DialAndSend(m)
}
