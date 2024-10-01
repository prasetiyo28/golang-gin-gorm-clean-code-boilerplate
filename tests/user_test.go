package tests

import (
	"test1/entities"
	"test1/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockUserRepository struct {
	users []entities.User
}

func (m *mockUserRepository) CreateUser(user entities.User) (entities.User, error) {
	m.users = append(m.users, user)
	return user, nil
}

func TestCreateUser(t *testing.T) {
	mockRepo := &mockUserRepository{}
	userUsecase := usecases.UserUsecase{UserRepo: mockRepo}

	user := entities.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	createdUser, err := userUsecase.CreateUser(user)
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", createdUser.Name)
	assert.Equal(t, "john@example.com", createdUser.Email)
}
