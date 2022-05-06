package controller

import (
	"getcare-notification/constant/errors"
	"getcare-notification/internal/model"
	"getcare-notification/internal/service"
	"getcare-notification/utils"

	"github.com/gin-gonic/gin"
)

type VerifyOtpController interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type verifyOtpController struct {
	Service service.VerifyOtpService
}

func NewVerifyOtpController(verifyOtpService service.VerifyOtpService) VerifyOtpController {
	return &verifyOtpController{
		Service: verifyOtpService,
	}
}

func (vo *verifyOtpController) List(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer
	page, pageSize := utils.ReadPagination(w, r)
	order := r.FormValue("order")

	if err := utils.ValidateRequest(utils.InitializeContext(r), r, "verify_otp", model.RetrieveMany); err != nil {
		utils.ReturnError(w, err)
		return
	}

	where := map[string]interface{}{
		"name": r.FormValue("name"),
	}
	utils.WhereTrim(where)

	records, totalRows, err := vo.Service.List(page, pageSize, order, where)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	result := &utils.PagedResults{Result: true, Page: page, PageSize: pageSize, Data: records, TotalRecords: totalRows}
	utils.WriteJSON(ctx, result)
}

func (vo *verifyOtpController) Get(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer

	argId, err := utils.ParseInt32(ctx.Param("id"))
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	if err := utils.ValidateRequest(utils.InitializeContext(r), r, "verify_otp", model.RetrieveOne); err != nil {
		utils.ReturnError(w, err)
		return
	}

	record, err := vo.Service.Get(argId)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	result := &utils.PagedResult{Result: true, Data: record}
	utils.WriteJSON(ctx, result)
}

func (vo *verifyOtpController) Create(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer

	verifyotp := &model.VerifyOtp{}

	if err := utils.ReadJSON(r, verifyotp); err != nil {
		utils.ReturnError(w, errors.ErrBadParams)
		return
	}

	// if err := verifyotp.BeforeSave(); err != nil {
	// 	service.ReturnError(w, r, repository.ErrBadParams)
	// }

	// verifyotp.Prepare()

	// if err := verifyotp.Validate(model.Create); err != nil {
	// 	service.ReturnError(w, r, repository.ErrBadParams)
	// 	return
	// }

	if err := utils.ValidateRequest(utils.InitializeContext(r), r, "verify_otp", model.Create); err != nil {
		utils.ReturnError(w, err)
		return
	}

	var err error
	verifyotp, _, err = vo.Service.Create(verifyotp)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	result := &utils.PagedResult{Result: true, Data: verifyotp}
	utils.WriteJSON(ctx, result)
}

func (vo *verifyOtpController) Update(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer

	argId, err := utils.ParseInt32(ctx.Param("id"))
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	verifyotp := &model.VerifyOtp{}
	if err := utils.ReadJSON(r, verifyotp); err != nil {
		utils.ReturnError(w, errors.ErrBadParams)
		return
	}

	// if err := verifyotp.BeforeSave(); err != nil {
	// 	service.ReturnError(w, r, repository.ErrBadParams)
	// }

	// verifyotp.Prepare()

	// if err := verifyotp.Validate(model.Update); err != nil {
	// 	service.ReturnError(w, r, repository.ErrBadParams)
	// 	return
	// }

	if err := utils.ValidateRequest(utils.InitializeContext(r), r, "verify_otp", model.Update); err != nil {
		utils.ReturnError(w, err)
		return
	}

	verifyotp, _, err = vo.Service.Update(
		argId,
		verifyotp)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	result := &utils.PagedResult{Result: true, Data: verifyotp}
	utils.WriteJSON(ctx, result)
}

func (vo *verifyOtpController) Delete(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer

	argId, err := utils.ParseInt32(ctx.Param("id"))
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	if err := utils.ValidateRequest(utils.InitializeContext(r), r, "verify_otp", model.Delete); err != nil {
		utils.ReturnError(w, err)
		return
	}

	rowsAffected, err := vo.Service.Delete(argId)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	result := &utils.PagedResult{Result: true, Data: rowsAffected}
	utils.WriteJSON(ctx, result)
}
