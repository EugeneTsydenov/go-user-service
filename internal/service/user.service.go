package service

import (
	"context"
	"fmt"
	"github.com/EugeneTsydenov/go-user-service/internal/model"
	"github.com/EugeneTsydenov/go-user-service/internal/proto"
	"github.com/EugeneTsydenov/go-user-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CheckPasswordHash(password, hash string) bool {
	fmt.Println(password, hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

type server struct {
	proto.UnimplementedUserServiceServer
}

func UserService(s *grpc.Server) {
	proto.RegisterUserServiceServer(s, &server{})
}

func (s *server) Register(ctx context.Context, in *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	_, err := repository.GetUserByUsername(in.Username)
	if err == nil {
		return &proto.RegisterResponse{Success: false, Message: "User with this username is already exist!"}, nil
	}
	hashPassword, err := HashPassword(in.Password)
	if err != nil {
		return &proto.RegisterResponse{Success: false, Message: "Something error"}, nil
	}
	err = repository.SaveUser(in.Username, hashPassword)
	if err != nil {
		return &proto.RegisterResponse{Success: false, Message: "Something Error"}, nil
	}
	return &proto.RegisterResponse{Success: true, Message: "User successfully saved"}, nil
}

func (s *server) Login(ctx context.Context, in *proto.LoginRequest) (*proto.LoginResponse, error) {
	userFromDB, err := repository.GetUserByUsername(in.Username)
	fmt.Println(userFromDB)
	if err != nil {
		return &proto.LoginResponse{Success: false, Message: "A user with this username does not exist!"}, nil
	}

	if !CheckPasswordHash(in.Password, userFromDB.HashPassword) {
		return &proto.LoginResponse{Success: false, Message: "Invalid password"}, nil
	}

	return &proto.LoginResponse{Success: true, Id: userFromDB.Id, Message: "Login success"}, nil
}

func (s *server) GetUser(ctx context.Context, in *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	userFromDb, err := repository.GetUserById(in.GetId())
	if err != nil {
		return &proto.GetUserResponse{UserData: nil, Message: "User not found", Success: false}, nil
	}

	return &proto.GetUserResponse{UserData: convertUserDataToProto(userFromDb), Success: true, Message: "User successfully retrieved"}, nil
}

func (s *server) DeleteUser(ctx context.Context, in *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	err := repository.DeleteUser(in.GetId())
	if err != nil {
		return &proto.DeleteUserResponse{Success: false, Message: "User cant be delete"}, nil
	}
	return &proto.DeleteUserResponse{Success: true, Message: "User successfully deleted"}, nil
}

func (s *server) ChangePassword(ctx context.Context, in *proto.ChangePasswordRequest) (*proto.ChangePasswordResponse, error) {
	userFromDB, err := repository.GetUserById(in.GetId())
	if err != nil {
		return &proto.ChangePasswordResponse{Success: false, Message: "User not found, you cant change password"}, nil
	}
	if !CheckPasswordHash(in.GetOldPassword(), userFromDB.HashPassword) {
		return &proto.ChangePasswordResponse{Success: false, Message: "Invalid password"}, nil
	}
	hashPassword, err := HashPassword(in.NewPassword)
	if err != nil {
		return &proto.ChangePasswordResponse{Success: false, Message: "Something Error"}, nil
	}
	err = repository.UpdatePassword(in.GetId(), hashPassword)
	fmt.Println(err)
	if err != nil {
		return &proto.ChangePasswordResponse{Success: false, Message: "Something Error"}, nil
	}
	return &proto.ChangePasswordResponse{Success: true, Message: "Password successfully updated"}, nil
}

func (s *server) UpdateUser(ctx context.Context, in *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	updateData := make(map[string]interface{})

	// Проверяем, установлено ли поле username
	if in.Username != "" {
		updateData["username"] = in.Username
	}

	// Проверяем, установлено ли поле avatar
	if in.Avatar != "" {
		updateData["avatar"] = in.Avatar
	}

	if len(updateData) > 0 {
		err := repository.UpdateUser(in.GetId(), updateData)
		if err != nil {
			return &proto.UpdateUserResponse{Message: "Something Error", Success: false}, nil
		}
		return &proto.UpdateUserResponse{Message: "User successfully updated", Success: true}, nil
	}

	return &proto.UpdateUserResponse{Message: "Nothing Update", Success: false}, nil
}

func convertUserDataToProto(userFromDb model.User) *proto.UserData {
	return &proto.UserData{
		Id:        userFromDb.Id,
		Username:  userFromDb.Username,
		Avatar:    userFromDb.Avatar,
		CreatedAt: timestamppb.New(userFromDb.CreatedAt),
	}
}
