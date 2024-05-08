package service

import (
	"github.com/EugeneTsydenov/go-user-service/pkg"
)

type LoginOutput struct {
	Success bool
	Message string
	Id      int64
}

func (s *Service) Login(username, password string) LoginOutput {
	userFromDB, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return LoginOutput{Success: false, Message: "A user with this username does not exist!", Id: -1}
	}

	if !pkg.CheckPasswordHash(password, userFromDB.HashPassword) {
		return LoginOutput{Success: false, Message: "Invalid password", Id: -1}
	}

	return LoginOutput{Success: true, Id: userFromDB.Id, Message: "Login success"}
}
