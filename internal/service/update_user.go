package service

import (
	"github.com/EugeneTsydenov/go-user-service/internal/domain/entity"
	"github.com/EugeneTsydenov/go-user-service/internal/proto"
	"reflect"
)

type UpdateUserOutput struct {
	UpdatedUser *entity.UserOutput
	Code        int32
	Message     string
}

func (s *Service) UpdateUser(request *proto.UpdateUserRequest) UpdateUserOutput {
	val := reflect.ValueOf(request.UpdatedField).Elem()
	typ := val.Type()
	updateData := make(map[string]interface{})

	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		fieldName := typ.Field(i).Name
		if fieldName == "sizeCache" || fieldName == "unknownFields" || fieldName == "state" {
			continue
		}
		if fieldValue.Interface() == "" {
			continue
		}
		updateData[fieldName] = fieldValue.Interface()
	}

	updatedData, err := s.repo.UpdateUser(request.GetId(), updateData)
	if err != nil {
		return UpdateUserOutput{Code: 500, Message: "Something error", UpdatedUser: nil}
	}
	return UpdateUserOutput{
		Code:    200,
		Message: "Successfully updated user data",
		UpdatedUser: &entity.UserOutput{
			Id:        updatedData.ID,
			Username:  updatedData.Username,
			Avatar:    updatedData.Avatar,
			CreatedAt: updatedData.CreatedAt,
		},
	}
}
