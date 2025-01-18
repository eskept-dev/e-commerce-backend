package services

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/errors"
	"eskept/internal/models"
	"eskept/internal/repositories"
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

func (s *AuthService) Login(email, password string) (*models.User, error) {
	// Find user by email
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	// Check password
	if !user.ComparePassword(password) {
		return nil, errors.ErrInvalidCredentials
	}

	return user, nil
}
