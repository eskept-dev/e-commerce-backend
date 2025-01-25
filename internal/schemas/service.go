package schemas

import "github.com/google/uuid"

type ServiceResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	IsEnabled bool      `json:"isEnabled"`
}
