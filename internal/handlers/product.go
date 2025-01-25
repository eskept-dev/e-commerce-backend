package handlers

import (
	"eskept/internal/app/context"
	"eskept/internal/models"
	"eskept/internal/schemas"
	"eskept/internal/services"
	"eskept/internal/types"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService *services.ProductService
	appCtx         *context.AppContext
}

func NewProductHandler(productService *services.ProductService, appCtx *context.AppContext) *ProductHandler {
	return &ProductHandler{
		ProductService: productService,
		appCtx:         appCtx,
	}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req schemas.ProductCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product := &models.Product{
		Name:               req.Name,
		ThumbnailURL:       req.ThumbnailURL,
		ImageURL:           req.ImageURL,
		Description:        req.Description,
		ShortDescription:   req.ShortDescription,
		UnitType:           req.UnitType,
		Details:            req.Details,
		Highlights:         strings.Join(req.Highlights, ";"),
		CancellationPolicy: req.CancellationPolicy,
		ServiceID:          req.ServiceID,
		ProviderID:         req.ProviderID,
	}
	err := h.ProductService.Create(product)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := schemas.ProductResponse{
		ID:                 product.ID,
		CreatedAt:          product.CreatedAt,
		UpdatedAt:          product.UpdatedAt,
		IsEnabled:          product.IsEnabled,
		Name:               product.Name,
		CodeName:           product.CodeName,
		ThumbnailURL:       product.ThumbnailURL,
		ImageURL:           product.ImageURL,
		Description:        product.Description,
		ShortDescription:   product.ShortDescription,
		UnitType:           product.UnitType,
		Details:            product.Details,
		Highlights:         strings.Split(product.Highlights, ";"),
		CancellationPolicy: product.CancellationPolicy,
		AnalysisMetrics: schemas.AnalysisMetrics{
			TotalBooking: product.AnalysisMetrics.TotalBooking,
			AvgRating:    product.AnalysisMetrics.AvgRating,
		},
		Service: schemas.ServiceResponse{
			ID:        product.Service.ID,
			Name:      product.Service.Name,
			Code:      product.Service.Code,
			IsEnabled: product.Service.IsEnabled,
		},
		Provider: schemas.ProviderResponse{
			ID:        product.Provider.ID,
			CreatedAt: product.Provider.CreatedAt,
			UpdatedAt: product.Provider.UpdatedAt,
			Name:      product.Provider.Name,
			CodeName:  product.Provider.CodeName,
			IsEnabled: product.Provider.IsEnabled,
		},
	}

	ctx.JSON(200, response)
}

func (h *ProductHandler) GetProduct(ctx *gin.Context) {
	codeName := ctx.Param("codeName")
	product, err := h.ProductService.FindByCodeName(codeName)
	if err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	response := schemas.ProductResponse{
		ID:                 product.ID,
		CreatedAt:          product.CreatedAt,
		UpdatedAt:          product.UpdatedAt,
		IsEnabled:          product.IsEnabled,
		Name:               product.Name,
		CodeName:           product.CodeName,
		ThumbnailURL:       product.ThumbnailURL,
		ImageURL:           product.ImageURL,
		Description:        product.Description,
		ShortDescription:   product.ShortDescription,
		UnitType:           product.UnitType,
		Details:            product.Details,
		Highlights:         strings.Split(product.Highlights, ";"),
		CancellationPolicy: product.CancellationPolicy,
		AnalysisMetrics: schemas.AnalysisMetrics{
			TotalBooking: product.AnalysisMetrics.TotalBooking,
			AvgRating:    product.AnalysisMetrics.AvgRating,
		},
		Service: schemas.ServiceResponse{
			ID:        product.Service.ID,
			Name:      product.Service.Name,
			Code:      product.Service.Code,
			IsEnabled: product.Service.IsEnabled,
		},
		Provider: schemas.ProviderResponse{
			ID:        product.Provider.ID,
			CreatedAt: product.Provider.CreatedAt,
			UpdatedAt: product.Provider.UpdatedAt,
			Name:      product.Provider.Name,
			CodeName:  product.Provider.CodeName,
			IsEnabled: product.Provider.IsEnabled,
		},
	}

	ctx.JSON(200, response)
}

func (h *ProductHandler) ListProducts(ctx *gin.Context) {
	// Get pagination parameters
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	keyword := ctx.DefaultQuery("keyword", "")
	serviceCode := ctx.DefaultQuery("serviceCode", "")

	log.Println("------------------- List products -------------------")
	log.Println("Page:", page)
	log.Println("PageSize:", pageSize)
	log.Println("Keyword:", keyword)
	log.Println("ServiceCode:", serviceCode)
	log.Println("------------------------------------------------------------")

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	pagination := &types.Pagination{
		Page:     page,
		PageSize: pageSize,
	}

	products, err := h.ProductService.List(
		serviceCode,
		keyword,
		pagination,
	)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	productResponse := []schemas.ProductResponse{}
	for _, product := range products {
		productResponse = append(productResponse, schemas.ProductResponse{
			ID:                 product.ID,
			CreatedAt:          product.CreatedAt,
			UpdatedAt:          product.UpdatedAt,
			IsEnabled:          product.IsEnabled,
			Name:               product.Name,
			CodeName:           product.CodeName,
			ThumbnailURL:       product.ThumbnailURL,
			ImageURL:           product.ImageURL,
			Description:        product.Description,
			ShortDescription:   product.ShortDescription,
			UnitType:           product.UnitType,
			Details:            product.Details,
			Highlights:         strings.Split(product.Highlights, ";"),
			CancellationPolicy: product.CancellationPolicy,
			AnalysisMetrics: schemas.AnalysisMetrics{
				TotalBooking: product.AnalysisMetrics.TotalBooking,
				AvgRating:    product.AnalysisMetrics.AvgRating,
			},
			Service: schemas.ServiceResponse{
				ID:        product.Service.ID,
				Name:      product.Service.Name,
				Code:      product.Service.Code,
				IsEnabled: product.Service.IsEnabled,
			},
			Provider: schemas.ProviderResponse{
				ID:        product.Provider.ID,
				CreatedAt: product.Provider.CreatedAt,
				UpdatedAt: product.Provider.UpdatedAt,
				Name:      product.Provider.Name,
				CodeName:  product.Provider.CodeName,
				IsEnabled: product.Provider.IsEnabled,
			},
		})
	}

	response := types.PaginatedResponse{
		Pagination: *pagination,
		Data:       productResponse,
	}

	ctx.JSON(200, response)
}
