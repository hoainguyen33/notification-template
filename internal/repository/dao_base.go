package repository

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// BuildInfo is used to define the application build info, and inject values into via the build process.
type BuildInfo struct {

	// BuildDate date string of when build was performed filled in by -X compile flag
	BuildDate string

	// LatestCommit date string of when build was performed filled in by -X compile flag
	LatestCommit string

	// BuildNumber date string of when build was performed filled in by -X compile flag
	BuildNumber string

	// BuiltOnIP date string of when build was performed filled in by -X compile flag
	BuiltOnIP string

	// BuiltOnOs date string of when build was performed filled in by -X compile flag
	BuiltOnOs string

	// RuntimeVer date string of when build was performed filled in by -X compile flag
	RuntimeVer string
}

type LogSql func(ctx context.Context, sql string)

type ResultCount struct {
	Count int `gorm:"column:Count;type:int;"`
}

type ResultSum struct {
	Sum float64 `gorm:"column:Sum;type:decimal;"`
}

var (
	// ErrNotFound error when record not found
	ErrNotFound = fmt.Errorf("Record not found")

	// ErrUnableToMarshalJSON error when json payload corrupt
	ErrUnableToMarshalJSON = fmt.Errorf("json payload corrupt")

	// ErrUpdateFailed error when update fails
	ErrUpdateFailed = fmt.Errorf("There was an error occurs in the process of updating the data. Please try again later!")

	// ErrInsertFailed error when insert fails
	ErrInsertFailed = fmt.Errorf("There was an error occurred in the process of adding new ones. Please try again later!")

	// ErrDeleteFailed error when delete fails
	ErrDeleteFailed = fmt.Errorf("There was an error occurs during deleting data. Please try again later!")

	// ErrBadParams error when bad params passed in
	ErrBadParams = fmt.Errorf("bad params error")

	ErrDuplicate = fmt.Errorf("New added data already exists")

	// DB reference to database
	// DB            *gorm.DB
	// DbLogMessager *mongo.Database
	// DbLog         *gorm.DB
	// DbMessenger   *gorm.DB

	// AppBuildInfo reference to build info
	AppBuildInfo *BuildInfo

	// Logger function that will be invoked before executing sql
	Logger LogSql
)

// Copy a src struct into a destination struct
func Copy(dst interface{}, src interface{}) error {
	dstV := reflect.Indirect(reflect.ValueOf(dst))
	srcV := reflect.Indirect(reflect.ValueOf(src))

	if !dstV.CanAddr() {
		return errors.New("copy to value is unaddressable")
	}

	if srcV.Type() != dstV.Type() {
		return errors.New("different types can be copied")
	}

	for i := 0; i < dstV.NumField(); i++ {
		f := srcV.Field(i)
		if !strings.Contains(f.Type().String(), "model.Getcare") && !isZeroOfUnderlyingType(f.Interface()) {
			dstV.Field(i).Set(f)
		}
	}

	return nil
}

func isZeroOfUnderlyingType(x interface{}) bool {
	return x == nil || reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func CompareTimeDate(time1 time.Time, time2 time.Time) int {
	if time1.Year() > time2.Year() {
		return 1
	}
	if time1.Year() < time2.Year() {
		return -1
	}
	if time1.YearDay() > time2.YearDay() {
		return 1
	}
	if time1.YearDay() < time2.YearDay() {
		return -1
	}

	return 0
}
