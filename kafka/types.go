package kafka

import "github.com/Shopify/sarama"

const (
	TopicLogout    = "logout"
	TopicSendEmail = "send_email"
)

type ConsumerMessage = sarama.ConsumerMessage

type KafkaWorker func(*ConsumerMessage) error

type LogoutMessage struct {
	TokenID  string
	ExpireAt int64
}

type SendEmailMessage struct {
	Recipient string
	Subject   string
	BodyType  string
	Body      string
}
