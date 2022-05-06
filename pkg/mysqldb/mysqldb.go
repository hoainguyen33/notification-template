package mysqldb

import (
	"getcare-notification/constant/config"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewConn Create new gormDB
func New(cfg *config.Config) (*gorm.DB, error) {
	mysql, err := gorm.Open(mysql.Open(cfg.MysqlDB.Address()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if cfg.MysqlDB.Mode != "release" {
		mysql.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,                           // Slow SQL threshold
				LogLevel:                  logger.LogLevel(cfg.MysqlDB.LogLevel), // Log level
				IgnoreRecordNotFoundError: false,                                 // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,                                  // Disable color
			},
		)
	}
	return mysql, nil
}
