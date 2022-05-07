package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	REQUEST_LOG_ID  = "log_id"
	ErrOTPIncorrect = errors.New("OTP incorrect! Please check and reenter the 6-digit OTP")
	ErrOTP60s       = errors.New("you cannot request OTP 2 times in a row no more than 60 seconds apart")
	ErrOTP1h        = errors.New("you can request and/or enter the OTP incorrectly a maximum of three times. After 3 unsuccessful OTP attempts, your OTP authentication will be blocked for 1 hour")
	ErrOTPExpired   = errors.New("OTP code is valid for 150 seconds. If you do not confirm the OTP within the lapsed time, please press “Resend” to request a new OTP")
	ErrOTP3Times    = errors.New("you have entered wrongly more than 3 times. Please request another OTP and try again")
)

type PagedResult struct {
	Result         bool        `json:"result"`
	Data           interface{} `json:"data"`
	AllowedActions []string    `json:"allowed_actions"`
	RedirectUrl    string      `json:"redirect_url,omitempty"`
}

type PagedResults struct {
	Result       bool        `json:"result"`
	Page         int         `json:"page"`
	PageSize     int         `json:"page_size"`
	TotalRecords int64       `json:"total_records"`
	Data         interface{} `json:"data"`
}

type HTTPError struct {
	Result    bool   `json:"result"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message" example:"status bad request"`
}

type Where struct {
	Values map[string]*WhereValue
}
type WhereValue struct {
	WhereType *WhereType
	Data      interface{}
}

func (w *Where) LoadData(r *http.Request) *Where {
	for k, v := range w.Values {
		str := r.FormValue(k)
		if str == "" {
			continue
		}
		switch v.WhereType.Type {
		case "":
			v.Data = str
		case "int":
			v.Data = StringToInt(str)
		case "int32":
			v.Data, _ = ParseInt32(str)
		case "string":
			v.Data = fmt.Sprintf("\"%s\"", str)
		case "time":
			v.Data, _ = StringToTime(str)
		default:
			v.Data = str
		}
	}
	return w
}

func (w *Where) Ok() bool {
	for range w.Values {
		return false
	}
	return true
}

func (w *Where) String() string {
	arr := []string{}
	for k, v := range w.Values {
		switch v.WhereType.Operator {
		case "":
			arr = append(arr, fmt.Sprintf("%s = %v", k, v.Data))
		case "LIKE":
			arr = append(arr, fmt.Sprintf("%s LIKE '%s\"%v\"%s'", k, "%", v.Data, "%"))
		default:
			arr = append(arr, fmt.Sprintf("%s %s %s", k, v.WhereType.Operator, v.Data))
		}
	}
	return strings.Join(arr, " AND ")
}

type WhereType struct {
	Type     string
	Operator string
}

func WT(t, o string) *WhereType {
	return &WhereType{
		Type:     t,
		Operator: o,
	}
}

type WhereMap map[string]*WhereType

// From map where {key: type, value: operator}
func (w *Where) FromMap(m map[string]*WhereType) *Where {
	result := &Where{}
	for k, wt := range m {
		result.Values[k] = &WhereValue{WhereType: wt}
	}
	return result
}
