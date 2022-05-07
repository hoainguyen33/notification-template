package repository

import (
	"getcare-notification/internal/model"
	"getcare-notification/utils"

	"gorm.io/gorm"
)

type RequestOtpRepository interface {
	List(page, pageSize int, order string, where *utils.Where) ([]*model.RequestOtp, int64, error)
	ListSort(sort, where string) ([]*model.RequestOtp, error)
	Get(argId int32) (*model.RequestOtp, error)
	Create(record *model.RequestOtp) (*model.RequestOtp, error)
	Update(argId int32, updated *model.RequestOtp) (*model.RequestOtp, error)
	Delete(argId int32) error
	DeleteWhere(where string) error
}

type requestOtpRepository struct {
	Table *gorm.DB
}

func NewRequestOtpRepository(table *gorm.DB) RequestOtpRepository {
	return &requestOtpRepository{
		Table: table,
	}
}

func (ro *requestOtpRepository) List(page, pageSize int, order string, where *utils.Where) (results []*model.RequestOtp, total int64, err error) {

	pipe := ro.Table.Model(&model.RequestOtp{})
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

func (ro *requestOtpRepository) ListSort(sort, where string) (results []*model.RequestOtp, err error) {
	if err = ro.Table.Where(where).Order(sort).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (ro *requestOtpRepository) Get(argId int32) (*model.RequestOtp, error) {
	record := &model.RequestOtp{}
	if err := ro.Table.First(record, argId).Error; err != nil {
		return record, err
	}
	return record, nil
}

func (ro *requestOtpRepository) Create(record *model.RequestOtp) (*model.RequestOtp, error) {
	if err := ro.Table.Create(record).Error; err != nil {
		return nil, err
	}
	return record, nil
}

func (ro *requestOtpRepository) Update(argId int32, updated *model.RequestOtp) (*model.RequestOtp, error) {
	result := &model.RequestOtp{}
	if err := ro.Table.Model(result).Where(argId).Updates(updated).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (ro *requestOtpRepository) Delete(argId int32) error {
	return ro.Table.Where(argId).Delete(&model.RequestOtp{}).Error
}

func (ro *requestOtpRepository) DeleteWhere(where string) error {
	return ro.Table.Delete(&model.RequestOtp{}, where).Error
}
