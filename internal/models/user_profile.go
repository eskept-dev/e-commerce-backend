package models

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	BaseModel

	User User `gorm:"foreignKey:UserId" json:"-"`

	FirstName   string     `gorm:"type:varchar(255)" json:"first_name"`
	LastName    string     `gorm:"type:varchar(255)" json:"last_name"`
	DateOfBirth *time.Time `gorm:"type:timestamptz" json:"date_of_birth"`
	Sex         string     `gorm:"type:varchar(255)" json:"sex"`
	Nationality string     `gorm:"type:varchar(255)" json:"nationality"`
	DialCode    string     `gorm:"type:varchar(255)" json:"dial_code"`
	PhoneNumber string     `gorm:"type:varchar(255)" json:"phone_number"`
	Email       string     `gorm:"type:varchar(255)" json:"email"`
	UserId      uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
}

func (UserProfile) TableName() string {
	return "user_profiles"
}
