package services

import (
	"eskept/internal/app/context"
	"eskept/internal/models"
	"eskept/internal/repositories"
	"eskept/internal/types"

	"github.com/google/uuid"
)

type ProviderService struct {
	repo   *repositories.ProviderRepository
	appCtx *context.AppContext
}

func NewProviderService(providerRepo *repositories.ProviderRepository, appCtx *context.AppContext) *ProviderService {
	return &ProviderService{
		appCtx: appCtx,
		repo:   providerRepo,
	}
}

func (s *ProviderService) CreateProvider(provider *models.Provider) error {
	return s.repo.Create(provider)
}

func (s *ProviderService) UpdateProvider(provider *models.Provider) error {
	return s.repo.Update(provider)
}

func (s *ProviderService) DeleteProvider(provider *models.Provider) error {
	return s.repo.Delete(provider)
}

func (s *ProviderService) FindByCodeName(codeName string) (*models.Provider, error) {
	return s.repo.FindByCodeName(codeName)
}

func (s *ProviderService) FindByProviderId(providerId uuid.UUID) (*models.Provider, error) {
	return s.repo.FindByProviderId(providerId)
}

func (s *ProviderService) FindAll(keyword *string, pagination *types.Pagination) ([]models.Provider, error) {
	if keyword == nil || *keyword == "" {
		return s.repo.List(pagination)
	}
	return s.repo.Search(*keyword, pagination)
}
