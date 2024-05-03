package service

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"product-app/domain"
	"product-app/service"
	"product-app/service/model"
	"testing"
)

var userService service.IUserService

func init() {
	initialUsers := []domain.User{
		{Model: gorm.Model{ID: 1}, Email: "test@gmail.com", Password: hashPassword("123456")},
		{Model: gorm.Model{ID: 2}, Email: "test1@gmail.com", Password: hashPassword("123457")},
		{Model: gorm.Model{ID: 3}, Email: "test2@gmail.com", Password: hashPassword("123458")},
		{Model: gorm.Model{ID: 4}, Email: "test3@gmail.com", Password: hashPassword("123459")},
	}
	fakeUserRepository := NewFakeUserRepository(initialUsers)
	userService = service.NewUserService(fakeUserRepository)
}

func TestAddUser(t *testing.T) {
	t.Run("AddUser", func(t *testing.T) {
		err := userService.AddUser(model.UserCreate{
			Email:    "test4@gmail.com",
			Password: "123460",
		})
		assert.Nil(t, err)
		users := userService.GetAllUsers()
		assert.Equal(t, 5, len(users))
		assert.Equal(t, "test4@gmail.com", users[len(users)-1].Email)
	})
}

func TestGetUserById(t *testing.T) {
	t.Run("GetUserById", func(t *testing.T) {
		user, err := userService.GetUserById(1)
		assert.Nil(t, err)
		assert.Equal(t, "test@gmail.com", user.Email)
	})
}

func TestGetUserByEmail(t *testing.T) {
	t.Run("GetUserByEmail", func(t *testing.T) {
		user, err := userService.GetUserByEmail("test@gmail.com")
		assert.Nil(t, err)
		assert.Equal(t, "test@gmail.com", user.Email)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("UpdateUser", func(t *testing.T) {
		err := userService.UpdateUser(model.UserUpdate{
			Id:       1,
			Email:    "updated@gmail.com",
			Password: "123456",
		})
		assert.Nil(t, err)
		user, err := userService.GetUserById(1)
		assert.Nil(t, err)
		assert.Equal(t, "updated@gmail.com", user.Email)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("DeleteUser", func(t *testing.T) {
		err := userService.DeleteUser(1)
		assert.Nil(t, err)
		_, err = userService.GetUserById(1)
		assert.NotNil(t, err)
	})
}

func TestGetAllUsers(t *testing.T) {
	t.Run("GetAllUsers", func(t *testing.T) {
		users := userService.GetAllUsers()
		assert.Equal(t, 4, len(users)) // We deleted one user in the previous test
	})
}
