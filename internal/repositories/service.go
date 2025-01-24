package repositories

import (
	"eskept/internal/app/context"
	"eskept/internal/models"

	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(ctx *context.AppContext) *ServiceRepository {
	return &ServiceRepository{
		db: ctx.DB,
	}
}

func (r *ServiceRepository) Create(service *models.Service) error {
	return r.db.Create(&service).Error
}

func (r *ServiceRepository) Update(service *models.Service) error {
	return r.db.Save(&service).Error
}

func (r *ServiceRepository) Delete(service *models.Service) error {
	return r.db.Delete(&service).Error
}

func (r *ServiceRepository) FindByCode(code string) (*models.Service, error) {
	var service models.Service
	err := r.db.First(&service, "code = ?", code).Error
	if err != nil {
		return nil, err
	}
	return &service, nil
}
