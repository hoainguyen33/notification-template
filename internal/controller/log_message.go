package controller

import (
	"getcare-notification/internal/model"
	"getcare-notification/internal/service"
)

type LogMessageController interface {
	Write(log *model.LogMessage) error
	// CreateLogMessage(msg *model.SendBroadcast, userId string) *model.LogMessage
	// AddEventWS(msg []byte, log *model.LogMessage) (*socket.SendBroadcast, error)
}

type logMessageController struct {
	Service service.LogMessageService
}

func NewLogMessageController(logMessageService service.LogMessageService) LogMessageController {
	return &logMessageController{
		Service: logMessageService,
	}
}

func (lm *logMessageController) Write(log *model.LogMessage) error {
	return lm.Service.Write(log)
}

// func (lm *logMessageController) CreateLogMessage(msg *model.SendBroadcast, userId string) *model.LogMessage {
// 	return lm.Service.CreateLogMessage(msg, userId)
// }

// func (lm *logMessageController) AddEventWS(msg []byte, log *model.LogMessage) (*socket.SendBroadcast, error) {
// 	data := &socket.SendBroadcast{}
// 	if err := json.Unmarshal(msg, &data); err != nil {
// 		return nil, err
// 	}
// 	log.EventRequied = data.Event
// 	return data, nil
// }

// func UpdateLogOption(ctx *gin.Context) {
// 	r := ctx.Request
// 	w := ctx.Writer

// 	logUpdate := &model.LogMessageUpdate{}
// 	if err := service.ReadJSON(r, logUpdate); err != nil {
// 		service.ReturnError(w, r, repository.ErrBadParams)
// 		return
// 	}
// 	token := ctx.GetHeader("token")
// 	userInfo, err := service.GetKeyCloakUserInfo(ctx, token)
// 	if err != nil {
// 		service.ReturnError(w, r, err)
// 		return
// 	}
// 	user, err := service.GetKeyCloakUserByID(&gin.Context{}, *userInfo.Sub)
// 	if err != nil {
// 		return
// 	}
// 	if err := service.UpdateLogOption(logUpdate, *user.ID); err != nil {
// 		service.ReturnError(w, r, err)
// 	}
// }
