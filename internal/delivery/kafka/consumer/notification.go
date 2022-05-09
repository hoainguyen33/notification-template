package consumer

import (
	"context"
	"getcare-notification/common"
	"getcare-notification/common/topics"
	"getcare-notification/internal/delivery/kafka"
	"sync"
)

type NotificationConsumer interface {
	Run(ctx context.Context, cancel context.CancelFunc, do DoFunc)
}

type notificationConsumer struct {
	Kafka *kafka.KafkaGroup
}

func NewNotificationConsumer(kafka *kafka.KafkaGroup) NotificationConsumer {
	return &notificationConsumer{
		Kafka: kafka,
	}
}

func (nc *notificationConsumer) consumeNotification(
	ctx context.Context,
	cancel context.CancelFunc,
	groupID string,
	topic string,
	workersNum int,
	do DoFunc,
) {
	r := nc.Kafka.NewReader(nc.Kafka.Brokers, topic, groupID)
	defer cancel()
	defer func() {
		if err := r.Close(); err != nil {
			nc.Kafka.Log.Errorf("r.Close", err)
			cancel()
		}
	}()
	w := nc.Kafka.NewWriter(topics.NotificationErrorConsume)
	defer func() {
		if err := w.Close(); err != nil {
			nc.Kafka.Log.Errorf("w.Close", err)
			cancel()
		}
	}()
	nc.Kafka.Log.Infof("Starting consumer group: %v", r.Config().GroupID)

	wg := &sync.WaitGroup{}
	for i := 0; i <= workersNum; i++ {
		wg.Add(1)
		go nc.NotificationWorker(ctx, cancel, r, w, wg, i, do)
	}
	wg.Wait()
}

// RunConsumers run kafka consumers
func (nc *notificationConsumer) Run(ctx context.Context, cancel context.CancelFunc,
	doMessage DoFunc) {
	go nc.consumeNotification(ctx, cancel, common.NotificationGroupId, topics.NotificationConsumer, common.NotificationWorker, doMessage)
}
