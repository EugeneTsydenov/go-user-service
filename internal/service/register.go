package service

import (
	"fmt"
	"github.com/EugeneTsydenov/go-user-service/pkg"
)

type RegisterOutput struct {
	Code    int32
	Message string
}

func (s *Service) Register(username, password string) RegisterOutput {
	_, err := s.repo.GetUserByUsername(username)
	if err == nil {
		return RegisterOutput{401, "User with this username is already exist!"}
	}
	hashPassword, err := pkg.HashPassword(password)
	if err != nil {
		return RegisterOutput{500, "Something Error"}
	}
	err = s.repo.SaveUser(username, hashPassword)
	fmt.Println(err)
	if err != nil {
		return RegisterOutput{500, "Something Error"}
	}
	return RegisterOutput{201, "User successfully saved"}
}
