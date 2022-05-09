package common

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	IncomingMessages = promauto.NewCounter(prometheus.CounterOpts{
		Name: "messages_incoming_kafka_messages_total",
		Help: "The total number of incoming Kafka messages",
	})
	SuccessMessages = promauto.NewCounter(prometheus.CounterOpts{
		Name: "messages_success_incoming_kafka_messages_total",
		Help: "The total number of success incoming success Kafka messages",
	})
	ErrorMessages = promauto.NewCounter(prometheus.CounterOpts{
		Name: "messages_error_incoming_kafka_message_total",
		Help: "The total number of error incoming success Kafka messages",
	})
)
