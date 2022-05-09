package kafka

import (
	"getcare-notification/common"
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
		MinBytes:      common.MinBytes,
		MaxBytes:      common.MaxBytes,
		QueueCapacity: common.QueueCapacity,
		Logger:        kafka.LoggerFunc(kg.Log.Debugf),
		ErrorLogger:   kafka.LoggerFunc(kg.Log.Errorf),
		MaxAttempts:   common.MaxAttempts,
		Dialer: &kafka.Dialer{
			Timeout: common.DialTimeout,
		},
		Partition:      common.PartitionNotification,
		CommitInterval: common.CommitInterval,
	})
}

func (kg *KafkaGroup) NewWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(kg.Brokers...),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireNone,
		MaxAttempts:  common.WriterMaxAttempts,
		Logger:       kafka.LoggerFunc(kg.Log.Debugf),
		ErrorLogger:  kafka.LoggerFunc(kg.Log.Errorf),
		// Compression:  compress.Snappy,
		ReadTimeout:  common.WriterReadTimeout,
		WriteTimeout: common.WriterWriteTimeout,
		BatchSize:    common.BatchSize,
		BatchTimeout: common.BatchTimeout,
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
