package models

import (
	"eskept/internal/constants/enums"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	UserTableName = "users"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Status    enums.UserStatus `json:"status" gorm:"default:pending_activation"`
	Role      enums.UserRoles  `json:"role" gorm:"default:user_guest"`
	Email     string           `json:"email" gorm:"uniqueIndex"`
	Password  string           `json:"-"` // Hide password from JSON responses
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
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func isHashed(value string) bool {
	return strings.HasPrefix(value, "$2a$") || strings.HasPrefix(value, "$2b$")
}
