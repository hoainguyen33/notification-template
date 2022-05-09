package firebase

import (
	"context"
	"fmt"

	"firebase.google.com/go/messaging"
)

type Firebase struct {
	Client *messaging.Client
}

func (fbc *Firebase) SendMessage(token string, data map[string]string) error {
	fmt.Println(token, data)
	message := &messaging.Message{
		Data:  data,
		Token: token,
	}
	d, err := fbc.Client.Send(context.Background(), message)
	if err != nil {
		fmt.Println("d", d, err)
		return err
	}
	fmt.Println("d", d, err)
	return nil
}

func (fbc *Firebase) SendMulticastMessage(tokens []string, data map[string]string) error {
	message := &messaging.MulticastMessage{
		Data:   data,
		Tokens: tokens,
	}
	d, err := fbc.Client.SendMulticast(context.Background(), message)
	if err != nil {
		return err
	}
	fmt.Println("d", d, err)
	return nil
}
