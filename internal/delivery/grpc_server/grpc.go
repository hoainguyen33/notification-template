package grpc_server

import (
	"getcare-notification/config"

	"getcare-notification/pkg/logger"

	"google.golang.org/grpc"
)

const (
	certFile        = "ssl/server-cert.pem"
	keyFile         = "ssl/server-key.pem"
	maxHeaderBytes  = 1 << 20
	gzipLevel       = 5
	stackSize       = 1 << 10 // 1 KB
	csrfTokenHeader = "X-CSRF-Token"
	bodyLimit       = "2M"
)

type Grpc struct {
	Server *grpc.Server
	Cfg    *config.Config
	Log    logger.Logger
}

func NewGrpcServer(cfg *config.Config,
	log logger.Logger) (*Grpc, error) {
	return &Grpc{
		Cfg: cfg,
		Log: log,
	}, nil
}
