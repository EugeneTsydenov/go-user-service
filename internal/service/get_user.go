package service

import (
	"github.com/EugeneTsydenov/go-user-service/internal/domain/entity"
)

type GetUserOutput struct {
	UserData *entity.UserOutput
	Message  string
	Code     int32
}

func (s *Service) GetUser(id int64) GetUserOutput {
	userFromDb, err := s.repo.GetUserById(id)
	if err != nil {
		return GetUserOutput{UserData: nil, Message: "User not found", Code: 404}
	}

	return GetUserOutput{UserData: &entity.UserOutput{
		Id: userFromDb.ID, Avatar: userFromDb.Avatar, Username: userFromDb.Username, CreatedAt: userFromDb.CreatedAt,
	}, Code: 200, Message: "User successfully retrieved"}
}
