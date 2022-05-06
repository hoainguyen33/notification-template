package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func NewKafkaConn(brokers string) (*kafka.Conn, error) {
	return kafka.DialContext(context.Background(), "tcp", brokers)
}
