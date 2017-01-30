package accounts

import (
	"log"

	"golang.org/x/net/context"

	pb "github.com/psankar/accounts-management/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type AccountsServer struct {
	// TODO: Yet to implement
}

func (s *AccountsServer) SignUp(c context.Context, req *pb.SignUpRequest) (*pb.GenericResponse, error) {
	log.Println(req.EmailAddress, req.PhoneNumber, req)
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
