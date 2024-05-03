package service

import (
	"gorm.io/gorm"
	"product-app/domain"
	"product-app/persistence"
	"product-app/service/model"
)

type IUserService interface {
	AddUser(userCreate model.UserCreate) error
	UpdateUser(userUpdate model.UserUpdate) error
	DeleteUser(userId uint) error
	GetAllUsers() []domain.User
	GetUserById(userId uint) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
}

type UserService struct {
	userRepository persistence.IUserRepository
}

func (userService UserService) AddUser(userCreate model.UserCreate) error {
	return userService.userRepository.AddUser(&domain.User{
		Email:    userCreate.Email,
		Password: userCreate.Password,
	})
}

func (userService UserService) UpdateUser(userUpdate model.UserUpdate) error {
	return userService.userRepository.UpdateUser(&domain.User{
		Model:    gorm.Model{ID: userUpdate.Id},
		Email:    userUpdate.Email,
		Password: userUpdate.Password,
	})
}

func (userService UserService) DeleteUser(userId uint) error {
	return userService.userRepository.DeleteUserById(userId)
}

func (userService UserService) GetAllUsers() []domain.User {
	return userService.userRepository.GetAllUsers()
}

func (userService UserService) GetUserById(userId uint) (domain.User, error) {
	return userService.userRepository.GetUserById(userId)
}

func (userService UserService) GetUserByEmail(email string) (domain.User, error) {
	return userService.userRepository.GetUserByEmail(email)
}

func NewUserService(userRepository persistence.IUserRepository) IUserService {
	return &UserService{userRepository: userRepository}
}
