package service

import (
	"fmt"
)

func (s *Service) GetUser(id int64) error {
	fmt.Println(id, "idshnik")
	return nil
}
