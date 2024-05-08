package service

import "fmt"

type RegisterOutput struct {
	Success bool
	Message string
}

func (s *Service) Register(username, password string) (*RegisterOutput, error) {
	fmt.Println(s, "service")
	err := s.repo.SaveUser(username, password)
	if err != nil {
		return &RegisterOutput{false, "Not registered"}, nil
	}
	return &RegisterOutput{true, "Registered"}, nil
}
