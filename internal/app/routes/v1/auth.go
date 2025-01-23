package v1

import (
	"eskept/internal/app/context"
	"eskept/internal/handlers"
	"eskept/internal/middleware"
	"eskept/internal/repositories"
	"eskept/internal/services"

	"github.com/gin-gonic/gin"
)

func setupAuthGroup(group *gin.RouterGroup, ctx *context.AppContext) {
	userRepository := repositories.NewUserRepository(ctx)
	authService := services.NewAuthService(userRepository, ctx)
	emailService := services.NewEmailService(userRepository, ctx)
	authHandler := handlers.NewAuthHandler(userRepository, authService, emailService, ctx)

	authGroupApi := group.Group("/auth")
	{

		authGroupApi.POST("/register", authHandler.Register)
		authGroupApi.POST("/login-by-authentication-token", authHandler.LoginByAuthenticationToken)
		authGroupApi.POST("/login", authHandler.Login)
		authGroupApi.POST("/send-authentication-email", authHandler.SendAuthenticationEmail)
		authGroupApi.POST("/send-activation-email", authHandler.SendActivationEmail)
		authGroupApi.POST("/send-verification-email", authHandler.SendVerificationEmail)
		authGroupApi.POST("/verify-email-token", authHandler.VerifyEmailToken)
		authGroupApi.POST("/activate", authHandler.Activate)
	}

	protectedGroupApi := group.Group("/auth")
	{
		// Apply auth middleware to protected routes
		protectedGroupApi.Use(middleware.AuthMiddleware(userRepository, ctx))
		{
			protectedGroupApi.GET("/verify-token", authHandler.VerifyToken)
		}
	}
}
