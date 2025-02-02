package services

import (
	"eskept/internal/app/context"
	"eskept/internal/models"
	"eskept/internal/repositories"

	"github.com/google/uuid"
)

type UserService struct {
	repo        *repositories.UserRepository
	profileRepo *repositories.UserProfileRepository
	appCtx      *context.AppContext
}

func NewUserService(
	repo *repositories.UserRepository,
	profileRepo *repositories.UserProfileRepository,
	appCtx *context.AppContext,
) *UserService {
	return &UserService{repo: repo, profileRepo: profileRepo, appCtx: appCtx}
}

func (s *UserService) CreateProfile(profile *models.UserProfile) error {
	return s.profileRepo.Create(profile)
}

func (s *UserService) GetProfileByUserId(userId uuid.UUID) (*models.UserProfile, error) {
	return s.profileRepo.FindByUserId(userId)
}
