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
	setupProfileGroup(group, ctx)
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

func setupUserGroup(group *gin.RouterGroup, ctx *context.AppContext) {
	userRepository := repositories.NewUserRepository(ctx)
	userProfileRepository := repositories.NewUserProfileRepository(ctx)
	userHandler := handlers.NewUserHandler(userRepository, userProfileRepository, ctx)

	// Apply auth middleware to user routes
	userGroupApi := group.Group("/users")
	{
		userGroupApi.Use(middleware.AuthMiddleware(userRepository, ctx))
		userGroupApi.GET("/me", userHandler.GetMe)
	}
}

func setupProfileGroup(group *gin.RouterGroup, ctx *context.AppContext) {
	userRepository := repositories.NewUserRepository(ctx)
	userProfileRepository := repositories.NewUserProfileRepository(ctx)
	userHandler := handlers.NewUserHandler(userRepository, userProfileRepository, ctx)

	// Apply auth middleware to user routes
	userProfileGroupApi := group.Group("/profiles")
	{
		userProfileGroupApi.Use(middleware.AuthMiddleware(userRepository, ctx))
		userProfileGroupApi.POST("", userHandler.CreateUserProfile)
	}
}
