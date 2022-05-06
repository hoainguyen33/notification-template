package consumer

import (
	"context"
	"encoding/json"
	"getcare-notification/constant"
	"getcare-notification/internal/model"
	"sync"

	"github.com/avast/retry-go"
	"github.com/segmentio/kafka-go"
)

func (nc *notificationConsumer) NotificationWorker(
	ctx context.Context,
	cancel context.CancelFunc,
	r *kafka.Reader,
	w *kafka.Writer,
	wg *sync.WaitGroup,
	workerID int,
	do DoFunc,
) {
	defer wg.Done()
	defer cancel()

	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			nc.Kafka.Log.Errorf("FetchMessage", err)
			return
		}
		nc.Kafka.LogInfo(workerID, m)
		constant.IncomingMessages.Inc()

		var km model.KafkaMessage
		if err := json.Unmarshal(m.Value, &km); err != nil {
			constant.ErrorMessages.Inc()
			nc.Kafka.Log.Errorf("json.Unmarshal", err)
			continue
		}
		// validate struct
		// if err := nc.Kafka.Validate.StructCtx(ctx, tma); err != nil {
		// 	constant.ErrorMessages.Inc()
		// 	nc.Kafka.Log.Errorf("validate.StructCtx", err)
		// 	continue
		// }

		if err := retry.Do(func() error {
			// do
			return do(&km)
		},
			// retry options
			retry.Attempts(constant.RetryAttempts),
			retry.Delay(constant.RetryDelay),
			retry.Context(ctx),
		); err != nil {
			// error do
			constant.ErrorMessages.Inc()
			if err := nc.publishErrorNotification(ctx, w, m, err); err != nil {
				nc.Kafka.Log.Errorf("publishErrorNotification", err)
				continue
			}
			nc.Kafka.Log.Errorf("notification.Add.publishErrorNotification", err)
			continue
		}
		if err := r.CommitMessages(ctx, m); err != nil {
			constant.ErrorMessages.Inc()
			nc.Kafka.Log.Errorf("CommitMessages", err)
			continue
		}
		nc.Kafka.Log.Infof("message: %v", km)
		constant.SuccessMessages.Inc()
	}
}

func (nc *notificationConsumer) publishErrorNotification(ctx context.Context, w *kafka.Writer, m kafka.Message, err error) error {
	errMsg := &model.KafkaErrorMessage{
		Offset:    m.Offset,
		Error:     err.Error(),
		Time:      m.Time.UTC(),
		Partition: m.Partition,
		Topic:     m.Topic,
	}

	errMsgBytes, err := json.Marshal(errMsg)
	if err != nil {
		return err
	}

	return w.WriteMessages(ctx, kafka.Message{
		Value: errMsgBytes,
	})
}
