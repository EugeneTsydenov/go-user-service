package service

type DeleteUserOutput struct {
	Code    int32
	Message string
}

func (s *Service) DeleteUser(id int64) DeleteUserOutput {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return DeleteUserOutput{Code: 400, Message: "User cant be delete"}
	}
	return DeleteUserOutput{Code: 200, Message: "User successfully deleted"}
}
