package v1

import (
	"eskept/internal/app/context"
	"eskept/internal/handlers"
	"eskept/internal/middleware"
	"eskept/internal/repositories"
	"eskept/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupV1Routes(group *gin.RouterGroup, ctx *context.AppContext) {
	setupAuthGroup(group, ctx)
	setupUserGroup(group, ctx)
}

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
		authGroupApi.POST("/activate", authHandler.Activate)
	}

	protectedGroupApi := group.Group("/auth")
	{
		// Apply auth middleware to protected routes
		protectedGroupApi.Use(middleware.AuthMiddleware(ctx))
		{
			protectedGroupApi.GET("/verify-token", authHandler.VerifyToken)
		}
	}
}

func setupUserGroup(group *gin.RouterGroup, ctx *context.AppContext) {
	userRepository := repositories.NewUserRepository(ctx)
	userHandler := handlers.NewUserHandler(userRepository, ctx)

	// Apply auth middleware to user routes
	userGroupApi := group.Group("/users")
	userGroupApi.Use(middleware.AuthMiddleware(ctx))
	{
		userGroupApi.GET("/me", userHandler.GetMe)
	}
}
