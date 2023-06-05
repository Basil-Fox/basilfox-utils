package kafka

const (
	TopicLogout = "logout"
)

type LogoutMessage struct {
	TokenID  string
	ExpireAt int64
}
