package controller

import (
	"getcare-notification/internal/model"
	"getcare-notification/internal/repository"
	"getcare-notification/internal/service"
	"getcare-notification/utils"

	"github.com/Nerzal/gocloak/v10"
	"github.com/gin-gonic/gin"
)

type UserFcmController interface {
	List(ctx *gin.Context)
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	PushUserFcm(ctx *gin.Context)
	Push(userID string, title string, body string, data interface{}) error
}

type userFcmController struct {
	Service service.UserFcmService
}

func NewUserFcmController(userFcmService service.UserFcmService) UserFcmController {
	return &userFcmController{
		Service: userFcmService,
	}
}

func (uf *userFcmController) List(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer
	page, pageSize := utils.ReadPagination(w, r)

	order := r.FormValue("order")
	where := map[string]interface{}{
		"name": r.FormValue("name"),
	}
	utils.WhereTrim(where)

	records, totalRows, err := uf.Service.List(page, pageSize, order, where)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	result := &utils.PagedResults{Result: true, Page: page, PageSize: pageSize, Data: records, TotalRecords: totalRows}
	utils.WriteJSON(ctx, result)
}

func (uf *userFcmController) Get(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer

	argId, err := utils.ParseInt32(ctx.Param("id"))
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	if err := utils.ValidateRequest(utils.InitializeContext(r), r, "user_fcm", model.RetrieveOne); err != nil {
		utils.ReturnError(w, err)
		return
	}

	record, err := uf.Service.Get(argId)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	result := &utils.PagedResult{Result: true, Data: record}
	utils.WriteJSON(ctx, result)
}

func (uf *userFcmController) Create(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer

	userFcmAdd := &model.UserFcmAdd{}

	if err := utils.ReadJSON(r, userFcmAdd); err != nil {
		utils.ReturnError(w, repository.ErrBadParams)
		return
	}
	user := ctx.Keys["user"].(*gocloak.UserInfo)
	if user != nil {
		userFcmAdd.UserID = *user.Sub
	} else {
		userFcmAdd.UserID = ""
	}
	record, err := uf.Service.Create(userFcmAdd)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	result := &utils.PagedResult{Result: true, Data: record}
	utils.WriteJSON(ctx, result)
}

func (uf *userFcmController) Update(ctx *gin.Context) {
	// r := ctx.Request
	w := ctx.Writer

	utils.ReturnError(w, repository.ErrBadParams)
}

func (uf *userFcmController) Delete(ctx *gin.Context) {
	// r := ctx.Request
	w := ctx.Writer

	argId, err := utils.ParseInt32(ctx.Param("id"))
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	rowsAffected, err := uf.Service.Delete(argId)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	result := &utils.PagedResult{Result: true, Data: rowsAffected}
	utils.WriteJSON(ctx, result)
}

func (uf *userFcmController) PushUserFcm(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer

	pushUserFcm := &model.PushUserFcm{}

	if err := utils.ReadJSON(r, pushUserFcm); err != nil {
		utils.ReturnError(w, repository.ErrBadParams)
		return
	}

	err := uf.Service.Push(pushUserFcm.UserID, pushUserFcm.Title, pushUserFcm.Body, pushUserFcm.Data)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}

	result := &utils.PagedResult{Result: true}
	utils.WriteJSON(ctx, result)
}

func (uf *userFcmController) Push(userID string, title string, body string, data interface{}) error {
	return uf.Service.Push(userID, title, body, data)
}
