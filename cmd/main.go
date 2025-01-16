package main

import (
	"eskept/pkg/config"
	"eskept/pkg/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set gin mode based on environment
	env := os.Getenv("APP_ENV")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Load configuration
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db := database.InitPostgres(cfg.Database.URL)
	_ = db // Temporary fix for unused variable

	// Initialize Gin router
	r := gin.Default()

	// Basic health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"env":    env,
		})
	})

	// Start server
	log.Printf("Server starting on port %s in %s mode", cfg.Server.Port, env)
	if err := r.Run(cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
