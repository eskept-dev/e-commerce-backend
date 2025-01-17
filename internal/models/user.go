package models

import "eskept/internal/constants"

const (
	UserTableName = "users"
)

type User struct {
	BaseModel
	DeletedAt int64                `json:"deleted_at"`
	Status    constants.UserStatus `json:"status" gorm:"default:pending_activation"`
	UserRoles constants.UserRoles  `json:"role" gorm:"default:user_guest"`
	Username  string               `json:"username"`
	Password  string               `json:"password"`
}

func (u *User) TableName() string {
	return UserTableName
}
