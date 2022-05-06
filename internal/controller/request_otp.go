package controller

import (
	"getcare-notification/constant/errors"
	"getcare-notification/internal/model"
	"getcare-notification/internal/service"
	"getcare-notification/utils"
	"os"

	"github.com/gin-gonic/gin"
)

type RequestOtpController interface {
	RequestOTP(ctx *gin.Context)
	VerifyOTP(ctx *gin.Context)
}

type requestOtpController struct {
	Service service.RequestOtpService
}

func NewRequestOtpController(service service.RequestOtpService) RequestOtpController {
	return &requestOtpController{
		Service: service,
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

	err := ro.Service.RequestOtpAdd(*requestOtpAdd)
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

	err := ro.Service.VerifyOTP(*verifyOTP)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	result := &utils.PagedResult{Result: true}
	utils.WriteJSON(ctx, result)
}
