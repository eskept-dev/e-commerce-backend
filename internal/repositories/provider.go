package repositories

import (
	"eskept/internal/app/context"
	"eskept/internal/models"

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

func (r *ProviderRepository) List() ([]models.Provider, error) {
	providers := []models.Provider{}
	err := r.db.Find(&providers).Error
	if err != nil {
		return nil, err
	}
	return providers, nil
}

func (r *ProviderRepository) Search(keyword string) ([]models.Provider, error) {
	searchValue := GenerateSearchValue(keyword)
	providers := []models.Provider{}
	err := r.db.Where("name LIKE ?", searchValue).Find(&providers).Error
	if err != nil {
		return nil, err
	}
	return providers, nil
}
