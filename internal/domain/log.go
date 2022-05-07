package domain

import (
	"getcare-notification/internal/model"
	"getcare-notification/internal/repository"
)

type LogMessageDomain interface {
	Write(log *model.LogMessage) error
}

type logMessageDomain struct {
	LogMessageRepository repository.LogMessageRepository
}

func NewLogMessageDomain(logMessageRepository repository.LogMessageRepository) LogMessageDomain {
	return &logMessageDomain{
		LogMessageRepository: logMessageRepository,
	}
}

func (lm *logMessageDomain) Write(log *model.LogMessage) error {
	return lm.LogMessageRepository.Write(log)
}
