package handlers

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/errors"
	"eskept/internal/models"
	"eskept/internal/repositories"
	"eskept/internal/schemas"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userProfileRepo *repositories.UserProfileRepository
	userRepo        *repositories.UserRepository
	appCtx          *context.AppContext
}

func NewUserHandler(
	userRepo *repositories.UserRepository,
	userProfileRepo *repositories.UserProfileRepository,
	appCtx *context.AppContext,
) *UserHandler {
	return &UserHandler{
		userProfileRepo: userProfileRepo,
		userRepo:        userRepo,
		appCtx:          appCtx,
	}
}

func (h *UserHandler) GetMe(c *gin.Context) {
	email := c.MustGet("email").(string)

	user, err := h.userRepo.FindByEmail(email)
	if err != nil {
		log.Println(err)
		if err == errors.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, schemas.UserGetMeResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
		Status:    string(user.Status),
		Email:     user.Email,
		Role:      string(user.Role),
	})
}

func (h *UserHandler) CreateUserProfile(c *gin.Context) {
	var req schemas.UserProfileCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(*models.User)

	isExists, _ := h.userProfileRepo.FindByUserId(user.ID)
	if isExists != nil {
		c.JSON(http.StatusConflict, gin.H{"error": errors.ErrAlreadyExists.Error()})
		return
	}

	profile := &models.UserProfile{
		UserId:      user.ID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateOfBirth: req.DateOfBirth,
		Sex:         req.Sex,
		Nationality: req.Nationality,
		DialCode:    req.DialCode,
		PhoneNumber: req.PhoneNumber,
		Email:       user.Email,
	}

	err := h.userProfileRepo.Create(profile)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		return
	}

	c.JSON(http.StatusCreated, schemas.UserProfileCreateResponse{
		ID:          profile.ID,
		CreatedAt:   profile.CreatedAt,
		UpdatedAt:   profile.UpdatedAt,
		UserId:      profile.UserId,
		FirstName:   profile.FirstName,
		LastName:    profile.LastName,
		DateOfBirth: profile.DateOfBirth,
		Sex:         profile.Sex,
		Nationality: profile.Nationality,
		DialCode:    profile.DialCode,
		PhoneNumber: profile.PhoneNumber,
		Email:       profile.Email,
	})
}
