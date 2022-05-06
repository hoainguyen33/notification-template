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

// Migrations struct is a row record of the migrations table in the getcare-dev database
type Migrations struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id" db:"id"`
	//[ 1] version                                        varchar              null: false  primary: false  isArray: false  auto: false  col: varchar         len: 0       default: []
	Version string `gorm:"column:version;type:varchar;" json:"version" db:"version"`
}

// TableName sets the insert table name for this struct type
func (m *Migrations) TableName() string {
	return "migrations"
}

// // BeforeSave invoked before saving, return an error if field is not populated.
// func (m *Migrations) BeforeSave() error {
// 	return nil
// }

// // Prepare invoked before saving, can be used to populate fields etc.
// func (m *Migrations) Prepare() {
// }

// // Validate invoked before performing action, return an error if field is not populated.
// func (m *Migrations) Validate(action Action) error {
// 	return nil
// }
