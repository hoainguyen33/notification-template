package errors

import "fmt"

var (
	ErrPermissionDenied = fmt.Errorf("Permission denied!")
	ErrUserNotFound     = fmt.Errorf("User not found!")
	ErrTypeDontSupport  = fmt.Errorf("Type don't support!")
	ErrBadRequestBody   = fmt.Errorf("Body request invalid!")
	ErrRecordExisted    = fmt.Errorf("record existed!")
	ErrRecordNotFound   = fmt.Errorf("record not found!")
	ErrAcceptExisted    = fmt.Errorf("accept existed!")
	ErrBadParams        = fmt.Errorf("bad params error")
)
