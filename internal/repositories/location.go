package repositories

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/enums"
	"eskept/internal/models"
	"eskept/internal/types"

	"gorm.io/gorm"
)

type LocationRepository struct {
	db     *gorm.DB
	appCtx *context.AppContext
}

func NewLocationRepository(appCtx *context.AppContext) *LocationRepository {
	return &LocationRepository{
		db:     appCtx.DB,
		appCtx: appCtx,
	}
}

func (r *LocationRepository) Create(location *models.Location) error {
	return r.db.Create(location).Error
}

func (r *LocationRepository) Update(location *models.Location) error {
	return r.db.Save(location).Error
}

func (r *LocationRepository) Delete(location *models.Location) error {
	return r.db.Delete(location).Error
}

func (r *LocationRepository) List(
	locationType enums.LocationType,
	keyword string,
	pagination *types.Pagination,
) ([]*models.Location, error) {
	query := r.db.Model(&models.Location{})

	if locationType != "" {
		query = query.Where("type = ?", locationType)
	}

	if keyword != "" {
		searchValue := GenerateSearchValue(keyword)
		query = query.Where("code_name LIKE ?", searchValue)
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	pagination.Total = total

	// Apply pagination
	offset := (pagination.Page - 1) * pagination.PageSize
	var locations []*models.Location
	err := query.Order("COALESCE(priority, 0) ASC").
		Offset(offset).
		Limit(pagination.PageSize).
		Find(&locations).Error
	if err != nil {
		return nil, err
	}
	return locations, nil
}
