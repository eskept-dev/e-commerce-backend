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
	gin.SetMode(gin.ReleaseMode)
	r.routerEngine = gin.New()
	
	// Disable automatic trailing slash redirects
	r.routerEngine.RedirectTrailingSlash = false
	r.routerEngine.RedirectFixedPath = false

	// Add recovery middleware
	r.routerEngine.Use(gin.Recovery())

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:5173",
		"http://localhost:3000",
		"http://127.0.0.1:5173",
		"http://127.0.0.1:3000",
	}
	config.AllowMethods = []string{
		"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS",
	}
	config.AllowHeaders = []string{
		"Origin",
		"Content-Length",
		"Content-Type",
		"Accept",
		"Authorization",
		"X-Requested-With",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Methods",
	}
	config.ExposeHeaders = []string{
		"Content-Length",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Methods",
	}
	config.AllowCredentials = true
	config.MaxAge = 12 * 60 * 60 // 12 hours

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
