package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"not null;unique"`
	Password     string `gorm:"not null"`
	EmailAddress string `gorm:"not null"`
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	Role         UserRole
	Banned       bool
	Deactivated  bool
	Reports      uint
	Points       uint
}
