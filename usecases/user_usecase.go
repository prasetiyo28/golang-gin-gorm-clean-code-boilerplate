package usecases

import (
	"test1/entities"
)

type UserRepository interface {
	CreateUser(user entities.User) (entities.User, error)
	GetUser(id string) (*entities.User, error)
}

type UserUsecase struct {
	UserRepo UserRepository
}

func (uc *UserUsecase) CreateUser(user entities.User) (entities.User, error) {
	return uc.UserRepo.CreateUser(user)
}

func (uc *UserUsecase) GetUser(id string) (*entities.User, error) {
	return uc.UserRepo.GetUser(id)
}
