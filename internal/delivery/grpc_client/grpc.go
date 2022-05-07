package grpc_client

type Grpcs struct {
	Account GrpcAccount
}

func NewGrpcClients() (*Grpcs, error) {
	grpcC, err := NewGrpcAccount()
	return &Grpcs{
		Account: grpcC,
	}, err
}
