package jwt

import (
	"eskept/internal/app/context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(
	email, role string,
	ctx *context.AppContext,
) (string, error) {
	claims := Claims{
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(time.Duration(ctx.Cfg.JWT.TokenExpirationTime)),
			)},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ctx.Cfg.JWT.Secret))
}

func GenerateRefreshToken(
	email, role string,
	ctx *context.AppContext,
) (string, error) {
	claims := Claims{
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(time.Duration(ctx.Cfg.JWT.RefreshTokenExpirationTime)),
			)},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ctx.Cfg.JWT.Secret))
}
