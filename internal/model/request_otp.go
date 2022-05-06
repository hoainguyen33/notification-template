package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// RequestOtp struct is a row record of the request_otp table in the getcare_messenger database
type RequestOtp struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id" db:"id"`
	//[ 1] phone                                          varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Phone string `gorm:"column:phone;type:varchar;size:255;" json:"phone" db:"phone"`
	//[ 2] code                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Code string `gorm:"column:code;type:varchar;size:255;" json:"code" db:"code"`
	//[ 3] expired_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	ExpiredAt time.Time `gorm:"column:expired_at;type:datetime;" json:"expired_at" db:"expired_at"`
	//[ 4] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime;" json:"created_at" db:"created_at"`
	//[ 5] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at" db:"updated_at"`
}

type RequestOtpAdd struct {
	Phone string `json:"phone"`
}

type RequestOTPParam struct {
	Phone      string `json:"phone"`
	IsRegister bool   `json:"is_register"`
}

type VerifyOTPParam struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

// TableName sets the insert table name for this struct type
func (r *RequestOtp) TableName() string {
	return "request_otp"
}

// // BeforeSave invoked before saving, return an error if field is not populated.
// func (r *RequestOtp) BeforeSave() error {
// 	return nil
// }

// // Prepare invoked before saving, can be used to populate fields etc.
// func (r *RequestOtp) Prepare() {
// }

// // Validate invoked before performing action, return an error if field is not populated.
// func (r *RequestOtp) Validate(action Action) error {
// 	return nil
// }
