package v1

import (
	"eskept/internal/app/context"
	"eskept/internal/handlers"
	"eskept/internal/middleware"
	"eskept/internal/repositories"
	"eskept/internal/services"

	"github.com/gin-gonic/gin"
)

func setupBusinessGroup(group *gin.RouterGroup, ctx *context.AppContext) {
	userRepository := repositories.NewUserRepository(ctx)
	userProfileRepository := repositories.NewUserProfileRepository(ctx)
	userService := services.NewUserService(userRepository, userProfileRepository, ctx)
	businessProfileRepository := repositories.NewBusinessProfileRepository(ctx)
	businessService := services.NewBusinessService(businessProfileRepository, userRepository, ctx)
	businessHandler := handlers.NewBusinessHandler(businessService, userService, ctx)

	businessGroupApi := group.Group("/business-profiles")
	{
		businessGroupApi.Use(middleware.AuthMiddleware(userRepository, ctx))
		businessGroupApi.GET("/:id", businessHandler.GetProfile)
		businessGroupApi.POST("", businessHandler.CreateProfile)
	}
}
