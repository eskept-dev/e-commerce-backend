package database

import (
	"eskept/internal/app/context"
	"eskept/internal/constants/enums"
	"eskept/internal/models"
	"fmt"
	"log"
)

// InitData initializes the data
func InitData(appCtx *context.AppContext) {
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
