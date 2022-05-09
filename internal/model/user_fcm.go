package model

import (
	"time"
)

// UserFcm struct is a row record of the getcare_user_fcm table in the getcare-dev database
type UserFcm struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id" db:"id"`
	//[ 1] getcare_user_id                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UserID string `gorm:"column:user_id;type:varchar;" json:"-" db:"user_id"`
	//[ 2] token                                          varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Token    string `gorm:"column:token;type:varchar;" json:"token" db:"token"`
	DeviceId string `gorm:"column:device_id;type:varchar;" json:"device_id" db:"device_id"`
	//[ 3] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime;" json:"created_at" db:"created_at"`
	//[ 4] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at" db:"updated_at"`
}

type UserFcmAdd struct {
	UserID   string `json:"user_id"`
	Token    string `json:"token"`
	DeviceId string `json:"device_id"`
}

type FcmPayload struct {
	To           string                 `json:"to"`
	Priority     string                 `json:"priority"`
	Notification FcmNotificationPayload `json:"notification"`
	Data         interface{}            `json:"data"`
}

type FcmNotificationPayload struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Image string `json:"image"`
}

type PushUserFcm struct {
	UserID string      `json:"user_id"`
	Title  string      `json:"title"`
	Body   string      `json:"body"`
	Data   interface{} `json:"data"`
}

// TableName sets the insert table name for this struct type
func (g *UserFcm) TableName() string {
	return "user_fcm"
}

// BeforeSave invoked before saving, return an error if field is not populated.
// func (g *UserFcm) BeforeSave() error {
// 	return nil
// }

// // Prepare invoked before saving, can be used to populate fields etc.
// func (g *UserFcm) Prepare() {
// }

// // Validate invoked before performing action, return an error if field is not populated.
// func (g *UserFcm) Validate(action Action) error {
// 	return nil
// }
