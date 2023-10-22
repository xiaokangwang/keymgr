package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/xiaokangwang/keymgr/def"
	"github.com/xiaokangwang/keymgr/tpmadm/tpmadmcommander"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
)

func StartServer() error {
	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		log.Fatalf("load key: %s", err)
	}

	certPEMBlock, err := os.ReadFile("ca.pem")
	if err != nil {
		log.Fatalf("load ca: %s", err)
	}
	cacertByte, _ := pem.Decode(certPEMBlock)
	cacert, err := x509.ParseCertificate(cacertByte.Bytes)
	if err != nil {
		log.Fatalf("load ca: %s", err)
	}
	pool := x509.NewCertPool()
	pool.AddCert(cacert)
	config := &tls.Config{Certificates: []tls.Certificate{cert}, ClientAuth: tls.RequireAndVerifyClientCert, ClientCAs: pool, NextProtos: []string{"h2"}}
	tlsCredentials := credentials.NewTLS(config)

	list, err := net.Listen("tcp", mustGetConfFromEnv("LISTEN_ADDR"))
	if err != nil {
		log.Fatalf("tcp listen: %s", err)
	}
	var opts []grpc.ServerOption
	logger := log.WithField("module", "gRPC")
	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_logrus.UnaryServerInterceptor(logger),
	),
		grpc_middleware.WithStreamServerChain(
			grpc_logrus.StreamServerInterceptor(logger),
		),
		grpc.Creds(tlsCredentials))
	grpcServer := grpc.NewServer(opts...)
	def.RegisterTPMKeyServiceServer(grpcServer, tpmadmcommander.NewTpmAdmService())
	log.Println(grpcServer.Serve(list).Error())
	return nil
}

func main() {
	panic(StartServer())
}

func mustGetConfFromEnv(name string) string {
	data, ok := os.LookupEnv(name)
	if ok {
		return data
	}
	panic(fmt.Sprintf("no env %v", name))
}
