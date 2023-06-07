package kafka

const (
	TopicLogout    = "logout"
	TopicSendEmail = "send_email"
)

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
