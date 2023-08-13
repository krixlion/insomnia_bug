package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"github.com/krixlion/insomnia_bug/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	serverCert, err := tls.LoadX509KeyPair("tls.crt", "tls.key")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.VerifyClientCertIfGiven,
	}

	creds := credentials.NewTLS(config)

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	reflection.Register(grpcServer)
	pb.RegisterGreeterServer(grpcServer, GreeterServer{})

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 50055))
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello %s!", req.GetName()),
	}, nil
}
