package models

import (
	"github.com/Kawar1mi/go-jwt/internals/db"
	"gorm.io/gorm"
)

func init() {
	db.Connect()
	db.DB.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

func (u *User) CreateUser() error {
	return db.DB.Create(u).Error
}

func (u *User) GetUserByEmail() error {
	return db.DB.Where("email = ?", u.Email).First(u).Error
}

func (u *User) GetUserByID(id string) error {
	return db.DB.First(u, id).Error
}
