package context

import (
	"eskept/pkg/config"

	"gorm.io/gorm"
)

// AppContext holds shared resources for the application
type AppContext struct {
	DB  *gorm.DB
	Cfg *config.Config
}

// NewAppContext initializes and returns a new AppContext
func NewAppContext(
	db *gorm.DB,
	cfg *config.Config,
) *AppContext {
	return &AppContext{
		DB:  db,
		Cfg: cfg,
	}
}
