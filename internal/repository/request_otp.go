package repository

import (
	"bytes"
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

type RequestOtpRepository interface {
	List(page, pageSize int, order string, where map[string]interface{}) (results []model.RequestOtp, totalRows int64, err error)
	ListSort(sort, where string) (results []model.RequestOtp, err error)
	Get(argId int32) (record *model.RequestOtp, err error)
	Create(record *model.RequestOtp) (result *model.RequestOtp, RowsAffected int64, err error)
	Update(argId int32, updated *model.RequestOtp) (result *model.RequestOtp, RowsAffected int64, err error)
	Delete(argId int32) (rowsAffected int64, err error)
	DeleteWhere(where string) (err error)
}

type requestOtpRepository struct {
	Table *gorm.DB
}

func NewRequestOtpRepository(table *gorm.DB) RequestOtpRepository {
	return &requestOtpRepository{
		Table: table,
	}
}

func (ro *requestOtpRepository) List(page, pageSize int, order string, where map[string]interface{}) (results []model.RequestOtp, totalRows int64, err error) {

	resultOrm := ro.Table.Model(&model.RequestOtp{})

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

func (ro *requestOtpRepository) ListSort(sort, where string) (results []model.RequestOtp, err error) {
	if err = ro.Table.Where(where).Order(sort).Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, err
	}

	return results, nil
}

func (ro *requestOtpRepository) Get(argId int32) (record *model.RequestOtp, err error) {
	record = &model.RequestOtp{}
	if err = ro.Table.First(record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

func (ro *requestOtpRepository) Create(record *model.RequestOtp) (result *model.RequestOtp, RowsAffected int64, err error) {
	db := ro.Table.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateRequestOtp is a function to update a single record from request_otp table in the getcare_messenger database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (ro *requestOtpRepository) Update(argId int32, updated *model.RequestOtp) (result *model.RequestOtp, RowsAffected int64, err error) {
	result = &model.RequestOtp{}
	db := ro.Table.First(result, argId)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

func (ro *requestOtpRepository) Delete(argId int32) (rowsAffected int64, err error) {

	record := &model.RequestOtp{}
	db := ro.Table.First(record, argId)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}

func (ro *requestOtpRepository) DeleteWhere(where string) (err error) {
	requestOTP := &model.RequestOtp{}
	return ro.Table.Delete(requestOTP, where).Error
}
