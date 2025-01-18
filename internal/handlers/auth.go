package handlers

import (
	"eskept/internal/constants/errors"
	"eskept/internal/schemas"
	"eskept/internal/services"
	"log"
	"net/http"

	"eskept/internal/app/context"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
	appCtx  *context.AppContext
}

func NewAuthHandler(service *services.AuthService, appCtx *context.AppContext) *AuthHandler {
	return &AuthHandler{service: service, appCtx: appCtx}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req schemas.AuthRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(req.Email, req.Password)
	if err != nil {
		log.Fatalln(err)
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

func (h *AuthHandler) Login(c *gin.Context) {
	var req schemas.AuthLoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenPair, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		log.Fatalln(err)
		if err == errors.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrInvalidCredentials.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": errors.ErrInternalServerError.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, schemas.AuthLoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	})
}
