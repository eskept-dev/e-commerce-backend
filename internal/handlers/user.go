package handlers

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/errors"
	"eskept/internal/repositories"
	"eskept/internal/schemas"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo   *repositories.UserRepository
	appCtx *context.AppContext
}

func NewUserHandler(repo *repositories.UserRepository, appCtx *context.AppContext) *UserHandler {
	return &UserHandler{
		repo:   repo,
		appCtx: appCtx,
	}
}

func (h *UserHandler) GetMe(c *gin.Context) {
	email := c.MustGet("email").(string)

	user, err := h.repo.FindByEmail(email)
	if err != nil {
		log.Println(err)
		if err == errors.ErrUserNotFound {
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
