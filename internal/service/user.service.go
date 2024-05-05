package service

import (
	"context"
	"fmt"
	"github.com/EugeneTsydenov/go-user-service/internal/proto"
	"github.com/EugeneTsydenov/go-user-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
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
