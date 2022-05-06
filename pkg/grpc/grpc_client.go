package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials(certFile string) (credentials.TransportCredentials, error) {
	pemServerCA, err := ioutil.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}
	config := &tls.Config{
		RootCAs: certPool,
	}
	return credentials.NewTLS(config), nil
}

func NewGrpcSrv(addr string, certFile string) (*grpc.ClientConn, error) {
	flag.Parse()
	tlsCredentials, err := loadTLSCredentials(certFile)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
