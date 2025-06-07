package firebase

import (
	"context"
	"errors"
	"sync"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var (
	apps = make(map[string]*firebase.App)
	mu   sync.RWMutex
)

// InitWithFile initializes Firebase using a credentials file.
func InitWithFile(namespace string, credentialsPath string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := apps[namespace]; exists {
		return nil // Already initialized
	}

	opt := option.WithCredentialsFile(credentialsPath)
	_app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	apps[namespace] = _app
	return nil
}

// InitWithJSON initializes Firebase using credentials JSON.
func InitWithJSON(namespace string, credentialsJSON []byte) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := apps[namespace]; exists {
		return nil // Already initialized
	}

	opt := option.WithCredentialsJSON(credentialsJSON)
	_app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	apps[namespace] = _app
	return nil
}

// GetApp returns the Firebase app instance or an error if not initialized.
func GetApp(namespace string) (*firebase.App, error) {
	mu.RLock()
	defer mu.RUnlock()

	app, ok := apps[namespace]
	if !ok {
		return nil, errors.New("firebase app is not initialized")
	}
	return app, nil
}
