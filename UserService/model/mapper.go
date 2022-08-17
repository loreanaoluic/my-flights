package model

func (user *User) ToRegisterDTO() RegisterDTO {

	return RegisterDTO{
		Username:     user.Username,
		Password:     user.Password,
		EmailAddress: user.EmailAddress,
	}
}

func (registerDTO *RegisterDTO) ToUser() User {

	return User{
		Username:     registerDTO.Username,
		Password:     registerDTO.Password,
		EmailAddress: registerDTO.EmailAddress,
	}
}
