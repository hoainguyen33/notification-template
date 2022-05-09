package route

import (
	"getcare-notification/internal/delivery/kafka/consumer"
	"getcare-notification/internal/delivery/kafka/producer"
)

type Closer func()

func (r *route) RunKafka() Closer {
	consumers := consumer.NewConsumers(
		consumer.NewNotificationConsumer(r.KafkaGroup),
		&consumer.DoConsumers{
			DoNotification: r.Controller.UserFcmController.KafkaPush,
		},
	)
	go consumers.Run()
	producers := &producer.Producers{
		NotificationProducer: producer.NewNotificationProducer(r.Log, r.KafkaGroup),
	}
	return producers.Close
}
