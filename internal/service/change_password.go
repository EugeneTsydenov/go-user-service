package service

import (
	"github.com/EugeneTsydenov/go-user-service/pkg"
)

type ChangePasswordOutput struct {
	Code    int32
	Message string
}

func (s *Service) ChangePassword(id int64, newPassword, oldPassword string) ChangePasswordOutput {
	userFromDB, err := s.repo.GetUserById(id)
	if err != nil {
		return ChangePasswordOutput{Code: 401, Message: "User not found, you cant change password"}
	}
	if !pkg.CheckPasswordHash(oldPassword, userFromDB.HashPassword) {
		return ChangePasswordOutput{Code: 401, Message: "Invalid password"}
	}
	hashPassword, err := pkg.HashPassword(newPassword)
	if err != nil {
		return ChangePasswordOutput{Code: 500, Message: "Something Error"}
	}
	err = s.repo.UpdatePassword(id, hashPassword)
	if err != nil {
		return ChangePasswordOutput{Code: 500, Message: "Something Error"}
	}
	return ChangePasswordOutput{Code: 200, Message: "Password successfully updated"}
}
