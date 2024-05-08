package service

import (
	"github.com/EugeneTsydenov/go-user-service/pkg"
)

type ChangePasswordOutput struct {
	Success bool
	Message string
}

func (s *Service) ChangePassword(id int64, newPassword, oldPassword string) ChangePasswordOutput {
	userFromDB, err := s.repo.GetUserById(id)
	if err != nil {
		return ChangePasswordOutput{Success: false, Message: "User not found, you cant change password"}
	}
	if !pkg.CheckPasswordHash(oldPassword, userFromDB.HashPassword) {
		return ChangePasswordOutput{Success: false, Message: "Invalid password"}
	}
	hashPassword, err := pkg.HashPassword(newPassword)
	if err != nil {
		return ChangePasswordOutput{Success: false, Message: "Something Error"}
	}
	err = s.repo.UpdatePassword(id, hashPassword)
	if err != nil {
		return ChangePasswordOutput{Success: false, Message: "Something Error"}
	}
	return ChangePasswordOutput{Success: true, Message: "Password successfully updated"}
}
