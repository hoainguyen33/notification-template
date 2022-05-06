package grpc

import (
	"crypto/tls"
	"getcare-notification/constant/config"
	"getcare-notification/pkg/logger"
	"time"

	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

func NewServer(cfg *config.Config, log logger.Logger, certFile, keyFile string) (*grpc.Server, error) {
	// server cert
	serverCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return grpc.NewServer(
		grpc.Creds(credentials.NewTLS(config)),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: cfg.Grpc.MaxConnectionIdle * time.Minute, // time.Duration
			Timeout:           cfg.Grpc.Timeout * time.Second,           // time.Duration
			MaxConnectionAge:  cfg.Grpc.MaxConnectionAge * time.Minute,  // time.Duration
			Time:              cfg.Grpc.Timeout * time.Minute,           // time.Duration
		}),
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpcrecovery.UnaryServerInterceptor(),
			log.Logger,
		),
	), nil
}
