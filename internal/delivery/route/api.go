package route

import (
	_ "github.com/satori/go.uuid"

	"getcare-notification/internal/controller"
	"getcare-notification/internal/delivery/route/api"
	"getcare-notification/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *route) RunAPI() error {
	r.ApiGin.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Our GetCare")
	})

	apiRoutes := r.ApiGin.Group("/api/v1")

	apiRoutes.GET("/heathcheck", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Status: OK")
	})

	apiRoutes.OPTIONS("/request_otp", api.Options, middleware.Cors())
	apiRoutes.OPTIONS("/request_otp/", api.Options, middleware.Cors())
	apiRoutes.OPTIONS("/request_otp/:id", api.Options, middleware.Cors())
	requestOtpRoutes := apiRoutes.Group("/request_otp", middleware.Base(), middleware.Cors())
	{
		requestOtpRoutes.GET("", controller.GetAllRequestOtp)
		requestOtpRoutes.GET("/", controller.GetAllRequestOtp)
		requestOtpRoutes.GET("/:id", controller.GetRequestOtp)
		requestOtpRoutes.POST("", controller.AddRequestOtp)
		requestOtpRoutes.POST("/", controller.AddRequestOtp)
		requestOtpRoutes.PUT("/:id", controller.UpdateRequestOtp)
		requestOtpRoutes.DELETE("/:id", controller.DeleteRequestOtp)
	}

	apiRoutes.OPTIONS("/user_fcm", api.Options, middleware.Cors())
	apiRoutes.OPTIONS("/user_fcm/", api.Options, middleware.Cors())
	apiRoutes.OPTIONS("/user_fcm/:id", api.Options, middleware.Cors())
	userFcmRoutes := apiRoutes.Group("/user_fcm", middleware.Base(), middleware.Cors())
	{
		userFcmRoutes.GET("", r.Controller.UserFcmController.List)
		userFcmRoutes.GET("/", r.Controller.UserFcmController.List)
		userFcmRoutes.GET("/:id", r.Controller.UserFcmController.Get)
		userFcmRoutes.POST("", r.Controller.UserFcmController.Create)
		userFcmRoutes.POST("/", r.Controller.UserFcmController.Create)
		userFcmRoutes.PUT("/:id", r.Controller.UserFcmController.Update)
		userFcmRoutes.DELETE("/:id", r.Controller.UserFcmController.Delete)
	}

	apiRoutes.OPTIONS("/verify_otp", api.Options, middleware.Cors())
	apiRoutes.OPTIONS("/verify_otp/", api.Options, middleware.Cors())
	apiRoutes.OPTIONS("/verify_otp/:id", api.Options, middleware.Cors())
	verifyOtpRoutes := apiRoutes.Group("/verify_otp", middleware.Base(), middleware.Cors())
	{
		verifyOtpRoutes.GET("", r.Controller.VerifyOtpController.List)
		verifyOtpRoutes.GET("/", r.Controller.VerifyOtpController.List)
		verifyOtpRoutes.GET("/:id", r.Controller.VerifyOtpController.Get)
		verifyOtpRoutes.POST("", r.Controller.VerifyOtpController.Create)
		verifyOtpRoutes.POST("/", r.Controller.VerifyOtpController.Create)
		verifyOtpRoutes.PUT("/:id", r.Controller.VerifyOtpController.Update)
		verifyOtpRoutes.DELETE("/:id", r.Controller.VerifyOtpController.Delete)
	}

	errGetcare := api.PhahubAPI(r.Cfg.Http.Address(), r.ApiGin, apiRoutes, r.Controller, r.GrpcsClient)
	if errGetcare != nil {
		return errGetcare
	}

	return nil
}
