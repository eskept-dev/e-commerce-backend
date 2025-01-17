package app

import (
	"eskept/internal/app/context"
	"eskept/internal/app/routes"
	"eskept/pkg/config"
	"eskept/pkg/database"
	"fmt"
)

// Server represents the server
type Server struct {
	config *config.Config
}

// NewServer returns a new Server instance
func NewServer(config *config.Config) *Server {
	return &Server{config: config}
}

// Run starts the server
func (s *Server) Run() {
	// Initialize database connection
	db, err := database.InitPostgres(s.config.Database)
	if err != nil {
		panic(fmt.Errorf("failed to initialize database: %w", err))
	}

	// App context
	AppContext := &context.AppContext{
		DB: db,
	}

	// Initialize router
	router := &routes.Router{}
	err = router.NewRouter(AppContext)
	if err != nil {
		panic(fmt.Errorf("failed to initialize router: %w", err))
	}

	// Setup router
	router.SetupRouter(AppContext)

	// Start the server
	router.Run(fmt.Sprintf(":%s", s.config.Server.Port))
}
