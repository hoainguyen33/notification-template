package grpc_client

import (
	grpcSrv "getcare-notification/pkg/grpc"
	ac "getcare-notification/proto/account"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type GrpcAccount interface {
	Login(ctx *gin.Context, username string, password string) (string, error)
}
type grpcAccount struct {
	Conn *grpc.ClientConn
}

func NewGrpcAccount() (GrpcAccount, error) {
	grpcSrvAccount, err := grpcSrv.NewGrpcSrv(os.Getenv("ACCOUNT_GRPC_SERVER"), os.Getenv("CERTFILE_ACCOUNT_GRPC_SERVER"))
	if err != nil {
		return nil, err
	}
	return &grpcAccount{
		Conn: grpcSrvAccount,
	}, nil
}

func (ga *grpcAccount) Login(ctx *gin.Context, username string, password string) (string, error) {
	c := ac.NewAccountsServiceClient(ga.Conn)
	r, err := c.Login(ctx, &ac.LoginReq{Username: username, Password: password})
	if err != nil {
		return "", err
	}
	return r.GetToken(), nil
}
