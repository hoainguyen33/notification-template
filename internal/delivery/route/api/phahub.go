package api

import (
	"getcare-notification/internal/controller"
	"getcare-notification/internal/middleware"
	"html/template"
	"net/http"
	"time"

	grpcClient "getcare-notification/internal/delivery/grpc_client"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	_ "github.com/satori/go.uuid"
)

func PhahubAPI(address string, r *gin.Engine, apiRoutes *gin.RouterGroup, c *controller.Controller, grpc *grpcClient.Grpcs) error {

	r.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "public",
		Extension: ".html",
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"socket.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})

	// apiRoutes.OPTIONS("/log_message", Options, middleware.Cors())
	// logMessage := apiRoutes.Group("/log_message", middleware.Cors())
	// {
	// 	logMessage.PUT("/", controller.UpdateLogOption)
	// }

	apiRoutes.OPTIONS("/otp", Options, middleware.Cors())
	apiRoutes.OPTIONS("/otp/request", Options, middleware.Cors())
	apiRoutes.OPTIONS("/otp/verify", Options, middleware.Cors())
	otpData := apiRoutes.Group("/otp", middleware.Cors())
	{
		otpData.POST("/request", c.RequestOtpController.RequestOTP)
		otpData.POST("/verify", c.RequestOtpController.VerifyOTP)
	}
	return r.Run(address)
}

func Options(ctx *gin.Context) {

}
