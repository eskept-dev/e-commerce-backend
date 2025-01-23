package v1

import (
	"eskept/internal/app/context"
	"eskept/internal/handlers"
	"eskept/internal/middleware"
	"eskept/internal/repositories"

	"github.com/gin-gonic/gin"
)

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
