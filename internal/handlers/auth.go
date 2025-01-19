package handlers

import (
	"eskept/internal/constants/errors"
	"eskept/internal/repositories"
	"eskept/internal/schemas"
	"eskept/internal/services"
	"log"
	"net/http"

	"eskept/internal/app/context"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	repo         *repositories.UserRepository
	service      *services.AuthService
	emailService *services.EmailService
	appCtx       *context.AppContext
}

func NewAuthHandler(
	repo *repositories.UserRepository,
	service *services.AuthService,
	emailService *services.EmailService,
	appCtx *context.AppContext,
) *AuthHandler {
	return &AuthHandler{
		repo:         repo,
		service:      service,
		emailService: emailService,
		appCtx:       appCtx,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req schemas.AuthRegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(req.Email, req.Password)
	if err != nil {
		log.Println(err.Error())
		if err == errors.ErrEmailExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
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

	_, err := h.service.IsAuthenticated(req.Email, req.Password)
	if err != nil {
		log.Println(err.Error())
		if err == errors.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	tokenPair, err := h.service.GenerateTokens(req.Email, req.Password)
	if err != nil {
		log.Println(err)
		if err == errors.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, schemas.AuthLoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	})
}

func (h *AuthHandler) LoginByAuthenticationToken(c *gin.Context) {
	var req schemas.AuthLoginByAuthenticationTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenPair, err := h.service.LoginByAuthenticationToken(req.AuthenticationToken)
	if err != nil {
		log.Println(err)
		if err == errors.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else if err == errors.ErrTokenExpired {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, schemas.AuthLoginByAuthenticationTokenResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	})
}

func (h *AuthHandler) SendAuthenticationEmail(c *gin.Context) {
	var req schemas.AuthSendAuthenticationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.repo.FindByEmail(req.Email)
	if err != nil {
		log.Println(err)
		if err == errors.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		}
		return
	}

	err = h.emailService.SendAuthenticationEmail(req.Email, string(user.Role))
	if err != nil {
		log.Println(err)
		if err == errors.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, schemas.AuthSendAuthenticationResponse{
		IsSuccess: true,
	})
}

func (h *AuthHandler) SendActivationEmail(c *gin.Context) {
	var req schemas.AuthSendActivationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.repo.FindByEmail(req.Email)
	if err != nil {
		log.Println(err)
		if err == errors.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		}
		return
	}

	err = h.emailService.SendActivationEmail(req.Email, string(user.Role))
	if err != nil {
		log.Println(err)
		if err == errors.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, schemas.AuthSendActivationResponse{
		IsSuccess: true,
	})
}

func (h *AuthHandler) Activate(c *gin.Context) {
	var req schemas.AuthActivateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.ActivateUserByActivationToken(req.ActivationToken)
	if err != nil {
		log.Println(err)
		switch err {
		case errors.ErrTokenExpired, errors.ErrInvalidToken:
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, schemas.AuthActivateResponse{
		IsActivated: true,
	})
}

func (h *AuthHandler) VerifyToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
