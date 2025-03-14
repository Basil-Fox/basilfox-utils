package firebase

import (
	"context"
	"errors"
	"sync"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var (
	app *firebase.App
	mu  sync.Mutex
)

// InitWithFile initializes Firebase using a credentials file.
func InitWithFile(credentialsPath string) error {
	mu.Lock()
	defer mu.Unlock()

	opt := option.WithCredentialsFile(credentialsPath)
	_app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	app = _app
	return nil
}

// InitWithJSON initializes Firebase using credentials JSON.
func InitWithJSON(credentialsJSON []byte) error {
	mu.Lock()
	defer mu.Unlock()

	opt := option.WithCredentialsJSON(credentialsJSON)
	_app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	app = _app
	return nil
}

// GetApp returns the Firebase app instance or an error if not initialized.
func GetApp() (*firebase.App, error) {
	mu.Lock()
	defer mu.Unlock()

	if app == nil {
		return nil, errors.New("firebase app is not initialized")
	}
	return app, nil
}
