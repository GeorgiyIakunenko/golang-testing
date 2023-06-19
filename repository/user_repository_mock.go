package repository

import (
	"auth/model"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryMock struct {
	users []*model.User
}

func NewUserRepositoryMock() IUserRepository {
	p1, _ := bcrypt.GenerateFromPassword([]byte("test_passw"), bcrypt.DefaultCost)
	p2, _ := bcrypt.GenerateFromPassword([]byte("test_passw_2"), bcrypt.DefaultCost)

	users := []*model.User{
		&model.User{
			ID:       1,
			Email:    "test-1@example.com",
			Name:     "Alex",
			Password: string(p1),
		},
		&model.User{
			ID:       2,
			Email:    "test-2@example.com",
			Name:     "Mary",
			Password: string(p2),
		},
	}

	return &UserRepository{users: users}
}

func (r *UserRepositoryMock) GetUserByEmail(email string) (*model.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserRepositoryMock) GetUserByID(id int) (*model.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}
