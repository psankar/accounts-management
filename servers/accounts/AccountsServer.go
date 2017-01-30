package main

import (
	"log"
	"net"

	"github.com/psankar/accounts-management/libs/accounts"
	pb "github.com/psankar/accounts-management/proto"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting accounts server")
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAccountsServer(grpcServer, &accounts.AccountsServer{})
	grpcServer.Serve(lis)
}
