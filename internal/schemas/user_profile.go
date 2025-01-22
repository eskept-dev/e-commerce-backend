package schemas

import (
	"time"

	"github.com/google/uuid"
)

type UserProfileCreateRequest struct {
	FirstName   string     `json:"firstName" binding:"required"`
	LastName    string     `json:"lastName" binding:"required"`
	DateOfBirth *time.Time `json:"dateOfBirth"`
	Sex         string     `json:"sex"`
	Nationality string     `json:"nationality"`
	DialCode    string     `json:"dialCode"`
	PhoneNumber string     `json:"phoneNumber"`
}

type UserProfileCreateResponse struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	UserId      uuid.UUID  `json:"userId"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	DateOfBirth *time.Time `json:"dateOfBirth"`
	Sex         string     `json:"sex"`
	Nationality string     `json:"nationality"`
	DialCode    string     `json:"dialCode"`
	PhoneNumber string     `json:"phoneNumber"`
	Email       string     `json:"email"`
}

type UserProfileUpdateRequest struct {
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	DateOfBirth *time.Time `json:"dateOfBirth"`
	Sex         string     `json:"sex"`
	Nationality string     `json:"nationality"`
	DialCode    string     `json:"dialCode"`
	PhoneNumber string     `json:"phoneNumber"`
}

type UserProfileResponse struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	UserId      uuid.UUID  `json:"userId"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	DateOfBirth *time.Time `json:"dateOfBirth"`
	Sex         string     `json:"sex"`
	Nationality string     `json:"nationality"`
	DialCode    string     `json:"dialCode"`
	PhoneNumber string     `json:"phoneNumber"`
	Email       string     `json:"email"`
}
