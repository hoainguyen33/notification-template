package route

import (
	grpcClient "getcare-notification/internal/grpc_client"
)

func NewGrpcClients() (*grpcClient.Grpcs, error) {
	grpcC, err := grpcClient.NewGrpcAccount()
	return &grpcClient.Grpcs{
		Account: grpcC,
	}, err
}
