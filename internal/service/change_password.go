package service

type ChangePasswordOutput struct {
	success bool
	message string
}

func (s *Service) ChangePassword(id int64, oldPassword, newPassword string) (*ChangePasswordOutput, error) {
	return nil, nil
}
