package repository

import (
	"errors"
	"net/mail"

	"github.com/my-flights/UserService/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) Register(user model.User) (model.User, error) {

	_, err := mail.ParseAddress(user.EmailAddress)

	if err != nil {
		return user, errors.New("Email format is not valid.")
	}

	createdUser := repo.db.Create(&user)

	if createdUser.Error != nil {
		return user, createdUser.Error
	}

	return user, nil
}
