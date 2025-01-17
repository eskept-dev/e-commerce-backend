package main

import (
	"eskept/internal/app"
	"eskept/pkg/config"
	"fmt"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig("./config")
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	server := app.NewServer(config)
	server.Run()
}
