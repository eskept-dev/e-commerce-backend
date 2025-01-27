package services

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/enums"
	"eskept/internal/models"
	"eskept/internal/repositories"
	"eskept/internal/types"
)

type LocationService struct {
	repo   *repositories.LocationRepository
	appCtx *context.AppContext
}

func NewLocationService(repo *repositories.LocationRepository, appCtx *context.AppContext) *LocationService {
	return &LocationService{repo: repo, appCtx: appCtx}
}

func (s *LocationService) Create(location *models.Location) error {
	return s.repo.Create(location)
}

func (s *LocationService) Update(location *models.Location) error {
	return s.repo.Update(location)
}

func (s *LocationService) Delete(location *models.Location) error {
	return s.repo.Delete(location)
}

func (s *LocationService) List(keyword string, locationType enums.LocationType, pagination *types.Pagination) ([]*models.Location, error) {
	return s.repo.List(locationType, keyword, pagination)
}
