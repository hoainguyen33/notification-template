package model

import (
	"time"
)

// VerifyOtp struct is a row record of the verify_otp table in the getcare_messenger database
type VerifyOtp struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id" db:"id"`
	//[ 1] request_otp_id                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	RequestOtpID int32 `gorm:"column:request_otp_id;type:int;" json:"request_otp_id" db:"request_otp_id"`
	//[ 2] code                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Code string `gorm:"column:code;type:varchar;size:255;" json:"code" db:"code"`
	//[ 3] created_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime;" json:"created_at" db:"created_at"`
	//[ 4] updated_at                                     datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at" db:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (v *VerifyOtp) TableName() string {
	return "verify_otp"
}

// BeforeSave invoked before saving, return an error if field is not populated.
// func (v *VerifyOtp) BeforeSave() error {
// 	return nil
// }

// // Prepare invoked before saving, can be used to populate fields etc.
// func (v *VerifyOtp) Prepare() {
// }

// // Validate invoked before performing action, return an error if field is not populated.
// func (v *VerifyOtp) Validate(action Action) error {
// 	return nil
// }
