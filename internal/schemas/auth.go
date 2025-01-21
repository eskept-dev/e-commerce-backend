package schemas

import (
	"eskept/internal/constants/enums"
	"time"

	"github.com/google/uuid"
)

type AuthRegisterRequest struct {
	Email    string          `json:"email" binding:"required"`
	Password string          `json:"password" binding:"required"`
	Role     enums.UserRoles `json:"role" binding:"required"`
}

type AuthRegisterResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthSendActivationRequest struct {
	Email string `json:"email" binding:"required"`
}

type AuthSendActivationResponse struct {
	IsSuccess bool `json:"is_success"`
}

type AuthActivateRequest struct {
	Token string `json:"token" binding:"required"`
}

type AuthActivateResponse struct {
	IsActivated bool `json:"is_activated"`
}

type AuthSendAuthenticationRequest struct {
	Email string `json:"email" binding:"required"`
}

type AuthSendAuthenticationResponse struct {
	IsSuccess bool `json:"is_success"`
}

type AuthLoginByAuthenticationTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

type AuthLoginByAuthenticationTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthSendVerificationEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

type AuthSendVerificationEmailResponse struct {
	IsSuccess bool `json:"is_success"`
}

type AuthVerifyEmailTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

type AuthVerifyEmailTokenResponse struct {
	IsVerified bool `json:"is_verified"`
}
