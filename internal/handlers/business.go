package handlers

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/enums"
	"eskept/internal/constants/errors"
	"eskept/internal/models"
	"eskept/internal/schemas"
	"eskept/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BusinessHandler struct {
	userService     *services.UserService
	businessService *services.BusinessService
	appCtx          *context.AppContext
}

func NewBusinessHandler(
	businessService *services.BusinessService,
	userService *services.UserService,
	ctx *context.AppContext,
) *BusinessHandler {
	return &BusinessHandler{
		businessService: businessService,
		userService:     userService,
		appCtx:          ctx,
	}
}

func (h *BusinessHandler) GetProfile(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	if user.Role != enums.UserRoleBusiness {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUnauthorized.Error()})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	businessProfile, err := h.businessService.GetProfile(id)
	if err != nil {
		log.Println(err)
		if err == errors.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		}
		return
	}

	if businessProfile.RepresentativeUserId != user.ID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUnauthorized.Error()})
		return
	}

	c.JSON(http.StatusOK, schemas.BusinessProfileGetResponse{
		ID:                      businessProfile.ID,
		CreatedAt:               businessProfile.CreatedAt,
		UpdatedAt:               businessProfile.UpdatedAt,
		BusinessName:            businessProfile.BusinessName,
		BusinessTaxId:           businessProfile.BusinessTaxId,
		BusinessAddress:         businessProfile.BusinessAddress,
		BusinessDialCode:        businessProfile.BusinessDialCode,
		BusinessPhoneNumber:     businessProfile.BusinessPhoneNumber,
		BusinessEmail:           businessProfile.BusinessEmail,
		BusinessWebsite:         businessProfile.BusinessWebsite,
		BusinessNationality:     businessProfile.BusinessNationality,
		RepresentativeUserId:    businessProfile.RepresentativeUserId,
		RepresentativeProfileId: businessProfile.RepresentativeProfileId,
	})
}

func (h *BusinessHandler) CreateProfile(c *gin.Context) {
	var req schemas.BusinessProfileCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(*models.User)
	if user.Role != enums.UserRoleBusiness {
		c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUnauthorized.Error()})
		return
	}

	representativeProfile, err := h.userService.GetProfileByUserId(user.ID)
	if err != nil {
		log.Println(err)
		if err == errors.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		}
		return
	}

	businessProfile := &models.BusinessProfile{
		BusinessName:            req.BusinessName,
		BusinessTaxId:           req.BusinessTaxId,
		BusinessAddress:         req.BusinessAddress,
		BusinessDialCode:        req.BusinessDialCode,
		BusinessPhoneNumber:     req.BusinessPhoneNumber,
		BusinessEmail:           req.BusinessEmail,
		BusinessWebsite:         req.BusinessWebsite,
		BusinessNationality:     req.BusinessNationality,
		RepresentativeUserId:    user.ID,
		RepresentativeProfileId: representativeProfile.ID,
	}

	err = h.businessService.CreateProfile(businessProfile)
	if err != nil {
		log.Println(err)
		if err == errors.ErrAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, schemas.BusinessProfileCreateResponse{
		ID:                      businessProfile.ID,
		CreatedAt:               businessProfile.CreatedAt,
		UpdatedAt:               businessProfile.UpdatedAt,
		BusinessName:            businessProfile.BusinessName,
		BusinessTaxId:           businessProfile.BusinessTaxId,
		BusinessAddress:         businessProfile.BusinessAddress,
		BusinessDialCode:        businessProfile.BusinessDialCode,
		BusinessPhoneNumber:     businessProfile.BusinessPhoneNumber,
		BusinessEmail:           businessProfile.BusinessEmail,
		BusinessWebsite:         businessProfile.BusinessWebsite,
		BusinessNationality:     businessProfile.BusinessNationality,
		RepresentativeUserId:    user.ID,
		RepresentativeProfileId: representativeProfile.ID,
	})
}
