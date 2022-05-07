package domain

import (
	"getcare-notification/internal/model"
	"getcare-notification/internal/repository"
	"getcare-notification/utils"
	"os"
)

type UserFcmDomain interface {
	List(page, pageSize int, order string, where *utils.Where) ([]*model.UserFcm, int64, error)
	Get(argId int32) (*model.UserFcm, error)
	Create(record *model.UserFcmAdd) (*model.UserFcm, error)
	Update(argId int32, updated *model.UserFcm) (*model.UserFcm, error)
	Delete(argId int32) error
	Push(userID string, title string, body string, data interface{}) error
}

type userFcmDomain struct {
	UserFcmRepository repository.UserFcmRepository
}

func NewUserFcmDomain(userFcmRepository repository.UserFcmRepository) UserFcmDomain {
	return &userFcmDomain{
		UserFcmRepository: userFcmRepository,
	}
}

func (uf *userFcmDomain) List(page, pageSize int, order string, where *utils.Where) ([]*model.UserFcm, int64, error) {
	return uf.UserFcmRepository.List(page, pageSize, order, where)
}

func (uf *userFcmDomain) Get(argId int32) (*model.UserFcm, error) {
	return uf.UserFcmRepository.Get(argId)
}

func (uf *userFcmDomain) Create(record *model.UserFcmAdd) (result *model.UserFcm, err error) {
	if record.DeviceId == "" {
		return uf.UserFcmRepository.Create(record)
	}
	if record.UserID != "" {
		return uf.UserFcmRepository.CreateByUser(record)
	}
	return uf.UserFcmRepository.CreateByDevice(record)
}

func (uf *userFcmDomain) Update(argId int32, updated *model.UserFcm) (*model.UserFcm, error) {
	return uf.UserFcmRepository.Update(argId, updated)
}

func (uf *userFcmDomain) Delete(argId int32) error {
	return uf.UserFcmRepository.Delete(argId)
}

// cần hỏi

func (uf *userFcmDomain) Push(userID string, title string, body string, data interface{}) error {
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
