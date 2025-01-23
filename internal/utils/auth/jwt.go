package jwt

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email     string          `json:"email"`
	Role      string          `json:"role"`
	ExpiresAt jwt.NumericDate `json:"expired_at"`
	jwt.RegisteredClaims
}

func GenerateToken(
	email, role string,
	expirationTime int,
	ctx *context.AppContext,
) (string, error) {
	claims := jwt.MapClaims{
		"email":      email,
		"role":       role,
		"expired_at": jwt.NewNumericDate(time.Now().Add(time.Duration(expirationTime) * time.Second)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ctx.Cfg.JWT.Secret))
}

func ValidateToken(
	tokenString string,
	ctx *context.AppContext,
) (*Claims, error) {

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

	// check if token is expired
	if claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
		return nil, errors.ErrTokenExpired
	}

	return claims, nil
}
