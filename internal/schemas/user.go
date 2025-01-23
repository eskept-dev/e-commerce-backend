package schemas

import (
	"time"

	"github.com/google/uuid"
)

type UserGetMeResponse struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
	Status    string     `json:"status"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
}
