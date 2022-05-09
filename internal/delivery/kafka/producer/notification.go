package producer

import (
	"context"
	"fmt"
	"getcare-notification/common/topics"
	"getcare-notification/pkg/logger"

	kafkaG "getcare-notification/internal/delivery/kafka"

	"github.com/segmentio/kafka-go"
)

// AccountsProducer interface
type NotificationProducer interface {
	PublishMessage(ctx context.Context, msgs ...kafka.Message) error
	Close()
	Run()
}

type notificationProducer struct {
	Log           logger.Logger
	Kafka         *kafkaG.KafkaGroup
	MessageWriter *kafka.Writer
}

// NewMessageProducer constructor
func NewNotificationProducer(log logger.Logger, kafka *kafkaG.KafkaGroup) NotificationProducer {
	return &notificationProducer{
		Log:   log,
		Kafka: kafka,
	}
}

// Run init producers writers
func (n *notificationProducer) Run() {
	// n.MessageWriter = n.Kafka.NewWriter(topics.MessageProducer)
	n.MessageWriter = n.Kafka.NewWriter(topics.NotificationProducer)
}

// Close close writers
func (n *notificationProducer) Close() {
	n.MessageWriter.Close()
}

// PublishMessage publish messages to message topic
func (n *notificationProducer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	fmt.Println(3)
	return n.MessageWriter.WriteMessages(ctx, msgs...)
}
