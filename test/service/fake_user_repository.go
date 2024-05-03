package service

import (
	"errors"
	"product-app/domain"
	"product-app/persistence"
)

type FakeUserRepository struct {
	users []domain.User
}

func (f *FakeUserRepository) AddUser(user *domain.User) error {
	user.ID = uint(len(f.users) + 1)
	f.users = append(f.users, *user)
	return nil
}

func (f *FakeUserRepository) DeleteUserById(userId uint) error {
	for i, user := range f.users {
		if user.ID == userId {
			f.users = append(f.users[:i], f.users[i+1:]...)
			return nil
		}
	}
	return errors.New("User not found")
}

func (f *FakeUserRepository) GetUserById(userId uint) (domain.User, error) {
	for _, user := range f.users {
		if user.ID == userId {
			return user, nil
		}
	}
	return domain.User{}, errors.New("User not found")
}

func (f *FakeUserRepository) GetUserByEmail(email string) (domain.User, error) {
	for _, user := range f.users {
		if user.Email == email {
			return user, nil
		}
	}
	return domain.User{}, errors.New("User not found")
}

func (f *FakeUserRepository) GetAllUsers() []domain.User {
	return f.users
}

func (f *FakeUserRepository) UpdateUser(user *domain.User) error {
	for i, u := range f.users {
		if u.ID == user.ID {
			f.users[i] = *user
			return nil
		}
	}
	return errors.New("User not found")
}

func NewFakeUserRepository(initialUsers []domain.User) persistence.IUserRepository {
	return &FakeUserRepository{
		users: initialUsers,
	}
}
