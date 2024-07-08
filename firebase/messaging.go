package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/FiberApps/common-library/kafka"
	"google.golang.org/api/option"
)

var Client *firebase.App

func InitApp(credentialsPath string) error {
	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	Client = app
	return nil
}

func SendToTokens(msg kafka.SendPushNotificationMessage) error {
	if Client == nil {
		return fmt.Errorf("firebase app isn't initialized yet")
	}

	ctx := context.Background()
	client, err := Client.Messaging(ctx)
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
