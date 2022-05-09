package model

import (
	"time"
)

type LogMessage struct {
	Id         string     `json:"id,omitempty" bson:"_id,omitempty"`
	Time       *time.Time `json:"time" bson:"time"`
	Status     string     `json:"status,omitempty" bson:"status,omitempty"`
	Error      string     `json:"error,omitempty" bson:"error,omitempty"`
	IsDontSave bool       `json:"-" bson:"-"`
}
