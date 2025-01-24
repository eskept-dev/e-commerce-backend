package v1

import (
	"eskept/internal/app/context"
	"eskept/internal/handlers"
	"eskept/internal/middleware"
	"eskept/internal/repositories"
	"eskept/internal/services"

	"github.com/gin-gonic/gin"
)

func setupProviderGroup(group *gin.RouterGroup, ctx *context.AppContext) {
	userRepository := repositories.NewUserRepository(ctx)
	providerRepository := repositories.NewProviderRepository(ctx)
	providerService := services.NewProviderService(providerRepository, ctx)
	providerHandler := handlers.NewProviderHandler(providerService, ctx)

	providerGroupApi := group.Group("/providers")
	{
		providerGroupApi.Use(middleware.AuthMiddleware(userRepository, ctx))
		providerGroupApi.POST("", providerHandler.CreateProvider)
		providerGroupApi.GET("/:code_name", providerHandler.GetProvider)
	}
}
