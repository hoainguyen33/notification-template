package firebase

import (
	"context"
	"getcare-notification/config"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var (
	firebaseClient *messaging.Client
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

func SendMessage(token string, data map[string]string) error {
	message := &messaging.Message{
		Data:  data,
		Token: token,
	}
	_, err := firebaseClient.Send(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}

func SendMulticastMessage(tokens []string, data map[string]string) error {
	message := &messaging.MulticastMessage{
		Data:   data,
		Tokens: tokens,
	}
	_, err := firebaseClient.SendMulticast(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}
