package v1

import (
	"eskept/internal/app/context"
	"eskept/internal/handlers"
	"eskept/internal/repositories"
	"eskept/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupV1Routes(group *gin.RouterGroup, ctx *context.AppContext) {
	setupAuthGroup(group, ctx)
}

func setupAuthGroup(group *gin.RouterGroup, ctx *context.AppContext) {
	userRepository := repositories.NewUserRepository(ctx)
	authService := services.NewAuthService(userRepository, ctx)
	authHandler := handlers.NewAuthHandler(userRepository, authService, ctx)

	userGroupApi := group.Group("/auth")
	{
		userGroupApi.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "OK"})
		})
		userGroupApi.POST("/register", authHandler.Register)
		userGroupApi.POST("/login", authHandler.Login)
		userGroupApi.POST("/send-activation-link", authHandler.SendActivationLink)
		userGroupApi.POST("/activate", authHandler.Activate)
	}
}
