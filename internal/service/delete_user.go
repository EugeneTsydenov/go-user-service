package service

type DeleteUserOutput struct {
	success bool
	message string
}

func (s *Service) DeleteUser(id int64) (*DeleteUserOutput, error) {
	return nil, nil
}
