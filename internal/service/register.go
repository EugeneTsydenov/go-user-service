package service

import "github.com/EugeneTsydenov/go-user-service/pkg"

type RegisterOutput struct {
	Success bool
	Message string
}

func (s *Service) Register(username, password string) RegisterOutput {
	_, err := s.repo.GetUserByUsername(username)
	if err == nil {
		return RegisterOutput{false, "User with this username is already exist!"}
	}
	hashPassword, err := pkg.HashPassword(password)
	if err != nil {
		return RegisterOutput{false, "Something error"}
	}
	err = s.repo.SaveUser(username, hashPassword)
	if err != nil {
		return RegisterOutput{false, "SomethingError"}
	}
	return RegisterOutput{true, "User successfully saved"}
}
