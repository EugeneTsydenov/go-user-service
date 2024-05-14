package grpc_port

import (
	"context"
	"github.com/EugeneTsydenov/go-user-service/internal/proto"
)

type Port interface {
	GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error)
	Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error)
	Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error)
	UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error)
	ChangePassword(ctx context.Context, req *proto.ChangePasswordRequest) (*proto.ChangePasswordResponse, error)
	DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error)
}
