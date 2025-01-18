package services

import (
	"eskept/internal/app/context"
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

func (s *AuthService) Login(email, password string) (types.TokenPair, error) {
	// Find user by email
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return types.TokenPair{}, err
	}

	// Check password
	if !user.ComparePassword(password) {
		return types.TokenPair{}, errors.ErrInvalidCredentials
	}

	accessToken, err := jwt.GenerateAccessToken(email, string(user.UserRoles), s.appCtx)
	if err != nil {
		return types.TokenPair{}, err
	}

	refreshToken, err := jwt.GenerateRefreshToken(email, string(user.UserRoles), s.appCtx)
	if err != nil {
		return types.TokenPair{}, err
	}

	return types.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
