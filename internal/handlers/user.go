package handlers

import (
	"eskept/internal/constants/errors"
	"eskept/internal/schemas"
	"eskept/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req schemas.UserRegisterRequest

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
	var req schemas.UserLoginRequest

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
