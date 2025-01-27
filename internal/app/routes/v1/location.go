package v1

import (
	"eskept/internal/app/context"
	"eskept/internal/handlers"
	"eskept/internal/repositories"
	"eskept/internal/services"

	"github.com/gin-gonic/gin"
)

func setUpLocationGroup(group *gin.RouterGroup, ctx *context.AppContext) {
	locationRepository := repositories.NewLocationRepository(ctx)
	locationService := services.NewLocationService(locationRepository, ctx)
	locationHandler := handlers.NewLocationHandler(locationRepository, locationService, ctx)

	locationGroupApi := group.Group("/locations")
	{
		locationGroupApi.GET("", locationHandler.List)
	}
}
