package producer

import (
	"context"
	"getcare-notification/constant/topics"
	"getcare-notification/pkg/logger"

	kafkaG "getcare-notification/internal/kafka"

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
	n.MessageWriter = n.Kafka.NewWriter(topics.MessageProducer)
}

// Close close writers
func (n *notificationProducer) Close() {
	n.MessageWriter.Close()
}

// PublishMessage publish messages to message topic
func (n *notificationProducer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	return n.MessageWriter.WriteMessages(ctx, msgs...)
}
