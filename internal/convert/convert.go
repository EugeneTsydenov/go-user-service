package convert

import (
	"github.com/EugeneTsydenov/go-user-service/internal/domain/entity"
	"github.com/EugeneTsydenov/go-user-service/internal/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertUserDataToProto(userFromDb *entity.UserOutput) *proto.UserData {
	if userFromDb == nil {
		return nil
	}
	return &proto.UserData{
		Id:        userFromDb.Id,
		Username:  userFromDb.Username,
		Avatar:    userFromDb.Avatar,
		CreatedAt: timestamppb.New(userFromDb.CreatedAt),
	}
}

func ConvertUserEntityToUserOutput(userEntity *entity.User) *entity.UserOutput {
	if userEntity == nil {
		return nil
	}

	return &entity.UserOutput{
		Id:        userEntity.ID,
		Username:  userEntity.Username,
		Avatar:    userEntity.Avatar,
		CreatedAt: userEntity.CreatedAt,
	}
}
