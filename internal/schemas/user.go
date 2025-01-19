package schemas

import (
	"time"

	"github.com/google/uuid"
)

type UserGetMeResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int64     `json:"deleted_at"`
	Status    string    `json:"status"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
}
