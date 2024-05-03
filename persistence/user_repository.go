package persistence

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"product-app/domain"
	"product-app/initializers"
)

type IUserRepository interface {
	AddUser(user *domain.User) error
	DeleteUserById(userId uint) error
	GetUserById(userId uint) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	GetAllUsers() []domain.User
	UpdateUser(user *domain.User) error
}

type UserRepository struct{}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (userRepository *UserRepository) AddUser(user *domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Error("Error while hashing password: %v\n", err)
		return errors.New("Error while hashing password")
	}

	user.Password = string(hash)
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		log.Error("Error while adding user: %v\n", result.Error)
		return result.Error
	}

	log.Info("User added with id: %v\n", user.ID)

	return nil
}

func (userRepository *UserRepository) DeleteUserById(userId uint) error {
	initializers.DB.Delete(&domain.User{}, userId)

	return nil
}

func (userRepository *UserRepository) GetUserById(userId uint) (domain.User, error) {
	var user domain.User
	result := initializers.DB.First(&user, userId)

	if result.Error != nil {
		log.Error("Error while fetching user: %v\n", result.Error)
		return domain.User{}, errors.New(fmt.Sprintf("User with id %d not found", userId))
	}

	return user, nil
}

func (userRepository *UserRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	result := initializers.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		log.Error("Error while fetching user: %v\n", result.Error)
		return domain.User{}, errors.New(fmt.Sprintf("User with email %s not found", email))
	}

	return user, nil
}

func (userRepository *UserRepository) GetAllUsers() []domain.User {
	var users []domain.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		log.Error("Error while fetching users: %v\n", result.Error)
		return []domain.User{}
	}

	return users
}

func (userRepository *UserRepository) UpdateUser(user *domain.User) error {
	err := userRepository.checkIfUserExists(user.ID)

	if err != nil {
		return err
	}

	tx := initializers.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if user.Email != "" {
		err := userRepository.checkIfEmailExistsExceptUser(user.Email, user.ID)

		if err != nil {
			tx.Rollback()
			return err
		}

		tx.Model(&user).Update("email", user.Email)
	}

	if user.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		if err != nil {
			log.Error("Error while hashing password: %v\n", err)
			tx.Rollback()
			return errors.New("Error while hashing password")
		}

		user.Password = string(hash)
		tx.Model(&user).Update("password", user.Password)

	}

	tx.Commit()

	log.Info("User updated with id: %v\n", user.ID)

	return nil
}

func (userRepository *UserRepository) checkIfEmailExists(email string) bool {
	var user domain.User
	result := initializers.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return false
	}

	return true
}

func (userRepository *UserRepository) checkIfEmailExistsExceptUser(email string, userId uint) error {
	var user domain.User
	result := initializers.DB.Where("email = ? AND id != ?", email, userId).First(&user)

	if result.Error != nil {
		return nil
	}

	return errors.New("Email already exists")
}

func (userRepository *UserRepository) checkIfUserExists(userId uint) error {
	_, getErr := userRepository.GetUserById(userId)

	if getErr != nil {
		log.Error("Error while checking if user exists: %v\n", getErr)
		return errors.New(fmt.Sprintf("User with id %d not found", userId))
	}

	return nil
}
