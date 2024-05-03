package service

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"product-app/domain"
	"product-app/persistence"
	"product-app/service"
	"product-app/service/model"
	"testing"
)

var authService service.IAuthService
var fakeUserRepository persistence.IUserRepository

func init() {
	initialUsers := []domain.User{
		{Model: gorm.Model{ID: 1}, Email: "test@gmail.com", Password: hashPassword("123456")},
		{Model: gorm.Model{ID: 2}, Email: "test1@gmail.com", Password: hashPassword("123457")},
		{Model: gorm.Model{ID: 3}, Email: "test2@gmail.com", Password: hashPassword("123458")},
		{Model: gorm.Model{ID: 4}, Email: "test3@gmail.com", Password: hashPassword("123459")},
	}
	fakeUserRepository = NewFakeUserRepository(initialUsers)
	authService = service.NewAuthService(fakeUserRepository)
}

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func TestRegister(t *testing.T) {
	t.Run("Register", func(t *testing.T) {
		token, err := authService.Register(model.UserCreate{
			Email:    "test5@gmail.com",
			Password: "123461",
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		users := fakeUserRepository.GetAllUsers()
		assert.Equal(t, 5, len(users))
		assert.Equal(t, "test5@gmail.com", users[len(users)-1].Email)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login", func(t *testing.T) {
		token, err := authService.Login(model.UserLogin{
			Email:    "test@gmail.com",
			Password: "123456",
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
	})
}
