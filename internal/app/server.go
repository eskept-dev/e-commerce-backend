package app

import (
	"eskept/internal/app/context"
	"eskept/internal/app/routes"
	"eskept/pkg/config"
	"fmt"
)

// Server represents the server
type Server struct {
	config *config.Config
	appCtx *context.AppContext
}

// NewServer returns a new Server instance
func NewServer(config *config.Config, appCtx *context.AppContext) *Server {
	return &Server{config: config, appCtx: appCtx}
}

// Run starts the server
func (s *Server) Run() {
	// Initialize router
	router := &routes.Router{}
	err := router.NewRouter(s.appCtx)
	if err != nil {
		panic(fmt.Errorf("failed to initialize router: %w", err))
	}

	// Setup router
	router.SetupRouter(s.appCtx)

	// Start the server
	router.Run(fmt.Sprintf(":%d", s.config.Server.Port))
}
