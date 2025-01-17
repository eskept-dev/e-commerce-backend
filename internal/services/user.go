package services

import (
	"eskept/internal/errors"
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
