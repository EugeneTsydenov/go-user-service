package service

type ChangePasswordOutput struct {
	Success bool
	Message string
}

func (s *Service) ChangePassword(id int64, oldPassword, newPassword string) ChangePasswordOutput {
	return ChangePasswordOutput{}
}
