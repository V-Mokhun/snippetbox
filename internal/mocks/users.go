package mocks

import (
	"time"

	"github.com/v-mokhun/snippetbox/internal/models"
)

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "exists@gmail.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "exists@gmail.com" && password == "password" {
		return 1, nil
	}

	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	if id == 1 {
		return true, nil
	}

	return false, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	if id == 1 {
		return &models.User{
			ID:      1,
			Name:    "Bob",
			Email:   "exists@gmail.com",
			Created: time.Now(),
		}, nil
	}

	return nil, models.ErrNoRecord
}

func (m *UserModel) PasswordUpdate(id int, currentPassword, newPassword string) error {
	if id == 1 {
		if currentPassword != "password" {
			return models.ErrInvalidCredentials
		}

		return nil
	}

	return models.ErrNoRecord
}
