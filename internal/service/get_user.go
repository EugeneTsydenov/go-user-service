package service

import (
	"fmt"
	"github.com/EugeneTsydenov/go-user-service/internal/domain/entity"
)

type GetUserOutput struct {
	UserData *entity.UserOutput
	Message  string
	Success  bool
}

func (s *Service) GetUser(id int64) GetUserOutput {
	fmt.Println(id, "idshnik")
	return GetUserOutput{}
}
