package service

import (
	"getcare-notification/internal/model"
	"getcare-notification/internal/repository"
)

type VerifyOtpService interface {
	List(page, pageSize int, order string, where map[string]interface{}) ([]model.VerifyOtp, int64, error)
	Get(argId int32) (*model.VerifyOtp, error)
	Create(record *model.VerifyOtp) (*model.VerifyOtp, int64, error)
	Update(argId int32, updated *model.VerifyOtp) (*model.VerifyOtp, int64, error)
	Delete(argId int32) (int64, error)
}

type verifyOtpService struct {
	VerifyOtpRepository repository.VerifyOtpRepository
}

func NewVerifyOtpService(verifyOtpRepository repository.VerifyOtpRepository) VerifyOtpService {
	return &verifyOtpService{
		VerifyOtpRepository: verifyOtpRepository,
	}
}

func (vo *verifyOtpService) List(page, pageSize int, order string, where map[string]interface{}) ([]model.VerifyOtp, int64, error) {
	return vo.VerifyOtpRepository.List(page, pageSize, order, where)
}

func (vo *verifyOtpService) Get(argId int32) (*model.VerifyOtp, error) {
	return vo.VerifyOtpRepository.Get(argId)
}
func (vo *verifyOtpService) Create(record *model.VerifyOtp) (result *model.VerifyOtp, RowsAffected int64, err error) {
	return vo.VerifyOtpRepository.Create(record)
}

func (vo *verifyOtpService) Update(argId int32, updated *model.VerifyOtp) (*model.VerifyOtp, int64, error) {
	return vo.VerifyOtpRepository.Update(argId, updated)
}

func (vo *verifyOtpService) Delete(argId int32) (int64, error) {
	return vo.VerifyOtpRepository.Delete(argId)
}
