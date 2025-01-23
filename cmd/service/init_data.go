package main

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/enums"
	"eskept/internal/models"
	"eskept/pkg/cache"
	"eskept/pkg/config"
	"eskept/pkg/database"
	"fmt"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database connection
	db, err := database.InitPostgres(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize cache connection
	cacheRedis, err := cache.InitRedis(&cfg.Cache)
	if err != nil {
		log.Fatalf("Failed to initialize cache: %v", err)
	}

	// Initialize app context properly
	appCtx := &context.AppContext{
		DB:    db,
		Cache: cacheRedis,
		Cfg:   cfg,
	}

	// Default services to create
	services := []models.Service{
		{
			Name:      "Airport Transfer",
			Type:      enums.AirportTransfer,
			IsEnabled: true,
		},
		{
			Name:      "Fast Track",
			Type:      enums.FastTrack,
			IsEnabled: true,
		},
		{
			Name:      "E-Visa",
			Type:      enums.EVisa,
			IsEnabled: true,
		},
	}

	// Create each service
	for _, service := range services {
		// Create new service
		if err := appCtx.DB.Create(&service).Error; err != nil {
			log.Printf("Failed to create service %s: %v\n", service.Name, err)
			continue
		}

		fmt.Printf("Successfully created service:\nName: %s\nCode: %s\nType: %s\nEnabled: %v\n\n",
			service.Name,
			service.Code,
			service.Type,
			service.IsEnabled,
		)
	}
}
