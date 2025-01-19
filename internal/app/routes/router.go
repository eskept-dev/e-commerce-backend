package routes

import (
	"eskept/internal/app/context"
	v1 "eskept/internal/app/routes/v1"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	routerEngine *gin.Engine
}

func (r *Router) NewRouter(ctx *context.AppContext) error {
	r.routerEngine = gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Add your frontend URL here
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true

	r.routerEngine.Use(cors.New(config))

	r.routerEngine.Use(func(c *gin.Context) {
		c.Set("db", ctx.DB)
		c.Set("cache", ctx.Cache)
		c.Set("cfg", ctx.Cfg)
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

	v1Group := r.routerEngine.Group("/api/v1")
	{
		v1.SetupV1Routes(v1Group, ctx)
	}
}
