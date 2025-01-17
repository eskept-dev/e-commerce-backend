package handlers

import (
	"eskept/internal/constants/errors"
	"eskept/internal/schemas"
	"eskept/internal/services"
	"net/http"

	"eskept/internal/app/context"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
	appCtx  *context.AppContext
}

func NewUserHandler(service *services.UserService, appCtx *context.AppContext) *UserHandler {
	return &UserHandler{service: service, appCtx: appCtx}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req schemas.AuthRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(req.Email, req.Password)
	if err != nil {
		if err == errors.ErrEmailExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":  "could not register user",
				"detail": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req schemas.AuthLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		if err == errors.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":  "could not login user",
				"detail": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
