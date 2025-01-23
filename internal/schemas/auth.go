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
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AuthLoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthSendActivationRequest struct {
	Email string `json:"email" binding:"required"`
}

type AuthSendActivationResponse struct {
	IsSuccess bool `json:"isSuccess"`
}

type AuthActivateRequest struct {
	Token string `json:"token" binding:"required"`
}

type AuthActivateResponse struct {
	IsActivated bool `json:"isActivated"`
}

type AuthSendAuthenticationRequest struct {
	Email string `json:"email" binding:"required"`
}

type AuthSendAuthenticationResponse struct {
	IsSuccess bool `json:"isSuccess"`
}

type AuthLoginByAuthenticationTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

type AuthLoginByAuthenticationTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthSendVerificationEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

type AuthSendVerificationEmailResponse struct {
	IsSuccess bool `json:"isSuccess"`
}

type AuthVerifyEmailTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

type AuthVerifyEmailTokenResponse struct {
	IsVerified bool `json:"isVerified"`
}
