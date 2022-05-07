package kafka

import (
	"context"
	"getcare-notification/config"

	"github.com/segmentio/kafka-go"
)

// connect kafka with broker[0]
func New(cfg *config.Config) (*kafka.Conn, error) {
	return kafka.DialContext(context.Background(), "tcp", cfg.Kafka.Brokers[0])
}
