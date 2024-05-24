package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/noonyuu/grpc-graphql/grpc/pb"
	"github.com/noonyuu/grpc-graphql/grpc/impl"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("grpc:%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, &impl.UserServiceServer{})
	grpcServer.Serve(lis)
}
