package repository

import (
	"github.com/EugeneTsydenov/go-user-service/cmd/user-service/db"
	"github.com/EugeneTsydenov/go-user-service/internal/model"
)

func GetUserById(id int64) {

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
