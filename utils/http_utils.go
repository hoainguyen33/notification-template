package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"getcare-notification/internal/model"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

func SliceContains(slice interface{}, e interface{}) bool {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("SliceContains() given a non-slice type")
	}
	for i := 0; i < s.Len(); i++ {
		if e == s.Index(i).Interface() {
			return true
		}
	}
	return false
}

func WhereTrim(where map[string]interface{}) {
	for k, v := range where {
		where[k] = strings.TrimSpace(v.(string))
	}
}

type RequestValidatorFunc func(ctx context.Context, r *http.Request, table string, action model.Action) error

var RequestValidator RequestValidatorFunc

func ValidateRequest(ctx context.Context, r *http.Request, table string, action model.Action) error {
	if RequestValidator != nil {
		return RequestValidator(ctx, r, table, action)
	}

	return nil
}

type ContextInitializerFunc func(r *http.Request) (ctx context.Context)

var ContextInitializer ContextInitializerFunc

func InitializeContext(r *http.Request) (ctx context.Context) {
	if ContextInitializer != nil {
		ctx = ContextInitializer(r)
	} else {
		ctx = r.Context()
	}
	return ctx
}

func JsonConvertCamelCase(j json.RawMessage) json.RawMessage {
	mArray := []map[string]json.RawMessage{}
	if err := json.Unmarshal([]byte(j), &mArray); err == nil {

		for i := range mArray {
			for k, v := range mArray[i] {
				fixed := SnakeCaseToCamelCase(k)
				delete(mArray[i], k)
				mArray[i][fixed] = JsonConvertCamelCase(v)
			}
		}

		b, err := json.Marshal(mArray)
		if err != nil {
			return j
		}

		return json.RawMessage(b)
	}

	m := make(map[string]json.RawMessage)
	if err := json.Unmarshal([]byte(j), &m); err == nil {
		for k, v := range m {
			fixed := SnakeCaseToCamelCase(k)
			delete(m, k)
			m[fixed] = JsonConvertCamelCase(v)
		}

		b, err := json.Marshal(m)
		if err != nil {
			return j
		}

		return json.RawMessage(b)
	}

	return j
}

func SnakeCaseToCamelCase(key string) string {
	result := ""
	parts := strings.Split(key, "_")
	for i, part := range parts {
		if i > 0 {
			part = fmt.Sprint(strings.Title(strings.ToLower(part)))
		}

		result += part
	}

	return result
}

func ReadInt(r *http.Request, param string, v int) int {
	p := r.FormValue(param)
	rs, err := strconv.Atoi(p)
	if err != nil {
		return v
	}
	return rs
}

func ReadPagination(w http.ResponseWriter, r *http.Request) (int, int) {
	return ReadInt(r, "page", 0), ReadInt(r, "page_size", 100)
}

func ReadJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, v)
}

func WriteJSON(ctx *gin.Context, v interface{}) {
	data, _ := json.Marshal(v)
	if ctx.Request.FormValue("camel_case") == "true" {
		data = JsonConvertCamelCase(data)
	}

	ctx.Writer.Header().Set("Content-Type", "application/json; charsetString=utf-8")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Write(data)

	log.Println(fmt.Sprintf("%s", string(data)))
}

func ReturnError(w http.ResponseWriter, err error) {
	er := HTTPError{
		Result:  false,
		Message: strings.Title(err.Error()),
	}

	ResponseWithJson(w, http.StatusBadRequest, er)
}

func ResponseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Cache-Control", "no-cache, no-store")
	writer.Header().Set("Pragma", "no-cache")
	writer.Header().Set("Expires", "0")
	writer.Header().Set("Content-Type", "application/json; charsetString=utf-8")
	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}
