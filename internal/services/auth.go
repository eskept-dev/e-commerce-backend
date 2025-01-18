package services

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/enums"
	"eskept/internal/constants/errors"
	"eskept/internal/models"
	"eskept/internal/repositories"
	"eskept/internal/types"
	jwt "eskept/internal/utils/auth"
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

func (s *AuthService) Register(email, password string) (*models.User, error) {
	// Check if user already exists
	isEmailExists := s.repo.IsEmailExists(email)
	if isEmailExists {
		return nil, errors.ErrEmailExists
	}

	// Create new user
	user := &models.User{
		Email:    email,
		Password: password,
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
	accessToken, err := jwt.GenerateAccessToken(email, role, s.appCtx)
	if err != nil {
		return types.TokenPair{}, err
	}

	refreshToken, err := jwt.GenerateRefreshToken(email, role, s.appCtx)
	if err != nil {
		return types.TokenPair{}, err
	}

	return types.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) GenerateActivationLink(email, role string) (string, error) {
	activationToken, err := jwt.GenerateActivationToken(email, role, s.appCtx)
	if err != nil {
		return "", err
	}

	activationLink := s.appCtx.Cfg.App.ActivationURL + "?activationToken=" + activationToken
	return activationLink, nil
}

func (s *AuthService) ActivateUser(activationToken string) error {
	claims, err := jwt.ValidateToken(activationToken, s.appCtx)
	if err != nil {
		return err
	}

	user, err := s.repo.FindByEmail(claims.Email)
	if err != nil {
		return err
	}

	user.Status = enums.UserStatusEnabled
	return s.repo.Update(user)
}
