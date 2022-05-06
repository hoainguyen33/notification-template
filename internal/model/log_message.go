package model

import (
	"time"
)

type LogMessage struct {
	Id            string      `json:"id,omitempty" bson:"_id,omitempty"`
	UserIDRequied string      `json:"user_id_requied,omitempty" bson:"user_id_requied,omitempty"`
	EventRequied  string      `json:"event_requied,omitempty" bson:"event_requied,omitempty"`
	Message       interface{} `json:"message,omitempty" bson:"message,omitempty"`
	Time          *time.Time  `json:"time" bson:"time"`
	TimeSend      *time.Time  `json:"time_send" bson:"time_send"`
	Status        string      `json:"status,omitempty" bson:"status,omitempty"`
	Error         string      `json:"error,omitempty" bson:"error,omitempty"`
	LogOption     bool        `json:"-" bson:"-"`
	IsDontSave    bool        `json:"-" bson:"-"`
}

type LogMessageUpdate struct {
	ThreadID  int32 `json:"thread_id,omitempty"`
	LogOption bool  `json:"log_option,omitempty"`
}

func (s *LogMessageUpdate) TableName() string {
	return "log_message"
}
