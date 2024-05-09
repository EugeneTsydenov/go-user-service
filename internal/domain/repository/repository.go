package repository

import (
	"github.com/EugeneTsydenov/go-user-service/internal/domain/entity"
	"gorm.io/gorm"
)

type RepoInterface interface {
	GetUserById(id int64) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	SaveUser(username, hashPassword string) error
	DeleteUser(id int64) error
	UpdatePassword(id int64, newPassword string) error
	UpdateUser(userID int64, updateData map[string]interface{}) (*entity.User, error)
}

var _ RepoInterface = (*Repository)(nil)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) GetUserById(id int64) (entity.User, error) {
	user := entity.User{}
	result := repo.db.First(&user, "id = ?", id)
	return user, result.Error
}

func (repo *Repository) GetUserByUsername(username string) (entity.User, error) {
	user := entity.User{}
	result := repo.db.First(&user, "username = ?", username)
	return user, result.Error
}

func (repo *Repository) SaveUser(username, hashPassword string) error {
	newUser := entity.NewUser{
		Username: username,
		Password: hashPassword,
	}
	tx := repo.db.Table("users").Create(&newUser)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *Repository) DeleteUser(id int64) error {
	result := repo.db.Delete(&entity.User{}, "id = ?", id)
	return result.Error
}

func (repo *Repository) UpdatePassword(id int64, hashPassword string) error {
	result := repo.db.Table("users").Where("id = ?", id).Update("hash_password", hashPassword)
	return result.Error
}

func (repo *Repository) UpdateUser(userID int64, updateData map[string]interface{}) (*entity.User, error) {
	updatedUser := entity.User{}
	result := repo.db.Model(&updatedUser).Where("id = ?", userID).Updates(updateData).First(&updatedUser)
	return &updatedUser, result.Error
}
