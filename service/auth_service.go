package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"product-app/domain"
	"product-app/persistence"
	"product-app/service/model"
	"time"
)

type IAuthService interface {
	Login(userLogin model.UserLogin) (string, error)
	Register(userCreate model.UserCreate) (string, error)
}

type AuthService struct {
	userRepository persistence.IUserRepository
}

func NewAuthService(userRepository persistence.IUserRepository) IAuthService {
	return &AuthService{userRepository: userRepository}
}

func (authService AuthService) Register(userCreate model.UserCreate) (string, error) {
	user := &domain.User{
		Email:    userCreate.Email,
		Password: userCreate.Password,
	}

	authService.userRepository.AddUser(user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    userCreate.Email,
		"password": userCreate.Password,
		"sub":      user.ID,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", errors.New("Error while generating token")
	}

	return tokenString, nil
}

func (authService AuthService) Login(userLogin model.UserLogin) (string, error) {
	user, err := authService.userRepository.GetUserByEmail(userLogin.Email)

	if err != nil {
		return "", errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password))

	if err != nil {
		return "", errors.New("Invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    userLogin.Email,
		"password": userLogin.Password,
		"sub":      user.ID,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", errors.New("Error while generating token")
	}

	return tokenString, nil
}
