package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/xiaokangwang/keymgr/def"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
	"strings"
)

func main() {
	grpcTarget := mustGetFromEnv("GRPC_TARGET")
	authority := os.Getenv("AUTHORITY")
	clientConn := getGRPC(grpcTarget, authority)
	defer clientConn.Close()
	tpmKeyServiceClient := def.NewTPMKeyServiceClient(clientConn)

	result, err := tpmKeyServiceClient.GetPrimaryKeyHash(context.Background(), &def.GetPrimaryKeyHashReq{})
	if err != nil {
		panic(err)
	}
	serverPrimaryKey := result.PrimaryKey
	fmt.Fprintf(os.Stderr, "Primary Key=%x\n", serverPrimaryKey)

	action := os.Args[1]
	switch action {
	case "ssh-agent":
		runSSHAgent(tpmKeyServiceClient)
		return
	case "ssh-keygen":
		KeySeed := os.Args[2]
		KeySend_WithSecret := KeySeed + mustGetFromEnv("KEY_SEED_SUFFIX")
		result, err := tpmKeyServiceClient.CreateECDSAKey(context.Background(), &def.CreateECDSAKeyReq{KeySeed: KeySend_WithSecret})
		if err != nil {
			panic(err)
		}
		SshPublicKey := strings.TrimRight(result.SshPublicKey, "\n")
		publicKeyResult := fmt.Sprintf("%s %s@%x.tpmkeymgr\n", SshPublicKey, KeySeed, serverPrimaryKey)
		fmt.Println(publicKeyResult)
	}

}

func mustGetFromEnv(name string) string {
	data, ok := os.LookupEnv(name)
	if ok {
		return data
	}
	panic("cannot find " + name)
}

func getGRPC(grpcTarget, authority string) *grpc.ClientConn {
	cert, err := tls.LoadX509KeyPair("client.pem", "client.key")
	if err != nil {
		fmt.Printf("load key: %s", err)
	}

	certPEMBlock, err := os.ReadFile("ca.pem")
	if err != nil {
		fmt.Printf("load ca: %s", err)
	}
	caCertByte, _ := pem.Decode(certPEMBlock)
	caCert, err := x509.ParseCertificate(caCertByte.Bytes)
	if err != nil {
		fmt.Printf("load ca: %s", err)
	}
	pool := x509.NewCertPool()
	pool.AddCert(caCert)
	config := &tls.Config{Certificates: []tls.Certificate{cert}, RootCAs: pool, NextProtos: []string{"h2"}}

	effectiveAuthority := authority
	if effectiveAuthority == "" {
		host, _, err := net.SplitHostPort(grpcTarget)
		if err != nil {
			panic(err)
		}
		effectiveAuthority = host
	}

	clientConn, err := grpc.Dial(grpcTarget, grpc.WithAuthority(effectiveAuthority), grpc.WithTransportCredentials(credentials.NewTLS(config)))

	if err != nil {
		fmt.Printf("create grpc connection: %s", err)
	}
	return clientConn
}
