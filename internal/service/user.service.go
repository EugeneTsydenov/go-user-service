package service

import (
	"github.com/EugeneTsydenov/go-user-service/internal/proto"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedUserServiceServer
}

func UserService(s *grpc.Server) {
	proto.RegisterUserServiceServer(s, &server{})
}
