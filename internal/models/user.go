package models

import (
	"eskept/internal/constants"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	UserTableName = "users"
)

type User struct {
	BaseModel
	DeletedAt int64                `json:"deleted_at"`
	Status    constants.UserStatus `json:"status" gorm:"default:pending_activation"`
	UserRoles constants.UserRoles  `json:"role" gorm:"default:user_guest"`
	Email     string               `json:"email"`
	Password  string               `json:"password"`
}

func (u *User) TableName() string {
	return UserTableName
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	return u.HashPassword(u.Password)
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	return u.HashPassword(u.Password)
}

func (u *User) HashPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
