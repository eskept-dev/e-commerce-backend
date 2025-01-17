package services

import (
	"eskept/internal/constants/errors"
	"eskept/internal/models"
	"eskept/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(email, password string) (*models.User, error) {
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

func (s *UserService) Login(email, password string) (*models.User, error) {
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
