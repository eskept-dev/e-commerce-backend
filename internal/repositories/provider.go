package repositories

import (
	"eskept/internal/app/context"
	"eskept/internal/models"
	"eskept/internal/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProviderRepository struct {
	db *gorm.DB
}

func NewProviderRepository(appCtx *context.AppContext) *ProviderRepository {
	return &ProviderRepository{
		db: appCtx.DB,
	}
}

func (r *ProviderRepository) Create(provider *models.Provider) error {
	return r.db.Create(&provider).Error
}

func (r *ProviderRepository) Update(provider *models.Provider) error {
	return r.db.Save(&provider).Error
}

func (r *ProviderRepository) Delete(provider *models.Provider) error {
	return r.db.Delete(&provider).Error
}

func (r *ProviderRepository) FindByCodeName(codeName string) (*models.Provider, error) {
	var provider models.Provider
	err := r.db.First(&provider, "code_name = ?", codeName).Error
	if err != nil {
		return nil, err
	}
	return &provider, nil
}

func (r *ProviderRepository) FindByProviderId(providerId uuid.UUID) (*models.Provider, error) {
	var provider models.Provider
	err := r.db.First(&provider, "id = ?", providerId).Error
	if err != nil {
		return nil, err
	}
	return &provider, nil
}

func (r *ProviderRepository) List(pagination *types.Pagination) ([]models.Provider, error) {
	providers := []models.Provider{}
	
	// Count total records
	var total int64
	if err := r.db.Model(&models.Provider{}).Count(&total).Error; err != nil {
		return nil, err
	}
	pagination.Total = total

	// Apply pagination
	offset := (pagination.Page - 1) * pagination.PageSize
	err := r.db.Offset(offset).Limit(pagination.PageSize).Find(&providers).Error
	if err != nil {
		return nil, err
	}
	return providers, nil
}

func (r *ProviderRepository) Search(keyword string, pagination *types.Pagination) ([]models.Provider, error) {
	searchValue := GenerateSearchValue(keyword)
	providers := []models.Provider{}

	// Count total records
	var total int64
	if err := r.db.Model(&models.Provider{}).Where("name LIKE ?", searchValue).Count(&total).Error; err != nil {
		return nil, err
	}
	pagination.Total = total

	// Apply pagination
	offset := (pagination.Page - 1) * pagination.PageSize
	err := r.db.Where("name LIKE ?", searchValue).Offset(offset).Limit(pagination.PageSize).Find(&providers).Error
	if err != nil {
		return nil, err
	}
	return providers, nil
}
