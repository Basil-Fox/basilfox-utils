package firebase

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/messaging"
	"github.com/FiberApps/common-library/kafka"
)

func SendToTokens(msg kafka.SendPushNotificationMessage) error {
	if App == nil {
		return fmt.Errorf("firebase app isn't initialized yet")
	}

	ctx := context.Background()
	client, err := App.Messaging(ctx)
	if err != nil {
		return err
	}

	for _, token := range msg.Tokens {
		message := &messaging.Message{
			Notification: &messaging.Notification{
				Title: msg.Title,
				Body:  msg.Body,
			},
			Data:  msg.Data,
			Token: token,
		}

		if _, err := client.Send(ctx, message); err != nil {
			return err
		}
	}

	return nil
}
