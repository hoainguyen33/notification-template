package service

import (
	"getcare-notification/internal/model"
	"getcare-notification/internal/repository"
)

type LogMessageService interface {
	Write(log *model.LogMessage) error
	// CreateLogMessage(msg *model.SendBroadcast, userId string) *model.LogMessage
}

type logMessageService struct {
	LogMessageRepository repository.LogMessageRepository
}

func NewLogMessageService(logMessageRepository repository.LogMessageRepository) LogMessageService {
	return &logMessageService{
		LogMessageRepository: logMessageRepository,
	}
}

// func (lm *logMessageService) CreateLogMessage(msg *model.SendBroadcast, userId string) *model.LogMessage {
// 	today := time.Now()
// 	if msg == nil {
// 		msg = &model.SendBroadcast{}
// 	}
// 	return &model.LogMessage{
// 		Time:          &today,
// 		TimeSend:      msg.Date,
// 		Id:            uuid.New().String(),
// 		Status:        socket.StatusFailed,
// 		Message:       msg.Data,
// 		UserIDRequied: userId,
// 		EventRequied:  msg.Event,
// 	}
// }

func (lm *logMessageService) Write(log *model.LogMessage) error {
	return lm.LogMessageRepository.Write(log)
}

func (lm *logMessageService) UpdateOption(log *model.LogMessageUpdate, threadUserId string) error {
	return nil
}
