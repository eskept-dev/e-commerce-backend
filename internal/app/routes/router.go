package routes

import (
	"eskept/internal/app/context"
	v1 "eskept/internal/app/routes/v1"
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
	v1Group := r.routerEngine.Group("/api/v1")
	{
		v1.SetupV1Routes(v1Group, ctx)
	}
}
