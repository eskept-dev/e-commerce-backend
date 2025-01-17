package app

import (
	"eskept/pkg/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(ctx *context.AppContext) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", ctx.DB)
		c.Next()
	})

	// ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return router
}
