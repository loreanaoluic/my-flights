package repository

import (
	"errors"
	"net/http"
	"net/mail"
	"strconv"

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

	user.Password = hash

	createdUser := repo.db.Create(&user)

	if createdUser.Error != nil {
		return user, createdUser.Error
	}

	return user, nil
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page < 0 {
			page = 0
		}

		pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := page * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (repo *Repository) FindAllUsers(r *http.Request) ([]model.User, int64, error) {
	var users []model.User
	var totalElements int64

	result := repo.db.Scopes(Paginate(r)).Table("users").Where("(deleted_at IS NULL)").Find(&users)
	repo.db.Table("users").Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	return users, totalElements, nil
}

func (repo *Repository) FindUserById(id uint) (*model.UserDTO, error) {
	var user model.User
	result := repo.db.Table("users").Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, errors.New("User not found!")
	}

	var retValue model.UserDTO = user.ToUserDTO()
	return &retValue, nil
}

func (repo *Repository) UpdateUser(userDTO *model.UserDTO) (*model.UserDTO, error) {
	var user model.User
	result := repo.db.Table("users").Where("ID = ?", userDTO.Id).First(&user)

	if result.Error != nil {
		return nil, errors.New("User cannot be found!")
	}

	user.FirstName = userDTO.FirstName
	user.LastName = userDTO.LastName
	user.EmailAddress = userDTO.EmailAddress

	result2 := repo.db.Table("users").Save(&user)

	if result2.Error != nil {
		return nil, errors.New("User cannot be updated!")
	}

	var retValue model.UserDTO = user.ToUserDTO()
	return &retValue, nil
}

func (repo *Repository) ActivateAccount(id uint) (*model.UserDTO, error) {
	var user model.User
	result := repo.db.Table("users").Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, errors.New("User not found!")
	}

	user.Deactivated = false

	result2 := repo.db.Table("users").Save(&user)

	if result2.Error != nil {
		return nil, errors.New("Account cannot be activated!")
	}

	var retValue model.UserDTO = user.ToUserDTO()
	return &retValue, nil
}

func (repo *Repository) DeactivateAccount(id uint) (*model.UserDTO, error) {
	var user model.User
	result := repo.db.Table("users").Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, errors.New("User not found!")
	}

	user.Deactivated = true

	result2 := repo.db.Table("users").Save(&user)

	if result2.Error != nil {
		return nil, errors.New("Account cannot be deactivated!")
	}

	var retValue model.UserDTO = user.ToUserDTO()
	return &retValue, nil
}

func (repo *Repository) BanUser(id uint) (*model.UserDTO, error) {
	var user model.User
	result := repo.db.Table("users").Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, errors.New("User not found!")
	}

	user.Banned = true

	result2 := repo.db.Table("users").Save(&user)

	if result2.Error != nil {
		return nil, errors.New("User cannot be banned!")
	}

	var retValue model.UserDTO = user.ToUserDTO()
	return &retValue, nil
}

func (repo *Repository) UnbanUser(id uint) (*model.UserDTO, error) {
	var user model.User
	result := repo.db.Table("users").Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, errors.New("User not found!")
	}

	user.Banned = false

	result2 := repo.db.Table("users").Save(&user)

	if result2.Error != nil {
		return nil, errors.New("User cannot be unbanned!")
	}

	var retValue model.UserDTO = user.ToUserDTO()
	return &retValue, nil
}
