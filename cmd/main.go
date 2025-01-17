package main

import (
	"eskept/internal/app"
	"eskept/pkg/config"
	"eskept/pkg/context"
	"eskept/pkg/database"
	"fmt"

	"gorm.io/gorm"
)

type AppContext struct {
	DB *gorm.DB
}

func main() {
	// Load configuration
	config, err := config.LoadConfig("./config")
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	// Initialize database connection
	db, err := database.InitPostgres(config.Database)
	if err != nil {
		panic(fmt.Errorf("failed to initialize database: %w", err))
	}

	// App context
	AppContext := &context.AppContext{
		DB: db,
	}

	// Initialize router
	router := app.NewRouter(AppContext)

	// Start the server
	router.Run(fmt.Sprintf(":%s", config.Server.Port))
}
