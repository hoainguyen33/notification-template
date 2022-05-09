package controller

import (
	"getcare-notification/common/errors"
	"getcare-notification/internal/domain"
	"getcare-notification/internal/model"
	"getcare-notification/utils"
	"os"

	"github.com/gin-gonic/gin"
)

type RequestOtpController interface {
	RequestOTP(ctx *gin.Context)
	VerifyOTP(ctx *gin.Context)
}

type requestOtpController struct {
	RequestOtpDomain domain.RequestOtpDomain
}

func NewRequestOtpController(requestOtpDomain domain.RequestOtpDomain) RequestOtpController {
	return &requestOtpController{
		RequestOtpDomain: requestOtpDomain,
	}
}

func GetAllRequestOtp(ctx *gin.Context) {
	// r := ctx.Request
	w := ctx.Writer

	utils.ReturnError(w, errors.ErrPermissionDenied)
}

func GetRequestOtp(ctx *gin.Context) {
	// r := ctx.Request
	w := ctx.Writer

	utils.ReturnError(w, errors.ErrPermissionDenied)
}

func AddRequestOtp(ctx *gin.Context) {
	// r := ctx.Request
	w := ctx.Writer

	utils.ReturnError(w, errors.ErrPermissionDenied)
}

func UpdateRequestOtp(ctx *gin.Context) {
	// r := ctx.Request
	w := ctx.Writer

	utils.ReturnError(w, errors.ErrPermissionDenied)
}

func DeleteRequestOtp(ctx *gin.Context) {
	// r := ctx.Request
	w := ctx.Writer

	utils.ReturnError(w, errors.ErrPermissionDenied)
}

func (ro *requestOtpController) RequestOTP(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer

	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN_SERVER1") && token != os.Getenv("TOKEN_SERVER2") {
		utils.ReturnError(w, errors.ErrPermissionDenied)
		return
	}

	requestOtpAdd := &model.RequestOtpAdd{}
	if err := utils.ReadJSON(r, requestOtpAdd); err != nil {
		utils.ReturnError(w, errors.ErrBadParams)
		return
	}

	err := ro.RequestOtpDomain.RequestOtpAdd(*requestOtpAdd)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	result := &utils.PagedResult{Result: true}
	utils.WriteJSON(ctx, result)
}

func (ro *requestOtpController) VerifyOTP(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer

	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN_SERVER1") && token != os.Getenv("TOKEN_SERVER2") {
		utils.ReturnError(w, errors.ErrPermissionDenied)
		return
	}

	verifyOTP := &model.VerifyOTPParam{}
	if err := utils.ReadJSON(r, verifyOTP); err != nil {
		utils.ReturnError(w, errors.ErrBadParams)
		return
	}

	err := ro.RequestOtpDomain.VerifyOTP(*verifyOTP)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	result := &utils.PagedResult{Result: true}
	utils.WriteJSON(ctx, result)
}
