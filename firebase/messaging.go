package firebase

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/messaging"
	"github.com/Basil-Fox/basilfox-utils/kafka"
)

// getMessagingClient safely retrieves the Firebase Messaging client.
func GetMessagingClient(ctx context.Context) (*messaging.Client, error) {
	app, err := GetApp()
	if err != nil {
		return nil, err
	}

	return app.Messaging(ctx)
}

func SendToTokens(ctx context.Context, msg kafka.SendPushNotificationMessage, silent bool) error {
	client, err := GetMessagingClient(ctx)
	if err != nil {
		return err
	}

	var message *messaging.MulticastMessage
	if silent {
		// Ensure the data map is initialized
		if msg.Data == nil {
			msg.Data = make(map[string]string)
		}

		// Add content-available for iOS silent notifications
		msg.Data["content-available"] = "1"

		// Prepare the silent notification (data-only)
		message = &messaging.MulticastMessage{
			Tokens: msg.Tokens,
			Data:   msg.Data,
			Android: &messaging.AndroidConfig{
				Priority: "high",
			},
			APNS: &messaging.APNSConfig{
				Headers: map[string]string{
					"apns-priority": "5",
				},
				Payload: &messaging.APNSPayload{
					Aps: &messaging.Aps{
						ContentAvailable: true,
					},
				},
			},
		}
	} else {
		// Prepare the normal notification
		message = &messaging.MulticastMessage{
			Tokens: msg.Tokens,
			Data:   msg.Data,
			Notification: &messaging.Notification{
				Title: msg.Title,
				Body:  msg.Body,
			},
			Android: &messaging.AndroidConfig{
				Priority: "high",
				Notification: &messaging.AndroidNotification{
					ChannelID: "default",
					Sound:     "default",
				},
			},
			APNS: &messaging.APNSConfig{
				Headers: map[string]string{
					"apns-priority": "10",
				},
				Payload: &messaging.APNSPayload{
					Aps: &messaging.Aps{
						Sound: "default",
					},
				},
			},
		}
	}

	// Send the message to all tokens in a single batch
	response, err := client.SendEachForMulticast(ctx, message)
	if err != nil {
		return fmt.Errorf("failed to send push notifications: %w", err)
	}

	// Check for failures
	if response.FailureCount > 0 {
		return fmt.Errorf("some notifications failed: %d/%d", response.FailureCount, len(msg.Tokens))
	}

	return nil
}
