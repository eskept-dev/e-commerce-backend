package models

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	BaseModel

	User User `gorm:"foreignKey:UserId" json:"-"`

	FirstName   string     `gorm:"type:varchar(255)" json:"firstName"`
	LastName    string     `gorm:"type:varchar(255)" json:"lastName"`
	DateOfBirth *time.Time `gorm:"type:timestamptz" json:"dateOfBirth"`
	Sex         string     `gorm:"type:varchar(255)" json:"sex"`
	Nationality string     `gorm:"type:varchar(255)" json:"nationality"`
	DialCode    string     `gorm:"type:varchar(255)" json:"dialCode"`
	PhoneNumber string     `gorm:"type:varchar(255)" json:"phoneNumber"`
	Email       string     `gorm:"type:varchar(255)" json:"email"`
	UserId      uuid.UUID  `gorm:"type:uuid;not null" json:"userId"`
}

func (UserProfile) TableName() string {
	return "user_profiles"
}
