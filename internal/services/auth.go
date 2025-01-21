package services

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/enums"
	"eskept/internal/constants/errors"
	"eskept/internal/models"
	"eskept/internal/repositories"
	"eskept/internal/types"
	jwt "eskept/internal/utils/auth"
	"log"
	"time"
)

type AuthService struct {
	repo   *repositories.UserRepository
	appCtx *context.AppContext
}

func NewAuthService(
	repo *repositories.UserRepository,
	appCtx *context.AppContext,
) *AuthService {
	return &AuthService{repo: repo, appCtx: appCtx}
}

func (s *AuthService) Register(
	email, password string,
	role enums.UserRoles,
	status enums.UserStatus,
) (*models.User, error) {
	isEmailExists := s.repo.IsEmailExists(email)
	if isEmailExists {
		return nil, errors.ErrEmailExists
	}

	user := &models.User{
		Email:    email,
		Password: password,
		Role:     role,
		Status:   status,
	}
	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) IsAuthenticated(email, password string) (bool, error) {
	// Find user by email
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return false, err
	}

	// Check if user exists
	if user == nil {
		return false, errors.ErrUserNotFound
	}

	// Check password
	if !user.ComparePassword(password) {
		return false, errors.ErrInvalidCredentials
	}

	// Check activation
	if user.Status != enums.UserStatusEnabled {
		return false, errors.ErrUserNotEnabled
	}
	return true, nil
}

func (s *AuthService) GenerateTokens(email, role string) (types.TokenPair, error) {
	accessToken, err := jwt.GenerateToken(email, role, s.appCtx.Cfg.JWT.TokenExpirationTime, s.appCtx)
	if err != nil {
		return types.TokenPair{}, err
	}

	refreshToken, err := jwt.GenerateToken(email, role, s.appCtx.Cfg.JWT.RefreshTokenExpirationTime, s.appCtx)
	if err != nil {
		return types.TokenPair{}, err
	}

	return types.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) ActivateUserByActivationToken(activationToken string) error {
	claims, err := jwt.ValidateToken(activationToken, s.appCtx)
	if err != nil {
		return err
	}

	user, err := s.repo.FindByEmail(claims.Email)
	if err != nil {
		return err
	}

	return s.repo.Activate(user)
}

func (s *AuthService) LoginByAuthenticationToken(authenticationToken string) (types.TokenPair, error) {
	claims, err := jwt.ValidateToken(authenticationToken, s.appCtx)
	if err != nil {
		return types.TokenPair{}, err
	}

	log.Println("------------------- Login by authentication token -------------------")
	log.Println("Email:", claims.Email)
	log.Println("Role:", claims.Role)
	log.Println("ExpiresAt:", claims.ExpiresAt.Time.Unix())
	log.Println("Comparison:", time.Now().Unix(), claims.ExpiresAt.Time.Unix() < time.Now().Unix())
	log.Println("------------------------------------------------------------")

	user, err := s.repo.FindByEmail(claims.Email)
	if err != nil {
		return types.TokenPair{}, err
	}

	if claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
		return types.TokenPair{}, errors.ErrTokenExpired
	}

	return s.GenerateTokens(user.Email, string(user.Role))
}

func (s *AuthService) VerifyEmailToken(token string) error {
	_, err := jwt.ValidateToken(token, s.appCtx)
	return err
}
