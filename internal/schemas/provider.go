package schemas

import (
	"time"

	"github.com/google/uuid"
)

type BusinessInformation struct {
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Website     string `json:"website"`
}

type ContactInformation struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Role        string `json:"role"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}

type ProviderCreateRequest struct {
	Name                string              `json:"name" validate:"required"`
	BusinessInformation BusinessInformation `json:"businessInformation" validate:"required"`
	ContactInformation  ContactInformation  `json:"contactInformation" validate:"required"`
}

type ProviderUpdateRequest struct {
	Name                string              `json:"name"`
	IsEnabled           bool                `json:"isEnabled"`
	BusinessInformation BusinessInformation `json:"businessInformation"`
	ContactInformation  ContactInformation  `json:"contactInformation"`
}

type ProviderResponse struct {
	ID                  uuid.UUID           `json:"id"`
	CreatedAt           time.Time           `json:"createdAt"`
	UpdatedAt           time.Time           `json:"updatedAt"`
	Name                string              `json:"name"`
	CodeName            string              `json:"codeName"`
	IsEnabled           bool                `json:"isEnabled"`
	BusinessInformation BusinessInformation `json:"businessInformation"`
	ContactInformation  ContactInformation  `json:"contactInformation"`
}
