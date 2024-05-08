package service

type LoginOutput struct {
	success bool
	message bool
	id      int64
}

func (s *Service) Login(username, password string) (*LoginOutput, error) {
	return nil, nil
}
