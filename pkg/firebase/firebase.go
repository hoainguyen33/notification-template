package firebase

import (
	"context"
	"getcare-notification/config"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func Init(cfg *config.Config) (*messaging.Client, error) {
	opt := option.WithCredentialsFile(cfg.Firebase.KeyPath)
	config := &firebase.Config{ProjectID: cfg.Firebase.ProjectID}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Messaging(context.Background())
	if err != nil {
		return nil, err
	}
	return client, nil
}
