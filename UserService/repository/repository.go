package repository

import (
	"errors"
	"fmt"
	"net/mail"

	"github.com/my-flights/UserService/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) CheckCredentials(username string, password string) (*model.User, error) {
	var user model.User

	repo.db.Table("users").Where("username = ?", username).Find(&user)

	if user.ID == 0 {
		return &user, errors.New("Invalid username!")
	}

	if user.Banned {
		return &user, errors.New("You are banned!")
	}

	if !DoPasswordsMatch(user.Password, password) {
		return &user, errors.New("Invalid password!")
	}

	return &user, nil
}

func HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

func DoPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}

func (repo *Repository) Register(user model.User) (model.User, error) {

	_, err := mail.ParseAddress(user.EmailAddress)

	if err != nil {
		return user, errors.New("Email format is not valid.")
	}

	user.Banned = false
	user.Deactivated = false
	user.Points = 0
	user.Reports = 0
	user.Role = model.USER
	hash, _ := HashPassword(user.Password)

	fmt.Println(DoPasswordsMatch(hash, user.Password))

	user.Password = hash

	createdUser := repo.db.Create(&user)

	if createdUser.Error != nil {
		return user, createdUser.Error
	}

	return user, nil
}
