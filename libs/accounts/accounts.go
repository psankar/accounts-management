package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	pb "github.com/psankar/accounts-management/libs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type AccountsServer struct {
	// TODO: Yet to implement
}

func (s *AccountsServer) SignUp(context.Context, *pb.SignUpRequest) (*pb.GenericResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "")
}
func (s *AccountsServer) CheckAvailability(context.Context, *pb.Username) (*pb.AvailabilityResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "")
}
func (s *AccountsServer) CreateAccount(context.Context, *pb.NewAccountDetails) (*pb.GenericResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "")
}
func (s *AccountsServer) SignIn(context.Context, *pb.SignInRequest) (*pb.GenericResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "")
}
func (s *AccountsServer) UpdateAccount(context.Context, *pb.Settings) (*pb.GenericResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "")
}
func (s *AccountsServer) UpdateUserName(context.Context, *pb.Username) (*pb.GenericResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "")
}
func (s *AccountsServer) ChangeEmailAddress(context.Context, *pb.GenericString) (*pb.GenericResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "")
}
func (s *AccountsServer) ConfirmEmailChange(context.Context, *pb.GenericString) (*pb.GenericResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "")
}
func (s *AccountsServer) ChangePhone(context.Context, *pb.GenericString) (*pb.GenericResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "")
}
func (s *AccountsServer) ConfirmPhoneChange(context.Context, *pb.GenericString) (*pb.GenericResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "")
}
func (s *AccountsServer) SignOut(context.Context, *pb.NilParam) (*pb.GenericResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "")
}
func (s *AccountsServer) DeleteAccount(context.Context, *pb.NilParam) (*pb.GenericResponse, error) {
	return nil, grpc.Errorf(codes.Unimplemented, "")
}

func main() {
	log.Println("Starting accounts server")
	lis, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAccountsServer(grpcServer, &AccountsServer{})
	grpcServer.Serve(lis)
}
