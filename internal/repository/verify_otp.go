package repository

import (
	"getcare-notification/internal/model"
	"getcare-notification/utils"

	"gorm.io/gorm"
)

type VerifyOtpRepository interface {
	List(page, pageSize int, order string, where *utils.Where) ([]*model.VerifyOtp, int64, error)
	ListSort(sort, where string) (results []*model.VerifyOtp, err error)
	Get(argId int32) (record *model.VerifyOtp, err error)
	Create(record *model.VerifyOtp) (*model.VerifyOtp, error)
	Update(argId int32, updated *model.VerifyOtp) (*model.VerifyOtp, error)
	Delete(argId int32) error
	DeleteWhere(where string) error
}

type verifyOtpRepository struct {
	Table *gorm.DB
}

func NewVerifyOtpRepository(table *gorm.DB) VerifyOtpRepository {
	return &verifyOtpRepository{
		Table: table,
	}
}

func (vo *verifyOtpRepository) List(page, pageSize int, order string, where *utils.Where) (results []*model.VerifyOtp, total int64, err error) {
	pipe := vo.Table.Model(&model.VerifyOtp{})
	if where.Ok() {
		pipe = pipe.Where(where.String())
	}
	pipe.Count(&total)
	offset := (page - 1) * pageSize
	pipe = pipe.Offset(offset).Limit(pageSize)
	if order != "" {
		pipe = pipe.Order(order)
	}
	if err = pipe.Find(&results).Error; err != nil {
		return nil, -1, err
	}

	return results, total, nil
}

func (vo *verifyOtpRepository) ListSort(sort, where string) (results []*model.VerifyOtp, err error) {
	table := vo.Table.Where(where)
	if sort != "" {
		table = table.Order(sort)
	}
	if err = table.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (vo *verifyOtpRepository) Get(argId int32) (*model.VerifyOtp, error) {
	record := &model.VerifyOtp{}
	if err := vo.Table.First(record, argId).Error; err != nil {
		return record, err
	}
	return record, nil
}

func (vo *verifyOtpRepository) Create(record *model.VerifyOtp) (result *model.VerifyOtp, err error) {
	if err = vo.Table.Create(record).Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (vo *verifyOtpRepository) Update(argId int32, updated *model.VerifyOtp) (result *model.VerifyOtp, err error) {
	result = &model.VerifyOtp{}
	if err := vo.Table.Model(result).Updates(updated).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (vo *verifyOtpRepository) Delete(argId int32) error {
	return vo.Table.Where(argId).Delete(&model.RequestOtp{}).Error
}

func (vo *verifyOtpRepository) DeleteWhere(where string) (err error) {
	verifyOTP := &model.VerifyOtp{}
	return vo.Table.Delete(verifyOTP, where).Error
}
