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

type VerifyOtpRepository interface {
	List(page, pageSize int, order string, where map[string]interface{}) (results []model.VerifyOtp, totalRows int64, err error)
	ListSort(sort, where string) (results []model.VerifyOtp, err error)
	Get(argId int32) (record *model.VerifyOtp, err error)
	Create(record *model.VerifyOtp) (result *model.VerifyOtp, RowsAffected int64, err error)
	Update(argId int32, updated *model.VerifyOtp) (result *model.VerifyOtp, RowsAffected int64, err error)
	Delete(argId int32) (rowsAffected int64, err error)
	DeleteWhere(where string) (err error)
}

type verifyOtpRepository struct {
	Table *gorm.DB
}

func NewVerifyOtpRepository(table *gorm.DB) VerifyOtpRepository {
	return &verifyOtpRepository{
		Table: table,
	}
}

func (vo *verifyOtpRepository) List(page, pageSize int, order string, where map[string]interface{}) (results []model.VerifyOtp, totalRows int64, err error) {

	resultOrm := vo.Table.Model(&model.VerifyOtp{})

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

func (vo *verifyOtpRepository) ListSort(sort, where string) (results []model.VerifyOtp, err error) {
	table := vo.Table.Where(where)
	if sort != "" {
		table = table.Order(sort)
	}
	if err = table.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, err
	}

	return results, nil
}

func (vo *verifyOtpRepository) Get(argId int32) (record *model.VerifyOtp, err error) {
	record = &model.VerifyOtp{}
	if err = vo.Table.First(record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

func (vo *verifyOtpRepository) Create(record *model.VerifyOtp) (result *model.VerifyOtp, RowsAffected int64, err error) {
	db := vo.Table.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

func (vo *verifyOtpRepository) Update(argId int32, updated *model.VerifyOtp) (result *model.VerifyOtp, RowsAffected int64, err error) {

	result = &model.VerifyOtp{}
	db := vo.Table.First(result, argId)
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

func (vo *verifyOtpRepository) Delete(argId int32) (rowsAffected int64, err error) {

	record := &model.VerifyOtp{}
	db := vo.Table.First(record, argId)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}

func (vo *verifyOtpRepository) DeleteWhere(where string) (err error) {
	verifyOTP := &model.VerifyOtp{}
	return vo.Table.Delete(verifyOTP, where).Error
}
