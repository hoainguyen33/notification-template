package repository

import (
	"context"
	"getcare-notification/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type LogMessageRepository interface {
	Write(log *model.LogMessage) error
	// UpdateOption(log *model.LogMessageUpdate, threadUserId string) error
}

type logMessageRepository struct {
	Database *mongo.Database
	// Table    *gorm.DB
}

func NewLogMessageRepository(db *mongo.Database) LogMessageRepository {
	return &logMessageRepository{
		Database: db,
	}
}

func (lm *logMessageRepository) Write(log *model.LogMessage) error {
	if log.IsDontSave {
		return nil
	}
	if !log.LogOption {
		log.Message = ""
	}
	collection := lm.Database.Collection("log_message")
	_, err := collection.InsertOne(context.Background(), log)
	return err
}

// func (lm *logMessageRepository) UpdateOption(log *model.LogMessageUpdate, threadUserId string) error {
// 	where := fmt.Sprintf("thread_id=%d AND user_id='%s'", log.ThreadID, threadUserId)
// 	if err := lm.Table.Where(where).Update("log_option", log.LogOption).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
