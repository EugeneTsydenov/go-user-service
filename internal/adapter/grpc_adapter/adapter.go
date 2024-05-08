package grpc_adapter

import (
	"context"
	"github.com/EugeneTsydenov/go-user-service/internal/ports/grpc_port"
	"github.com/EugeneTsydenov/go-user-service/internal/proto"
	"github.com/EugeneTsydenov/go-user-service/internal/service"
)

var _ grpc_port.Port = (*Implementation)(nil)

type Implementation struct {
	proto.UnimplementedUserServiceServer
	userService service.ServiceInterface
}

func NewImplementation(userService *service.Service) *Implementation {
	return &Implementation{userService: userService}
}

func (i *Implementation) GetUser(_ context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	/*userData := i.userService.GetUser(req.GetId())*/
	return nil, nil
}

func (i *Implementation) Login(_ context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	res := i.userService.Login(req.GetUsername(), req.GetPassword())
	return &proto.LoginResponse{Id: res.Id, Message: res.Message, Success: res.Success}, nil
}

func (i *Implementation) Register(_ context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	res := i.userService.Register(req.GetUsername(), req.GetPassword())
	return &proto.RegisterResponse{Message: res.Message, Success: res.Success}, nil
}

func (i *Implementation) UpdateUser(_ context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	return &proto.UpdateUserResponse{}, nil
}

func (i *Implementation) DeleteUser(_ context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	res := i.userService.DeleteUser(req.GetId())
	return &proto.DeleteUserResponse{Success: res.Success, Message: res.Message}, nil
}

func (i *Implementation) ChangePassword(_ context.Context, req *proto.ChangePasswordRequest) (*proto.ChangePasswordResponse, error) {
	res := i.userService.ChangePassword(req.GetId(), req.GetNewPassword(), req.GetOldPassword())
	return &proto.ChangePasswordResponse{Success: res.Success, Message: res.Message}, nil
}
