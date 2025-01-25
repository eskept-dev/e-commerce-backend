package v1

import (
	"eskept/internal/app/context"
	"eskept/internal/handlers"
	"eskept/internal/middleware"
	"eskept/internal/repositories"
	"eskept/internal/services"

	"github.com/gin-gonic/gin"
)

func setupProductGroup(group *gin.RouterGroup, ctx *context.AppContext) {
	userRepository := repositories.NewUserRepository(ctx)
	productRepository := repositories.NewProductRepository(ctx)
	productService := services.NewProductService(productRepository, ctx)
	productHandler := handlers.NewProductHandler(productService, ctx)

	productGroupApi := group.Group("/products")
	{
		productGroupApi.Use(middleware.AuthMiddleware(userRepository, ctx))
		productGroupApi.GET("", productHandler.ListProducts)
		productGroupApi.GET("/:code_name", productHandler.GetProduct)
		productGroupApi.POST("", productHandler.CreateProduct)
	}
}
