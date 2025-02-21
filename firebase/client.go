package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitWithFile(credentialsPath string) error {
	opt := option.WithCredentialsFile(credentialsPath)
	_app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	App = _app
	return nil
}

func InitWithJson(credentialsJSON []byte) error {
	opt := option.WithCredentialsJSON(credentialsJSON)
	_app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	App = _app
	return nil
}
