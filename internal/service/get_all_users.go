package service

import (
	"github.com/EugeneTsydenov/go-user-service/internal/convert"
	"github.com/EugeneTsydenov/go-user-service/internal/proto"
)

type GetAllUsersOutput struct {
	Users []*proto.UserData
}

func (s *Service) GetAllUsers() *GetAllUsersOutput {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return &GetAllUsersOutput{Users: nil}
	}

	usersOutput := make([]*proto.UserData, len(*users))
	for i, user := range *users {
		userOutput := convert.ConvertUserEntityToUserOutput(&user)
		usersOutput[i] = convert.ConvertUserDataToProto(userOutput)
	}

	return &GetAllUsersOutput{Users: usersOutput}
}
