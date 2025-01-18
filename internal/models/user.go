package models

import (
	"eskept/internal/constants/enums"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	UserTableName = "users"
)

type User struct {
	BaseModel
	DeletedAt int64            `json:"deleted_at"`
	Status    enums.UserStatus `json:"status" gorm:"default:pending_activation"`
	Role      enums.UserRoles  `json:"role" gorm:"default:user_guest"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
}

func (u *User) TableName() string {
	return UserTableName
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Password != "" {
		u.Password = u.HashPassword(u.Password)
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if u.Password != "" {
		if !isHashed(u.Password) {
			u.Password = u.HashPassword(u.Password)
		}
	}
	return nil
}

func (u *User) ComparePassword(password string) bool {
	log.Println(u.Password, u.HashPassword(password), bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(u.HashPassword(password))))
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) HashPassword(password string) string {
	log.Println(password)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func isHashed(value string) bool {
	return strings.HasPrefix(value, "$2a$") || strings.HasPrefix(value, "$2b$")
}
