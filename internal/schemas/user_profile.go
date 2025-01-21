package schemas

import (
	"time"

	"github.com/google/uuid"
)

type UserProfileCreateRequest struct {
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Sex         string     `json:"sex"`
	Nationality string     `json:"nationality"`
	DialCode    string     `json:"dial_code"`
	PhoneNumber string     `json:"phone_number"`
}

type UserProfileCreateResponse struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UserId      uuid.UUID  `json:"user_id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	Sex         string     `json:"sex"`
	Nationality string     `json:"nationality"`
	DialCode    string     `json:"dial_code"`
	PhoneNumber string     `json:"phone_number"`
	Email       string     `json:"email"`
}
