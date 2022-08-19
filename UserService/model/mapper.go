package model

import "gorm.io/gorm"

func (user *User) ToRegisterDTO() RegisterDTO {

	return RegisterDTO{
		Username:     user.Username,
		Password:     user.Password,
		EmailAddress: user.EmailAddress,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
	}
}

func (registerDTO *RegisterDTO) ToUser() User {

	return User{
		Model:        gorm.Model{},
		Username:     registerDTO.Username,
		Password:     registerDTO.Password,
		EmailAddress: registerDTO.EmailAddress,
		FirstName:    registerDTO.FirstName,
		LastName:     registerDTO.LastName,
	}
}

func (user *User) ToUserDTO() UserDTO {

	return UserDTO{
		Id:           user.ID,
		Username:     user.Username,
		EmailAddress: user.EmailAddress,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
	}
}

func (userDTO *UserDTO) ToUser() User {

	return User{
		Model:        gorm.Model{},
		Username:     userDTO.Username,
		Password:     userDTO.Password,
		EmailAddress: userDTO.EmailAddress,
		FirstName:    userDTO.FirstName,
		LastName:     userDTO.LastName,
	}
}
