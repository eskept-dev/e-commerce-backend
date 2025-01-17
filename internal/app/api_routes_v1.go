package app

import (
	"eskept/internal/context"
	"eskept/internal/handlers"
	"eskept/internal/repositories"
	"eskept/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) setupV1Routes(apiV1 *gin.RouterGroup, ctx *context.AppContext) {
	setupUserGroup(apiV1, ctx)
}

func setupUserGroup(apiV1 *gin.RouterGroup, ctx *context.AppContext) {
	userRepository := repositories.NewUserRepository(ctx)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	userGroupApi := apiV1.Group("/users")
	{
		userGroupApi.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "OK"})
		})
		userGroupApi.POST("/register", userHandler.Register)
	}
}
