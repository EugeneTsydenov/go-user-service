package service

type DeleteUserOutput struct {
	Success bool
	Message string
}

func (s *Service) DeleteUser(id int64) DeleteUserOutput {
	return DeleteUserOutput{}
}
