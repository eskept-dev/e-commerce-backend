package services

import (
	"eskept/internal/app/context"
	"eskept/internal/models"
	"eskept/internal/repositories"
	"eskept/internal/types"
)

type ProductService struct {
	ProductRepository *repositories.ProductRepository
	appCtx            *context.AppContext
}

func NewProductService(productRepository *repositories.ProductRepository, appCtx *context.AppContext) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
		appCtx:            appCtx,
	}
}

func (s *ProductService) Create(product *models.Product) error {
	return s.ProductRepository.Create(product)
}

func (s *ProductService) Update(product *models.Product) error {
	return s.ProductRepository.Update(product)
}

func (s *ProductService) Delete(product *models.Product) error {
	return s.ProductRepository.Delete(product)
}

func (s *ProductService) FindByCodeName(codeName string) (*models.Product, error) {
	return s.ProductRepository.FindByCodeName(codeName)
}

func (s *ProductService) List(
	serviceCode string,
	keyword string,
	pagination *types.Pagination,
) ([]models.Product, error) {
	return s.ProductRepository.FindProducts(serviceCode, keyword, pagination)
}
