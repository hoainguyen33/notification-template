package repository

import (
	"context"
	"getcare-notification/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type LogMessageRepository interface {
	Write(log *model.LogMessage) error
}

type logMessageRepository struct {
	Database *mongo.Database
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
	collection := lm.Database.Collection("log")
	_, err := collection.InsertOne(context.Background(), log)
	return err
}
