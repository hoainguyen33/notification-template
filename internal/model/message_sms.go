package model

import "time"

type MessageSMS struct {
	ID        int32      `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id" db:"id"`
	UserID    int32      `gorm:"column:user_id;type:int;" json:"user_id" db:"user_id"`
	Message   string     `gorm:"column:message;type:text;" json:"message" db:"message"`
	Phone     string     `gorm:"column:phone;type:varchar;" json:"phone,omitempty" db:"phone"`
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime;" json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at,omitempty" db:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (ms *MessageSMS) TableName() string {
	return "thread_message"
}
