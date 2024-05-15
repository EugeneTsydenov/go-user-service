package service

import (
	"github.com/EugeneTsydenov/go-user-service/pkg"
)

type UpdatePasswordOutput struct {
	Code    int32
	Message string
}

func (s *Service) UpdatePassword(id int64, newPassword, oldPassword string) UpdatePasswordOutput {
	userFromDB, err := s.repo.GetUserById(id)
	if err != nil {
		return UpdatePasswordOutput{Code: 401, Message: "User not found, you cant change password"}
	}
	if !pkg.CheckPasswordHash(oldPassword, userFromDB.HashPassword) {
		return UpdatePasswordOutput{Code: 401, Message: "Invalid password"}
	}
	hashPassword, err := pkg.HashPassword(newPassword)
	if err != nil {
		return UpdatePasswordOutput{Code: 500, Message: "Something Error"}
	}
	err = s.repo.UpdatePassword(id, hashPassword)
	if err != nil {
		return UpdatePasswordOutput{Code: 500, Message: "Something Error"}
	}
	return UpdatePasswordOutput{Code: 200, Message: "Password successfully updated"}
}
