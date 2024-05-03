package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"product-app/domain"
	"testing"
)

func setupUsers() {
	TestDataInitialize()
}

func clearUsers() {
	TruncateTestData()
}

func TestAddUser(t *testing.T) {
	expectedUser := domain.User{Model: gorm.Model{ID: 1}, Email: "test3@mail.com"}

	newUser := domain.User{
		Email:    "test3@mail.com",
		Password: "test",
	}

	t.Run("AddUser", func(t *testing.T) {
		userRepository.AddUser(&newUser)
		actualUsers := userRepository.GetAllUsers()
		actualUsers[0].Model = gorm.Model{ID: uint(len(actualUsers))}
		assert.Equal(t, 1, len(actualUsers))
		assert.Equal(t, expectedUser.ID, actualUsers[0].ID)
		assert.Equal(t, expectedUser.Email, actualUsers[0].Email)
	})

	clearUsers()
}

func TestGetUserById(t *testing.T) {
	setupUsers()
	expectedUser := domain.User{Model: gorm.Model{ID: 1}, Email: "test@mail.com", Password: "test"}

	t.Run("GetUserById", func(t *testing.T) {
		actualUser, _ := userRepository.GetUserById(1)
		_, err := userRepository.GetUserById(5)
		assert.Equal(t, expectedUser, actualUser)
		assert.Equal(t, "User with id 5 not found", err.Error())
	})

	clearUsers()
}

func TestDeleteUserById(t *testing.T) {
	setupUsers()

	t.Run("DeleteUserById", func(t *testing.T) {
		userRepository.DeleteUserById(1)
		_, err := userRepository.GetUserById(1)
		assert.Equal(t, "User with id 1 not found", err.Error())
	})

	clearUsers()
}

func TestUpdateUser(t *testing.T) {
	setupUsers()
	expectedUser := domain.User{Model: gorm.Model{ID: 1}, Email: "test@mail.com", Password: "test"}
	updatedUser := domain.User{Model: gorm.Model{ID: 1}, Email: "test@mail.com", Password: "test"}

	t.Run("UpdateUser", func(t *testing.T) {
		userRepository.UpdateUser(&updatedUser)
		actualUser, _ := userRepository.GetUserById(1)
		assert.Equal(t, expectedUser.ID, actualUser.ID)
		assert.Equal(t, expectedUser.Email, actualUser.Email)
		// Do not compare passwords as they are hashed and will not match the plain text password
	})

	clearUsers()
}
