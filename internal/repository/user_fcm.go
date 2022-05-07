package repository

import (
	"fmt"

	"getcare-notification/internal/model"
	"getcare-notification/utils"

	"gorm.io/gorm"
)

type UserFcmRepository interface {
	List(page, pageSize int, order string, where *utils.Where) ([]*model.UserFcm, int64, error)
	Get(argId int32) (*model.UserFcm, error)
	GetByUserId(userID string) ([]*model.UserFcm, error)
	GetBySystemId(systemId string) ([]string, error)
	GetBySystemIdNotUserId(systemId string, userId string) ([]string, error)
	Create(userFcmAdd *model.UserFcmAdd) (result *model.UserFcm, err error)
	CreateByUser(userFcmAdd *model.UserFcmAdd) (*model.UserFcm, error)
	CreateByDevice(userFcmAdd *model.UserFcmAdd) (*model.UserFcm, error)
	Update(argId int32, updated *model.UserFcm) (*model.UserFcm, error)
	Delete(argId int32) error
	Begin() *gorm.DB
}

type userFcmRepository struct {
	Table *gorm.DB
}

func NewUserFcmRepository(table *gorm.DB) UserFcmRepository {
	return &userFcmRepository{
		Table: table,
	}
}

func (uf *userFcmRepository) List(page, pageSize int, order string, where *utils.Where) (results []*model.UserFcm, total int64, err error) {
	pipe := uf.Table.Model(&model.UserFcm{})
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

func (uf *userFcmRepository) Get(argId int32) (record *model.UserFcm, err error) {
	record = &model.UserFcm{}
	if err = uf.Table.First(record, argId).Error; err != nil {
		return record, err
	}
	return record, nil
}

func (uf *userFcmRepository) CreateByUser(userFcmAdd *model.UserFcmAdd) (result *model.UserFcm, err error) {
	userFcm := &model.UserFcm{}
	where := fmt.Sprintf("user_id='%s' AND device_id='%s'", userFcmAdd.UserID, userFcmAdd.DeviceId)
	if err := uf.Table.Where(where).First(userFcm).Error; err != nil {
		userFcm = &model.UserFcm{
			UserID:   userFcmAdd.UserID,
			Token:    userFcmAdd.Token,
			DeviceId: userFcmAdd.DeviceId,
		}
		if err = uf.Table.Model(&model.UserFcm{}).Create(userFcm).Error; err != nil {
			return nil, err
		}
		return userFcm, nil
	}
	if userFcm.Token == userFcmAdd.Token {
		return userFcm, nil
	}
	if err := uf.Table.Model(&model.UserFcm{}).Where(userFcm.ID).Update("token", userFcmAdd.Token).Error; err != nil {
		return nil, err
	}
	return userFcm, nil
}

func (uf *userFcmRepository) CreateByDevice(userFcmAdd *model.UserFcmAdd) (result *model.UserFcm, err error) {
	userFcm := &model.UserFcm{}
	where := fmt.Sprintf("device_id='%s'", userFcmAdd.DeviceId)
	if err := uf.Table.Where(where).First(userFcm).Error; err != nil {
		userFcm = &model.UserFcm{
			Token:    userFcmAdd.Token,
			DeviceId: userFcmAdd.DeviceId,
		}
		if err = uf.Table.Model(&model.UserFcm{}).Create(userFcm).Error; err != nil {
			return nil, err
		}
	}
	if userFcm.Token == userFcmAdd.Token {
		return userFcm, nil
	}
	if err := uf.Table.Model(&model.UserFcm{}).Where(userFcm.ID).Update("token", userFcmAdd.Token).Error; err != nil {
		return nil, err
	}
	return userFcm, nil
}

func (uf *userFcmRepository) Create(userFcmAdd *model.UserFcmAdd) (result *model.UserFcm, err error) {
	userFcm := &model.UserFcm{
		UserID:   userFcmAdd.UserID,
		Token:    userFcmAdd.Token,
		DeviceId: userFcmAdd.DeviceId,
	}
	if err = uf.Table.Model(&model.UserFcm{}).Create(userFcm).Error; err != nil {
		return nil, err
	}
	return userFcm, nil
}

func (uf *userFcmRepository) Update(argId int32, updated *model.UserFcm) (*model.UserFcm, error) {
	result := &model.UserFcm{}
	if err := uf.Table.Model(result).Updates(updated).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (uf *userFcmRepository) Delete(argId int32) error {
	return uf.Table.Where(argId).Delete(&model.RequestOtp{}).Error
}

func (uf *userFcmRepository) GetByUserId(userID string) ([]*model.UserFcm, error) {
	records := []*model.UserFcm{}
	where := fmt.Sprintf("user_id='%s'", userID)
	if err := uf.Table.Where(where).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (uf *userFcmRepository) GetBySystemId(systemId string) ([]string, error) {
	records := []string{}
	selects := "user_fcm.token"
	from := "system_user LEFT JOIN user_fcm ON system_user.user_id COLLATE utf8mb4_unicode_ci = user_fcm.user_id"
	where := fmt.Sprintf("system_user.system_id='%s'", systemId)
	if err := uf.Table.Table(from).Select(selects).Where(where).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (uf *userFcmRepository) GetBySystemIdNotUserId(systemId string, userId string) ([]string, error) {
	records := []string{}
	selects := "user_fcm.token"
	from := "system_user LEFT JOIN user_fcm ON system_user.user_id COLLATE utf8mb4_unicode_ci = user_fcm.user_id"
	where := fmt.Sprintf("system_user.system_id='%s' AND system_user.user_id<>'%s'", systemId, userId)
	if err := uf.Table.Table(from).Select(selects).Where(where).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (uf *userFcmRepository) Begin() *gorm.DB {
	return uf.Table.Begin()
}
