package firebase

import (
	"context"

	"firebase.google.com/go/messaging"
)

type Firebase struct {
	Client *messaging.Client
}

func (fbc *Firebase) SendMessage(token string, data map[string]string) error {
	message := &messaging.Message{
		Data:  data,
		Token: token,
	}
	_, err := fbc.Client.Send(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}

func (fbc *Firebase) SendMulticastMessage(tokens []string, data map[string]string) error {
	message := &messaging.MulticastMessage{
		Data:   data,
		Tokens: tokens,
	}
	_, err := fbc.Client.SendMulticast(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}
