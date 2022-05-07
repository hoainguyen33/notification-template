package controller

import (
	"getcare-notification/internal/domain"
	"getcare-notification/internal/model"
)

type LogMessageController interface {
	Write(log *model.LogMessage) error
	// CreateLogMessage(msg *model.SendBroadcast, userId string) *model.LogMessage
	// AddEventWS(msg []byte, log *model.LogMessage) (*socket.SendBroadcast, error)
}

type logMessageController struct {
	LogMessageDomain domain.LogMessageDomain
}

func NewLogMessageController(logMessageDomain domain.LogMessageDomain) LogMessageController {
	return &logMessageController{
		LogMessageDomain: logMessageDomain,
	}
}

func (lm *logMessageController) Write(log *model.LogMessage) error {
	return lm.LogMessageDomain.Write(log)
}
