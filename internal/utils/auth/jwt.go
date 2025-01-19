package jwt

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email     string          `json:"email"`
	Role      string          `json:"role"`
	ExpiresAt jwt.NumericDate `json:"expired_at"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(
	email, role string,
	ctx *context.AppContext,
) (string, error) {
	claims := jwt.MapClaims{
		"email":      email,
		"role":       role,
		"expired_at": jwt.NewNumericDate(time.Now().Add(time.Duration(ctx.Cfg.JWT.TokenExpirationTime) * time.Second)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ctx.Cfg.JWT.Secret))
}

func GenerateRefreshToken(
	email, role string,
	ctx *context.AppContext,
) (string, error) {
	claims := jwt.MapClaims{
		"email":      email,
		"role":       role,
		"expired_at": jwt.NewNumericDate(time.Now().Add(time.Duration(ctx.Cfg.JWT.RefreshTokenExpirationTime) * time.Second)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ctx.Cfg.JWT.Secret))
}

func GenerateActivationToken(
	email, role string,
	ctx *context.AppContext,
) (string, error) {
	claims := jwt.MapClaims{
		"email":      email,
		"role":       role,
		"expired_at": jwt.NewNumericDate(time.Now().Add(time.Duration(ctx.Cfg.JWT.ActivationTokenExpirationTime) * time.Second)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ctx.Cfg.JWT.Secret))
}

func GenerateAuthenticationToken(
	email, role string,
	ctx *context.AppContext,
) (string, error) {
	claims := jwt.MapClaims{
		"email":      email,
		"role":       role,
		"expired_at": jwt.NewNumericDate(time.Now().Add(time.Duration(ctx.Cfg.JWT.AuthenticationTokenExpirationTime) * time.Second)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ctx.Cfg.JWT.Secret))
}

func ValidateToken(tokenString string, ctx *context.AppContext) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(ctx.Cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	log.Println(claims)
	log.Println(claims.ExpiresAt.Time.Unix(), time.Now().Unix(), claims.ExpiresAt.Time.Unix() < time.Now().Unix())

	// check if token is expired
	if claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
		return nil, errors.ErrTokenExpired
	}

	return claims, nil
}
