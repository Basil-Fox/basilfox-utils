package kafka

import "github.com/Shopify/sarama"

const (
	TopicLogout               = "logout"
	TopicSendEmail            = "send_email"
	TopicAccountLink          = "account_link"
	TopicAccountLinkFinished  = "account_link_finished"
	TopicAccountLinkFailed    = "account_link_failed"
	TopicSendPushNotification = "send_push_notification"
)

type ConsumerMessage = sarama.ConsumerMessage

type KafkaWorker func(*ConsumerMessage) error

type LogoutMessage struct {
	TokenID  string
	ExpireAt int64
}

type AccountLinkMessage struct {
	UserID    string
	GuestID   string
	Namespace string
}

type SendEmailMessage struct {
	Recipient string
	Subject   string
	BodyType  string
	Body      string
}

type SendPushNotificationMessage struct {
	Tokens []string
	Title  string
	Body   string
	Data   map[string]string
}
