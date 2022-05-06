package main

import (
	"context"
	"getcare-notification/constant"
	"getcare-notification/constant/config"
	"getcare-notification/internal/controller"
	kafkaGroup "getcare-notification/internal/kafka"
	"getcare-notification/internal/repository"
	"getcare-notification/internal/route"
	"getcare-notification/internal/service"
	"getcare-notification/pkg/firebase"
	"getcare-notification/pkg/jaeger"
	kafkaSrv "getcare-notification/pkg/kafka"
	"getcare-notification/pkg/logger"
	"getcare-notification/pkg/mongodb"
	"getcare-notification/pkg/mysqldb"
	"io"
	"log"
	"os"

	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Srv struct {
	cfg      *config.Config
	logger   logger.Logger
	tracer   opentracing.Tracer
	validate *validator.Validate
	mongoDB  *mongo.Database
	mysqlDB  *gorm.DB
	kafka    *kafka.Conn
	redis    *redis.Client
	firebase *messaging.Client
	Routes   route.Routes
}

var srv = &Srv{}

func init() {
	log.Println("init application")
	// load env in localhost
	srv.LoadEnv()
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
	// config srv
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}
	srv.cfg = cfg
}

func (srv *Srv) LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func (srv *Srv) Start() {
	// config log
	srv.Logger()
	// config tracer
	closer := srv.Tracer()
	defer closer.Close()
	// connect kafka
	srv.Kafka()
	// connect mysql
	srv.Mysql()
	mysqlDB, _ := srv.mysqlDB.DB()
	defer mysqlDB.Close()
	// connect mongodb
	ctx := srv.MongoDB()
	defer srv.mongoDB.Client().Disconnect(ctx)
	// create validate
	srv.Validate()
	// connect firebase
	srv.Firebase()
	// config route
	srv.configRoute()
	// run server
	log.Fatal(srv.Routes.Run())
}

func (srv *Srv) MongoDB() context.Context {
	ctx := context.Background()
	conn, err := mongodb.New(ctx, srv.cfg)
	if err != nil {
		srv.logger.Fatal("mongodb.New", err)
	}
	srv.mongoDB = conn
	srv.logger.Infof("MongoDB connected: %v", srv.cfg.MongoDB.URI)
	return ctx
}

func (srv *Srv) Mysql() {
	mysql, err := mysqldb.New(srv.cfg)
	if err != nil {
		srv.logger.Fatal("mysqldb.New", err)
	}
	srv.mysqlDB = mysql
	srv.logger.Infof("MysqlDB connected: %s:%s", srv.cfg.MysqlDB.Host, srv.cfg.MysqlDB.Port)
}

func (srv *Srv) Logger() {
	logger := logger.New(srv.cfg)
	logger.Init()
	logger.Infof(
		"AppVersion: %s, LogLevel: %s, DevelopmentMode: %s",
		srv.cfg.AppVersion,
		srv.cfg.Logger.Level,
		srv.cfg.Grpc.Development,
	)
	srv.logger = logger
}

func (srv *Srv) Tracer() io.Closer {
	tracer, closer, err := jaeger.New(srv.cfg)
	if err != nil {
		srv.logger.Fatal("jaeger.New", err)
		return closer
	}
	srv.tracer = tracer
	opentracing.SetGlobalTracer(srv.tracer)
	srv.logger.Info("Opentracing connected")
	return closer
}

func (srv *Srv) Kafka() {
	conn, err := kafkaSrv.NewKafkaConn(srv.cfg.Kafka.Brokers[0])
	if err != nil {
		srv.logger.Fatal("NewKafkaConn", err)
	}
	srv.kafka = conn
	brokers, err := conn.Brokers()
	if err != nil {
		srv.logger.Fatal("conn.Brokers", err)
	}
	srv.logger.Infof("Kafka connected: %v", brokers)
}

func (srv *Srv) Validate() {
	srv.validate = validator.New()
	srv.logger.Infof("Validate created")
}

func (srv *Srv) Firebase() {
	client, err := firebase.Init(srv.cfg)
	if err != nil {
		srv.logger.Fatal("firebase.Init", err)
	}
	srv.firebase = client
	srv.logger.Infof("Firebase ProjectID: %v", srv.cfg.Firebase.ProjectID)
}

func (srv *Srv) configRoute() {
	userFcmRepository := repository.NewUserFcmRepository(srv.mysqlDB)
	requestOtpRepository := repository.NewRequestOtpRepository(srv.mysqlDB)
	verifyOtpRepository := repository.NewVerifyOtpRepository(srv.mysqlDB)
	logMessageRepository := repository.NewLogMessageRepository(srv.mongoDB)
	userFcmService := service.NewUserFcmService(userFcmRepository)
	logMessageService := service.NewLogMessageService(logMessageRepository)
	srv.Routes = route.NewRoute(srv.grpcClient, kafkaGroup.NewKafkaGroup(srv.cfg.Kafka.Brokers,
		constant.NotificationGroupId, srv.logger, srv.validate), gin.Default(), &controller.Controller{
		RequestOtpController: controller.NewRequestOtpController(
			service.NewRequestOtpService(
				requestOtpRepository,
				verifyOtpRepository,
			),
		),
		VerifyOtpController: controller.NewVerifyOtpController(
			service.NewVerifyOtpService(
				verifyOtpRepository,
			),
		),
		UserFcmController: controller.NewUserFcmController(
			userFcmService,
		),
		LogMessageController: controller.NewLogMessageController(
			logMessageService,
		),
	},
		srv.cfg,
		srv.logger,
		srv.validate,
	)
}
