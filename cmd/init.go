package main

import (
	"context"
	"getcare-notification/common"
	"getcare-notification/common/flags"
	"getcare-notification/config"
	"getcare-notification/internal/controller"
	firebaseSrv "getcare-notification/internal/delivery/firebase"
	kafkaGroup "getcare-notification/internal/delivery/kafka"
	"getcare-notification/internal/delivery/route"
	"getcare-notification/internal/domain"
	"getcare-notification/internal/repository"
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
	"github.com/joho/godotenv"
	"github.com/opentracing/opentracing-go"
	"github.com/segmentio/kafka-go"
	"github.com/urfave/cli"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	app          = NewApp()
	server       = new(Srv)
	startCommand = cli.Command{
		Action:      flags.MigrateFlags(Start),
		Name:        "start",
		Usage:       "start server",
		ArgsUsage:   "<genesisPath>",
		Flags:       []cli.Flag{},
		Description: `start server`,
		//SkipFlagParsing: true,
	}
)

type Srv struct {
	cfg      *config.Config
	logger   logger.Logger
	tracer   opentracing.Tracer
	validate *validator.Validate
	mongoDB  *mongo.Database
	mysqlDB  *gorm.DB
	kafka    *kafka.Conn
	// redis    *redis.Client
	firebase *messaging.Client
	Routes   route.Routes
}

func init() {
	app.Action = cli.ShowAppHelp
	app.Commands = []cli.Command{
		startCommand,
	}
	app.Flags = []cli.Flag{
		flags.HttpHostFlag,
		flags.HttpPortFlag,
		flags.ServerNameFlag,
		flags.ServerVersionFlag,

		flags.MongoDatabaseNameFlag,
		flags.MongoURIFlag,

		flags.StorageAccessKeyFlag,
		flags.StorageSecretKeyFlag,
		flags.StorageRegionFlag,
		flags.StorageNameFlag,

		flags.JaegerHostFlag,
		flags.JaegerPortFlag,
	}
}

// NewApp creates an app with sane defaults.
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Action = cli.ShowAppHelp
	app.Name = "Getcare Notification"
	app.Author = "Getcare"
	app.Email = "getcare@getcare.com"
	app.Usage = "Server Notification API"
	return app
}

// Start ...
func Start(ctx *cli.Context) error {
	server.Start(ctx)
	return nil
}

func (server *Srv) Start(cli *cli.Context) {
	// load env in localhost
	server.LoadEnv()
	// load config
	server.LoadConfig(cli)
	// config log
	server.Logger()
	// config tracer
	closer := server.Tracer()
	defer closer.Close()
	// connect kafka
	server.Kafka()
	// connect mysql
	server.Mysql()
	mysqlDB, _ := server.mysqlDB.DB()
	defer mysqlDB.Close()
	// connect mongodb
	ctx := server.MongoDB()
	defer server.mongoDB.Client().Disconnect(ctx)
	// create validate
	server.Validate()
	// connect firebase
	server.Firebase()
	// config route
	server.configRoute()
	// run server
	log.Fatal(server.Routes.Run())
}

func (server *Srv) LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
}

func (src *Srv) LoadConfig(ctx *cli.Context) {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	// config flags
	config.GetFlags(cfg.Http, ctx,
		flags.HttpHostFlag.GetName(),
	)
	server.cfg = cfg
}

func (server *Srv) MongoDB() context.Context {
	ctx := context.Background()
	conn, err := mongodb.New(ctx, server.cfg)
	if err != nil {
		server.logger.Fatal("mongodb.New: ", err)
	}
	server.mongoDB = conn
	server.logger.Infof("MongoDB connected: %v", server.cfg.MongoDB.URI)
	return ctx
}

func (server *Srv) Mysql() {
	mysql, err := mysqldb.New(server.cfg)
	if err != nil {
		server.logger.Fatal("mysqldb.New: ", err)
	}
	server.mysqlDB = mysql
	server.logger.Infof("MysqlDB connected: %s:%s", server.cfg.MysqlDB.Host, server.cfg.MysqlDB.Port)
}

func (server *Srv) Logger() {
	logger := logger.New(server.cfg)
	logger.Init()
	logger.Infof(
		"AppVersion: %s, LogLevel: %s, DevelopmentMode: %s",
		server.cfg.AppVersion,
		server.cfg.Logger.Level,
		server.cfg.Grpc.Development,
	)
	server.logger = logger
}

func (server *Srv) Tracer() io.Closer {
	tracer, closer, err := jaeger.New(server.cfg)
	if err != nil {
		server.logger.Fatal("jaeger.New: ", err)
		return closer
	}
	server.tracer = tracer
	opentracing.SetGlobalTracer(server.tracer)
	server.logger.Info("Opentracing connected")
	return closer
}

func (server *Srv) Kafka() {
	conn, err := kafkaSrv.New(server.cfg)
	if err != nil {
		server.logger.Fatal("NewKafkaConn: ", err)
	}
	server.kafka = conn
	brokers, err := conn.Brokers()
	if err != nil {
		server.logger.Fatal("conn.Brokers: ", err)
	}
	server.logger.Infof("Kafka connected: %v", brokers)
}

func (server *Srv) Validate() {
	server.validate = validator.New()
	server.logger.Infof("Validate created")
}

func (server *Srv) Firebase() {
	client, err := firebase.Init(server.cfg)
	if err != nil {
		server.logger.Fatal("firebase.Init: ", err)
	}
	server.firebase = client
	server.logger.Infof("Firebase ProjectID: %v", server.cfg.Firebase.ProjectID)
}

func (server *Srv) configRoute() {
	userFcmRepository := repository.NewUserFcmRepository(server.mysqlDB)
	requestOtpRepository := repository.NewRequestOtpRepository(server.mysqlDB)
	verifyOtpRepository := repository.NewVerifyOtpRepository(server.mysqlDB)
	logMessageRepository := repository.NewLogMessageRepository(server.mongoDB)
	userFcmDomain := domain.NewUserFcmDomain(userFcmRepository, &firebaseSrv.Firebase{
		Client: server.firebase,
	})
	logMessageDomain := domain.NewLogMessageDomain(logMessageRepository)
	server.Routes = route.NewRoute(kafkaGroup.NewKafkaGroup(server.cfg.Kafka.Brokers,
		common.NotificationGroupId, server.logger, server.validate), gin.Default(), &controller.Controller{
		RequestOtpController: controller.NewRequestOtpController(
			domain.NewRequestOtpDomain(
				requestOtpRepository,
				verifyOtpRepository,
			),
		),
		VerifyOtpController: controller.NewVerifyOtpController(
			domain.NewVerifyOtpDomain(
				verifyOtpRepository,
			),
		),
		UserFcmController: controller.NewUserFcmController(
			userFcmDomain,
		),
		LogMessageController: controller.NewLogMessageController(
			logMessageDomain,
		),
	},
		server.cfg,
		server.logger,
		server.validate,
	)
}
