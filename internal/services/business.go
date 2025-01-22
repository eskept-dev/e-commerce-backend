package services

import (
	"eskept/internal/app/context"
	"eskept/internal/models"
	"eskept/internal/repositories"

	"github.com/google/uuid"
)

type BusinessService struct {
	businessProfileRepo *repositories.BusinessProfileRepository
	userRepo            *repositories.UserRepository
	appCtx              *context.AppContext
}

func NewBusinessService(
	businessProfileRepo *repositories.BusinessProfileRepository,
	userRepo *repositories.UserRepository,
	appCtx *context.AppContext,
) *BusinessService {
	return &BusinessService{
		businessProfileRepo: businessProfileRepo,
		userRepo:            userRepo,
		appCtx:              appCtx,
	}
}

func (s *BusinessService) GetProfile(id uuid.UUID) (*models.BusinessProfile, error) {
	return s.businessProfileRepo.FindByBusinessProfileId(id)
}

func (s *BusinessService) CreateProfile(profile *models.BusinessProfile) error {
	return s.businessProfileRepo.Create(profile)
}
