package service

import (
	"github.com/EugeneTsydenov/go-user-service/internal/domain/repository"
	"github.com/EugeneTsydenov/go-user-service/internal/proto"
)

type ServiceInterface interface {
	GetUser(id int64) error
	Login(username, password string) (*LoginOutput, error)
	Register(username, password string) (RegisterOutput, error)
	DeleteUser(id int64) (*DeleteUserOutput, error)
	UpdateUser(request *proto.UpdateUserRequest) (*UpdateUserOutput, error)
	ChangePassword(id int64, newPassword, oldPassword string) (*ChangePasswordOutput, error)
}

var _ ServiceInterface = (*Service)(nil)

type Service struct {
	repo repository.RepoInterface
}

func New(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}
