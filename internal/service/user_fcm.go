package service

import (
	"getcare-notification/internal/model"
	"getcare-notification/internal/repository"
	"os"
)

type UserFcmService interface {
	List(page, pageSize int, order string, where map[string]interface{}) ([]*model.UserFcm, int64, error)
	Get(argId int32) (*model.UserFcm, error)
	Create(record *model.UserFcmAdd) (result *model.UserFcm, err error)
	Update(argId int32, updated *model.UserFcm) (*model.UserFcm, int64, error)
	Delete(argId int32) (int64, error)
	Push(userID string, title string, body string, data interface{}) error
}

type userFcmService struct {
	UserFcmRepository repository.UserFcmRepository
}

func NewUserFcmService(userFcmRepository repository.UserFcmRepository) UserFcmService {
	return &userFcmService{
		UserFcmRepository: userFcmRepository,
	}
}

func (uf *userFcmService) List(page, pageSize int, order string, where map[string]interface{}) ([]*model.UserFcm, int64, error) {
	return uf.UserFcmRepository.List(page, pageSize, order, where)
}

func (uf *userFcmService) Get(argId int32) (*model.UserFcm, error) {
	return uf.UserFcmRepository.Get(argId)
}

func (uf *userFcmService) Create(record *model.UserFcmAdd) (result *model.UserFcm, err error) {
	if record.DeviceId == "" {
		return uf.UserFcmRepository.Create(record)
	}
	if record.UserID != "" {
		return uf.UserFcmRepository.CreateByUser(record)
	}
	return uf.UserFcmRepository.CreateByDevice(record)
}

func (uf *userFcmService) Update(argId int32, updated *model.UserFcm) (*model.UserFcm, int64, error) {
	return uf.UserFcmRepository.Update(argId, updated)
}

func (uf *userFcmService) Delete(argId int32) (int64, error) {
	return uf.UserFcmRepository.Delete(argId)
}

// cần hỏi

func (uf *userFcmService) Push(userID string, title string, body string, data interface{}) error {
	records, err := uf.UserFcmRepository.GetByUserId(userID)
	if err != nil {
		return err
	}
	for _, record := range records {
		fcmPayload := &model.FcmPayload{
			To:       record.Token,
			Priority: "HIGH",
			Notification: model.FcmNotificationPayload{
				Title: title,
				Body:  body,
				Image: os.Getenv("URL_CDN") + os.Getenv("PATH_LOGO"),
			},
			Data: data,
		}
		FcmSend(fcmPayload)
	}

	return nil
}

func FcmSend(data interface{}) interface{} {
	return FcmSendPostWithHeader(os.Getenv("FCM_URL_API")+"/send", data)
}
