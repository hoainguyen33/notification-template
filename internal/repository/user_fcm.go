package repository

import (
	"bytes"
	"fmt"
	"time"

	"getcare-notification/internal/model"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

type UserFcmRepository interface {
	List(page, pageSize int, order string, where map[string]interface{}) (results []*model.UserFcm, totalRows int64, err error)
	Get(argId int32) (record *model.UserFcm, err error)
	GetByUserId(userID string) ([]*model.UserFcm, error)
	GetBySystemId(systemId string) ([]string, error)
	GetBySystemIdNotUserId(systemId string, userId string) ([]string, error)
	Create(userFcmAdd *model.UserFcmAdd) (result *model.UserFcm, err error)
	CreateByUser(userFcmAdd *model.UserFcmAdd) (result *model.UserFcm, err error)
	CreateByDevice(userFcmAdd *model.UserFcmAdd) (result *model.UserFcm, err error)
	Update(argId int32, updated *model.UserFcm) (result *model.UserFcm, RowsAffected int64, err error)
	Delete(argId int32) (rowsAffected int64, err error)
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

func (uf *userFcmRepository) List(page, pageSize int, order string, where map[string]interface{}) (results []*model.UserFcm, totalRows int64, err error) {

	resultOrm := uf.Table.Model(&model.UserFcm{})

	var sql bytes.Buffer
	sql.WriteString("1")
	for field, value := range where {
		if value != "" {
			switch field {
			default:
				sql.WriteString(" AND " + field + " LIKE '%" + value.(string) + "%'")
			}
		}
	}

	if sql.String() != "1" {
		resultOrm = resultOrm.Where(sql.String())
	}

	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pageSize
		resultOrm = resultOrm.Offset(offset).Limit(pageSize)
	} else {
		resultOrm = resultOrm.Limit(pageSize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

func (uf *userFcmRepository) Get(argId int32) (record *model.UserFcm, err error) {
	record = &model.UserFcm{}
	if err = uf.Table.First(record, argId).Error; err != nil {
		err = ErrNotFound
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

// UpdateUserFcm is a function to update a single record from user_fcm table in the getcare_messenger database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (uf *userFcmRepository) Update(argId int32, updated *model.UserFcm) (*model.UserFcm, int64, error) {
	result := &model.UserFcm{}
	db := uf.Table.First(argId)
	if err := db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err := Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err := db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

func (uf *userFcmRepository) Delete(argId int32) (rowsAffected int64, err error) {

	record := &model.UserFcm{}
	db := uf.Table.First(record, argId)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
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
