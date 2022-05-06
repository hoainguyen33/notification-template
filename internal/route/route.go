package route

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/satori/go.uuid"

	"getcare-notification/constant/config"
	"getcare-notification/internal/controller"
	grpcClient "getcare-notification/internal/grpc_client"
	"getcare-notification/internal/kafka"
	"getcare-notification/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Routes interface {
	Run() error
	RunAPI() error
	RunKafka() Closer
}

type route struct {
	Cfg         *config.Config
	Log         logger.Logger
	Validate    *validator.Validate
	GrpcsClient *grpcClient.Grpcs
	KafkaGroup  *kafka.KafkaGroup
	ApiGin      *gin.Engine
	Controller  *controller.Controller
}

func NewRoute(
	kafkaGroup *kafka.KafkaGroup,
	apiGin *gin.Engine,
	controller *controller.Controller,
	cfg *config.Config,
	log logger.Logger,
	validate *validator.Validate) Routes {
	return &route{
		ApiGin:     apiGin,
		Controller: controller,
		KafkaGroup: kafkaGroup,
		Cfg:        cfg,
		Log:        log,
		Validate:   validate,
	}
}

func (r *route) Run() error {
	go r.RunKafka()
	return r.RunAPI()
}
