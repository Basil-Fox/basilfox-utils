package firebase

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/auth"
)

func VerifyIDToken(idToken string) (*auth.Token, error) {
	if app == nil {
		return nil, fmt.Errorf("firebase app isn't initialized yet")
	}

	ctx := context.Background()
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}
