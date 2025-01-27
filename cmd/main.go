package main

import (
	cmdServiceDatabase "eskept/cmd/service/database"
	cmdServiceLocation "eskept/cmd/service/location"
	"eskept/internal/app"
	"eskept/internal/app/context"
	"eskept/pkg/cache"
	"eskept/pkg/config"
	"eskept/pkg/database"

	"flag"
	"fmt"
)

func main() {
	// Define command-line flags
	initData := flag.Bool("init-data", false, "Initialize data")
	initLocation := flag.Bool("init-location", false, "Initialize location")
	flag.Parse()

	// Default behavior
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

	// Initialize cache connection
	cacheRedis, err := cache.InitRedis(&config.Cache)
	if err != nil {
		panic(fmt.Errorf("failed to initialize cache: %w", err))
	}

	// App context
	appContext := &context.AppContext{
		DB:    db,
		Cache: cacheRedis,
		Cfg:   config,
	}

	// Call appropriate function based on flags
	if *initData {
		runInitData(appContext)
	} else if *initLocation {
		runInitLocation(appContext)
	} else {
		server := app.NewServer(config, appContext)
		server.Run()
	}
}

func runInitData(appCtx *context.AppContext) {
	// Logic from init_data.go main function
	fmt.Println("Initializing data...")
	cmdServiceDatabase.InitData(appCtx)
	// Add the actual logic here
}

func runInitLocation(appCtx *context.AppContext) {
	// Logic from init_location.go main function
	fmt.Println("Initializing location...")
	cmdServiceLocation.InitLocation(appCtx)
	// Add the actual logic here
}
