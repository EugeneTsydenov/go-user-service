package service

import (
	"github.com/EugeneTsydenov/go-user-service/pkg"
)

type LoginOutput struct {
	Code    int32
	Message string
	Id      int64
}

func (s *Service) Login(username, password string) LoginOutput {
	userFromDB, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return LoginOutput{Code: 401, Message: "A user with this username does not exist!", Id: 0}
	}

	if !pkg.CheckPasswordHash(password, userFromDB.HashPassword) {
		return LoginOutput{Code: 401, Message: "Invalid password", Id: 0}
	}

	return LoginOutput{Code: 200, Id: userFromDB.ID, Message: "Login success"}
}
