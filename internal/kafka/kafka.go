package kafka

import (
	"getcare-notification/constant"
	"getcare-notification/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/segmentio/kafka-go"
)

// KafkaGroup struct
type KafkaGroup struct {
	Brokers  []string
	GroupID  string
	Log      logger.Logger
	Validate *validator.Validate
}

// NewKafkaGroup constructor
func NewKafkaGroup(
	brokers []string,
	groupID string,
	log logger.Logger,
	validate *validator.Validate,
) *KafkaGroup {
	return &KafkaGroup{
		Brokers:  brokers,
		GroupID:  groupID,
		Log:      log,
		Validate: validate,
	}
}

func (kg *KafkaGroup) NewReader(kafkaURLs []string, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:       kafkaURLs,
		Topic:         topic,
		GroupID:       groupID,
		MinBytes:      constant.MinBytes,
		MaxBytes:      constant.MaxBytes,
		QueueCapacity: constant.QueueCapacity,
		Logger:        kafka.LoggerFunc(kg.Log.Debugf),
		ErrorLogger:   kafka.LoggerFunc(kg.Log.Errorf),
		MaxAttempts:   constant.MaxAttempts,
		Dialer: &kafka.Dialer{
			Timeout: constant.DialTimeout,
		},
		Partition:      constant.PartitionNotification,
		CommitInterval: constant.CommitInterval,
	})
}

func (kg *KafkaGroup) NewWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(kg.Brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireNone,
		MaxAttempts:  constant.WriterMaxAttempts,
		Logger:       kafka.LoggerFunc(kg.Log.Debugf),
		ErrorLogger:  kafka.LoggerFunc(kg.Log.Errorf),
		// Compression:  compress.Snappy,
		ReadTimeout:  constant.WriterReadTimeout,
		WriteTimeout: constant.WriterWriteTimeout,
		BatchSize:    constant.BatchSize,
		BatchTimeout: constant.BatchTimeout,
	}
}

func (kg *KafkaGroup) LogInfo(workerID int, m kafka.Message) {
	kg.Log.Infof(
		"WORKER: %v, message at topic/partition/offset %v/%v/%v: %s = %s\n",
		workerID,
		m.Topic,
		m.Partition,
		m.Offset,
		string(m.Key),
		string(m.Value),
	)
}
