package main

import (
	"eskept/internal/app"
	"eskept/pkg/config"
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

	server := app.NewServer(config)
	server.Run()
}
