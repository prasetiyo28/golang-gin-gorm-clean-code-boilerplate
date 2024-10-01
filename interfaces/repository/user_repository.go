package repository

import (
	"test1/entities"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func (r *GormUserRepository) CreateUser(user entities.User) (entities.User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *GormUserRepository) GetUser(id string) (*entities.User, error) {
	var user entities.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
