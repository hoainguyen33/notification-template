package consumer

import (
	"context"
	"getcare-notification/internal/model"
)

type DoFunc func(kmsg *model.KafkaMessage) error

type DoConsumers struct {
	DoNotification DoFunc
}

type Consumers interface {
	Run()
}

type consumers struct {
	NotificationConsumer
	doConsumers *DoConsumers
}

func NewConsumers(noti NotificationConsumer, doConsumers *DoConsumers) Consumers {
	return &consumers{
		NotificationConsumer: noti,
		doConsumers:          doConsumers,
	}
}

func (c *consumers) Run() {
	ctx, cancel := context.WithCancel(context.Background())
	// will update
	go c.NotificationConsumer.Run(ctx, cancel, c.doConsumers.DoNotification)
}
