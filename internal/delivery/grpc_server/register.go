package grpc_server

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	grpc_server "getcare-notification/pkg/grpc"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"google.golang.org/grpc/reflection"
)

func (g *Grpc) Register() error {
	server, err := grpc_server.NewServer(g.Cfg, g.Log, keyFile, certFile)
	if err != nil {
		return err
	}
	g.Server = server
	// register grpc
	l, err := net.Listen("tcp", g.Cfg.Grpc.Port)
	if err != nil {
		return errors.Wrap(err, "net.Listen")
	}
	defer l.Close()
	grpc_prometheus.Register(g.Server)
	go func() {
		g.Log.Infof("GRPC Server is listening on port: %s", g.Cfg.Grpc.Port)
		g.Log.Fatal(g.Server.Serve(l))
	}()
	if g.Cfg.Grpc.Development {
		reflection.Register(g.Server)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// select {
	// case v := <-quit:
	// 	g.Log.Errorf("signal.Notify: %v", v)
	// case done := <-ctx.Done():
	// 	g.Log.Errorf("ctx.Done: %v", done)
	// }

	g.Server.GracefulStop()
	g.Log.Info("Server Exited Properly")
	return nil
}
