package repository

import (
	"github.com/EugeneTsydenov/go-user-service/cmd/user-service/db"
	"github.com/EugeneTsydenov/go-user-service/internal/model"
)

func GetUserById(id int64) (model.User, error) {
	user := model.User{}
	result := db.DB.First(&user, "id = ?", id)
	return user, result.Error
}

func GetUserByUsername(username string) (model.User, error) {
	user := model.User{}
	result := db.DB.First(&user, "username = ?", username)
	return user, result.Error
}

func SaveUser(username, hashPassword string) error {
	newUser := model.NewUser{
		Username: username,
		Password: hashPassword,
	}
	tx := db.DB.Table("users").Create(&newUser)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteUser(id int64) error {
	result := db.DB.Delete(&model.User{}, "id = ?", id)
	return result.Error
}
