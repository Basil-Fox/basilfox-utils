package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog"
)

const (
	TopicLogout               = "logout"
	TopicSendEmail            = "send_email"
	TopicAccountLink          = "account_link"
	TopicAccountLinkFinished  = "account_link_finished"
	TopicAccountLinkFailed    = "account_link_failed"
	TopicSendPushNotification = "send_push_notification"
)

type ConsumerMessage = sarama.ConsumerMessage

type KafkaWorker func(*ConsumerMessage, *zerolog.Logger) error

type LogoutMessage struct {
	RequestID string
	TokenID   string
	ExpireAt  int64
}

type AccountLinkMessage struct {
	RequestID string
	UserID    string
	GuestID   string
	Namespace string
}

type SendEmailMessage struct {
	RequestID string
	Recipient string
	Subject   string
	BodyType  string
	Body      string
}

type SendPushNotificationMessage struct {
	RequestID string
	Tokens    []string
	Title     string
	Body      string
	Data      map[string]string
}
