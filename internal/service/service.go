package service

import (
	"github.com/EugeneTsydenov/go-user-service/internal/domain/repository"
	"github.com/EugeneTsydenov/go-user-service/internal/proto"
)

type ServiceInterface interface {
	GetUser(id int64) GetUserOutput
	Login(username, password string) LoginOutput
	Register(username, password string) RegisterOutput
	DeleteUser(id int64) DeleteUserOutput
	UpdateUser(request *proto.UpdateUserRequest) UpdateUserOutput
	UpdatePassword(id int64, newPassword, oldPassword string) UpdatePasswordOutput
	GetAllUsers() *GetAllUsersOutput
}

var _ ServiceInterface = (*Service)(nil)

type Service struct {
	repo repository.RepoInterface
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}
