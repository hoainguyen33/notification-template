package model

import (
	"fmt"
	"time"
)

// ErrorMessage
type KafkaErrorMessage struct {
	MessageID string    `json:"messageId"`
	Offset    int64     `json:"offset"`
	Partition int       `json:"partition"`
	Topic     string    `json:"topic"`
	Error     string    `json:"error"`
	Time      time.Time `json:"time"`
}

type KafkaMessage struct {
	UserId string      `json:"user_id"`
	Key    []byte      `json:"key"`
	Msg    interface{} `json:"msg"`
}

func (kmsg *KafkaMessage) String() string {
	return fmt.Sprintf("user_id: %v, msg: %v", kmsg.UserId, kmsg.Msg)
}
