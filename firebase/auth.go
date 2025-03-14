package firebase

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

// GetAuthClient safely retrieves the Firebase Auth client.
func GetAuthClient() (*auth.Client, error) {
	app, err := GetApp()
	if err != nil {
		return nil, err
	}

	return app.Auth(context.Background())
}

// VerifyIDToken verifies the given Firebase ID token and returns the decoded token.
func VerifyIDToken(idToken string) (*auth.Token, error) {
	client, err := GetAuthClient()
	if err != nil {
		return nil, err
	}

	return client.VerifyIDToken(context.Background(), idToken)
}

// SetCustomTokenClaims sets custom claims for a Firebase user.
func SetCustomTokenClaims(firebaseUID string, claims map[string]interface{}) error {
	client, err := GetAuthClient()
	if err != nil {
		return err
	}

	return client.SetCustomUserClaims(context.Background(), firebaseUID, claims)
}
