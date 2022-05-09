package controller

import (
	"getcare-notification/common/errors"
	"getcare-notification/internal/domain"
	"getcare-notification/internal/model"
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
	Push(ctx *gin.Context)
	KafkaPush(msg *model.KafkaMessage) error
}

type userFcmController struct {
	UserFcmDomain domain.UserFcmDomain
}

func NewUserFcmController(userFcmDomain domain.UserFcmDomain) UserFcmController {
	return &userFcmController{
		UserFcmDomain: userFcmDomain,
	}
}

func (uf *userFcmController) List(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer
	page, pageSize := utils.ReadPagination(w, r)
	order := r.FormValue("order")
	where := (&utils.Where{}).FromMap(utils.WhereMap{
		"name": utils.WT("string", "LIKE"),
	}).LoadData(r)
	records, total, err := uf.UserFcmDomain.List(page, pageSize, order, where)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}
	result := &utils.PagedResults{Result: true, Page: page, PageSize: pageSize, Data: records, TotalRecords: total}
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
	record, err := uf.UserFcmDomain.Get(argId)
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
		utils.ReturnError(w, errors.ErrBadParams)
		return
	}
	user := ctx.Keys["user"].(*gocloak.UserInfo)
	if user != nil {
		userFcmAdd.UserID = *user.Sub
	} else {
		userFcmAdd.UserID = ""
	}
	record, err := uf.UserFcmDomain.Create(userFcmAdd)
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

	utils.ReturnError(w, errors.ErrBadParams)
}

func (uf *userFcmController) Delete(ctx *gin.Context) {
	// r := ctx.Request
	w := ctx.Writer

	argId, err := utils.ParseInt32(ctx.Param("id"))
	if err != nil {
		utils.ReturnError(w, err)
		return
	}
	err = uf.UserFcmDomain.Delete(argId)
	if err != nil {
		utils.ReturnError(w, err)
		return
	}
	result := &utils.PagedResult{Result: true, Data: argId}
	utils.WriteJSON(ctx, result)
}

func (uf *userFcmController) Push(ctx *gin.Context) {
	r := ctx.Request
	w := ctx.Writer
	msg := &model.KafkaMessage{}
	if err := utils.ReadJSON(r, msg); err != nil {
		utils.ReturnError(w, errors.ErrBadParams)
		return
	}
	if err := uf.KafkaPush(msg); err != nil {
		utils.ReturnError(w, errors.ErrBadParams)
		return
	}
	result := &utils.PagedResult{Result: true}
	utils.WriteJSON(ctx, result)
}

func (uf *userFcmController) KafkaPush(msg *model.KafkaMessage) error {
	switch msg.Event {
	case "push/user":
		data := msg.Msg.(model.MessageUser)
		return uf.UserFcmDomain.PushUser(&data)
	case "push/users":
		data := msg.Msg.(model.MessageUsers)
		return uf.UserFcmDomain.PushUsers(&data)
	}
	return nil
}
