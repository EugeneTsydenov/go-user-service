package service

type DeleteUserOutput struct {
	Success bool
	Message string
}

func (s *Service) DeleteUser(id int64) DeleteUserOutput {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return DeleteUserOutput{Success: false, Message: "User cant be delete"}
	}
	return DeleteUserOutput{Success: true, Message: "User successfully deleted"}
}
