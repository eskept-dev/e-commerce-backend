package app

import (
	"eskept/pkg/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	routerEngine *gin.Engine
}

func (r *Router) NewRouter(ctx *context.AppContext) error {
	r.routerEngine = gin.Default()
	r.routerEngine.Use(func(c *gin.Context) {
		c.Set("db", ctx.DB)
		c.Next()
	})

	return nil
}

func (r *Router) Run(port string) {
	r.routerEngine.Run(port)
}

func (r *Router) SetupRouter(ctx *context.AppContext) {
	// Health Check Route
	r.routerEngine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// Setup API versions
	apiV1 := r.routerEngine.Group("/api/v1")
	{
		r.setupV1Routes(apiV1, ctx)
	}
}
