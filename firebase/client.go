package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var app *firebase.App

func InitApp(credentialsPath string) error {
	opt := option.WithCredentialsFile(credentialsPath)
	_app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	app = _app
	return nil
}
