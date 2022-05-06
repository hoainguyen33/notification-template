package repository

import (
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

type MigrationsRepository interface {
	List(page, pagesize int, order string) (results []model.Migrations, totalRows int64, err error)
	Get(argId int32) (record *model.Migrations, err error)
	Create(record *model.Migrations) (result *model.Migrations, RowsAffected int64, err error)
	Update(argId int32, updated *model.Migrations) (result *model.Migrations, RowsAffected int64, err error)
	Delete(argId int32) (rowsAffected int64, err error)
}

type migrationsRepository struct {
	Table *gorm.DB
}

func NewMigrationsRepository(table *gorm.DB) MigrationsRepository {
	return &migrationsRepository{
		Table: table,
	}
}

// GetAllMigrations is a function to get a slice of record(s) from migrations table in the getcare-dev database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func (m *migrationsRepository) List(page, pagesize int, order string) (results []model.Migrations, totalRows int64, err error) {

	resultOrm := m.Table.Model(&model.Migrations{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
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

// GetMigrations is a function to get a single record from the migrations table in the getcare-dev database
// error - ErrNotFound, db Find error
func (m *migrationsRepository) Get(argId int32) (record *model.Migrations, err error) {
	record = &model.Migrations{}
	if err = m.Table.First(record, argId).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddMigrations is a function to add a single record to migrations table in the getcare-dev database
// error - ErrInsertFailed, db save call failed
func (m *migrationsRepository) Create(record *model.Migrations) (result *model.Migrations, RowsAffected int64, err error) {
	db := m.Table.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateMigrations is a function to update a single record from migrations table in the getcare-dev database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func (m *migrationsRepository) Update(argId int32, updated *model.Migrations) (result *model.Migrations, RowsAffected int64, err error) {

	result = &model.Migrations{}
	db := m.Table.First(result, argId)
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

// DeleteMigrations is a function to delete a single record from migrations table in the getcare-dev database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func (m *migrationsRepository) Delete(argId int32) (rowsAffected int64, err error) {

	record := &model.Migrations{}
	db := m.Table.First(record, argId)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
