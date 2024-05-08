package service

import "github.com/EugeneTsydenov/go-user-service/internal/proto"

type UpdateUserOutput struct {
	success bool
	message string
}

func (s *Service) UpdateUser(request *proto.UpdateUserRequest) (*UpdateUserOutput, error) {
	return nil, nil
}
