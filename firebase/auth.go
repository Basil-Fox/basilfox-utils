package firebase

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/auth"
)

func VerifyIDToken(idToken string) (*auth.Token, error) {
	if App == nil {
		return nil, fmt.Errorf("firebase app isn't initialized yet")
	}

	ctx := context.Background()
	client, err := App.Auth(ctx)
	if err != nil {
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func SetCustomTokenClaims(firebaseUID string, claims map[string]interface{}) error {
	if App == nil {
		return fmt.Errorf("firebase app isn't initialized yet")
	}

	ctx := context.Background()
	client, err := App.Auth(ctx)
	if err != nil {
		return err
	}

	err = client.SetCustomUserClaims(ctx, firebaseUID, claims)
	if err != nil {
		return err
	}

	return nil
}
