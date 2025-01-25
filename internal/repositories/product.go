package repositories

import (
	"eskept/internal/app/context"
	"eskept/internal/models"
	"eskept/internal/types"
	"log"

	"gorm.io/gorm"
)

type ProductRepository struct {
	appCtx *context.AppContext
	db     *gorm.DB
}

func NewProductRepository(appCtx *context.AppContext) *ProductRepository {
	return &ProductRepository{
		appCtx: appCtx,
		db:     appCtx.DB,
	}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(product *models.Product) error {
	return r.db.Delete(product).Error
}

func (r *ProductRepository) FindByCodeName(codeName string) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, "code_name = ?", codeName).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) FindProducts(
	serviceCode string,
	keyword string,
	pagination *types.Pagination,
) ([]models.Product, error) {
	query := r.db.Model(&models.Product{})

	if serviceCode != "" {
		var service models.Service
		err := r.appCtx.DB.Where("code = ?", serviceCode).First(&service).Error
		if err != nil {
			return nil, err
		}
		query = query.Joins("Service", "Service.code = ?", serviceCode)
	}

	if keyword != "" {
		searchValue := GenerateSearchValue(keyword)
		query = query.Where("code_name LIKE ?", searchValue)
		log.Println("SearchValue:", searchValue)
	}

	// Count total records
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	pagination.Total = total

	// Apply pagination
	offset := (pagination.Page - 1) * pagination.PageSize
	var products []models.Product
	err := query.Offset(offset).Limit(pagination.PageSize).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
