package domain

import (
	"getcare-notification/internal/model"
	"getcare-notification/internal/repository"
	"getcare-notification/utils"
)

type VerifyOtpDomain interface {
	List(page, pageSize int, order string, where *utils.Where) ([]*model.VerifyOtp, int64, error)
	Get(argId int32) (*model.VerifyOtp, error)
	Create(record *model.VerifyOtp) (*model.VerifyOtp, error)
	Update(argId int32, updated *model.VerifyOtp) (*model.VerifyOtp, error)
	Delete(argId int32) error
}

type verifyOtpDomain struct {
	VerifyOtpRepository repository.VerifyOtpRepository
}

func NewVerifyOtpDomain(verifyOtpRepository repository.VerifyOtpRepository) VerifyOtpDomain {
	return &verifyOtpDomain{
		VerifyOtpRepository: verifyOtpRepository,
	}
}

func (vo *verifyOtpDomain) List(page, pageSize int, order string, where *utils.Where) ([]*model.VerifyOtp, int64, error) {
	return vo.VerifyOtpRepository.List(page, pageSize, order, where)
}

func (vo *verifyOtpDomain) Get(argId int32) (*model.VerifyOtp, error) {
	return vo.VerifyOtpRepository.Get(argId)
}
func (vo *verifyOtpDomain) Create(record *model.VerifyOtp) (*model.VerifyOtp, error) {
	return vo.VerifyOtpRepository.Create(record)
}

func (vo *verifyOtpDomain) Update(argId int32, updated *model.VerifyOtp) (*model.VerifyOtp, error) {
	return vo.VerifyOtpRepository.Update(argId, updated)
}

func (vo *verifyOtpDomain) Delete(argId int32) error {
	return vo.VerifyOtpRepository.Delete(argId)
}
