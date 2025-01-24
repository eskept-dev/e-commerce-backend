package handlers

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/errors"
	"eskept/internal/models"
	"eskept/internal/schemas"
	"eskept/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProviderHandler struct {
	appCtx          *context.AppContext
	ProviderService *services.ProviderService
}

func NewProviderHandler(
	providerService *services.ProviderService,
	appCtx *context.AppContext,
) *ProviderHandler {
	return &ProviderHandler{
		appCtx:          appCtx,
		ProviderService: providerService,
	}
}

func (h *ProviderHandler) CreateProvider(c *gin.Context) {
	var req schemas.ProviderCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	businessInfo := models.BusinessInformation{
		PhoneNumber: req.BusinessInformation.PhoneNumber,
		Email:       req.BusinessInformation.Email,
		Address:     req.BusinessInformation.Address,
		Website:     req.BusinessInformation.Website,
	}
	contactInfo := models.ContactInformation{
		FirstName:   req.ContactInformation.FirstName,
		LastName:    req.ContactInformation.LastName,
		Role:        req.ContactInformation.Role,
		Gender:      req.ContactInformation.Gender,
		PhoneNumber: req.ContactInformation.PhoneNumber,
		Email:       req.ContactInformation.Email,
	}

	provider := &models.Provider{
		Name:                req.Name,
		BusinessInformation: businessInfo,
		ContactInformation:  contactInfo,
	}

	err := h.ProviderService.CreateProvider(provider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		return
	}

	businessInformationResponse := schemas.BusinessInformation{
		PhoneNumber: provider.BusinessInformation.PhoneNumber,
		Email:       provider.BusinessInformation.Email,
		Address:     provider.BusinessInformation.Address,
		Website:     provider.BusinessInformation.Website,
	}
	contactInformationResponse := schemas.ContactInformation{
		FirstName:   provider.ContactInformation.FirstName,
		LastName:    provider.ContactInformation.LastName,
		Role:        provider.ContactInformation.Role,
		Gender:      provider.ContactInformation.Gender,
		PhoneNumber: provider.ContactInformation.PhoneNumber,
		Email:       provider.ContactInformation.Email,
	}

	response := schemas.ProviderResponse{
		ID:                  provider.ID,
		CreatedAt:           provider.CreatedAt,
		UpdatedAt:           provider.UpdatedAt,
		Name:                provider.Name,
		CodeName:            provider.CodeName,
		BusinessInformation: businessInformationResponse,
		ContactInformation:  contactInformationResponse,
		IsEnabled:           provider.IsEnabled,
	}

	c.JSON(http.StatusOK, response)
}

func (h *ProviderHandler) GetProvider(c *gin.Context) {
	providerCodeName := c.Param("code_name")

	provider, err := h.ProviderService.FindByCodeName(providerCodeName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		return
	}

	businessInformationResponse := schemas.BusinessInformation{
		PhoneNumber: provider.BusinessInformation.PhoneNumber,
		Email:       provider.BusinessInformation.Email,
		Address:     provider.BusinessInformation.Address,
		Website:     provider.BusinessInformation.Website,
	}
	contactInformationResponse := schemas.ContactInformation{
		FirstName:   provider.ContactInformation.FirstName,
		LastName:    provider.ContactInformation.LastName,
		Role:        provider.ContactInformation.Role,
		Gender:      provider.ContactInformation.Gender,
		PhoneNumber: provider.ContactInformation.PhoneNumber,
		Email:       provider.ContactInformation.Email,
	}

	response := schemas.ProviderResponse{
		ID:                  provider.ID,
		CreatedAt:           provider.CreatedAt,
		UpdatedAt:           provider.UpdatedAt,
		Name:                provider.Name,
		CodeName:            provider.CodeName,
		BusinessInformation: businessInformationResponse,
		ContactInformation:  contactInformationResponse,
		IsEnabled:           provider.IsEnabled,
	}

	c.JSON(http.StatusOK, response)
}
