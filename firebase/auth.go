package firebase

import (
	"context"

	"firebase.google.com/go/v4/auth"
)

// GetAuthClient safely retrieves the Firebase Auth client.
func GetAuthClient(ctx context.Context) (*auth.Client, error) {
	app, err := GetApp()
	if err != nil {
		return nil, err
	}

	return app.Auth(ctx)
}

// VerifyIDToken verifies the given Firebase ID token and returns the decoded token.
func VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	client, err := GetAuthClient(ctx)
	if err != nil {
		return nil, err
	}

	return client.VerifyIDToken(ctx, idToken)
}

// SetCustomTokenClaims sets custom claims for a Firebase user.
func SetCustomTokenClaims(ctx context.Context, firebaseUID string, claims map[string]interface{}) error {
	client, err := GetAuthClient(ctx)
	if err != nil {
		return err
	}

	return client.SetCustomUserClaims(ctx, firebaseUID, claims)
}
