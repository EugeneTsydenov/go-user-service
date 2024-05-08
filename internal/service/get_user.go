package service

import (
	"github.com/EugeneTsydenov/go-user-service/internal/domain/entity"
)

type GetUserOutput struct {
	UserData *entity.UserOutput
	Message  string
	Success  bool
}

func (s *Service) GetUser(id int64) GetUserOutput {
	userFromDb, err := s.repo.GetUserById(id)
	if err != nil {
		return GetUserOutput{UserData: nil, Message: "User not found", Success: false}
	}

	return GetUserOutput{UserData: &entity.UserOutput{
		Id: userFromDb.Id, Avatar: userFromDb.Avatar, Username: userFromDb.Username, CreatedAt: userFromDb.CreatedAt,
	}, Success: true, Message: "User successfully retrieved"}
}
