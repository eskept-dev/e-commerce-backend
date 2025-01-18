package context

import (
	"eskept/pkg/cache"
	"eskept/pkg/config"

	"gorm.io/gorm"
)

// AppContext holds shared resources for the application
type AppContext struct {
	DB    *gorm.DB
	Cache *cache.RedisClient
	Cfg   *config.Config
}

// NewAppContext initializes and returns a new AppContext
func NewAppContext(
	db *gorm.DB,
	cache *cache.RedisClient,
	cfg *config.Config,
) *AppContext {
	return &AppContext{
		DB:    db,
		Cache: cache,
		Cfg:   cfg,
	}
}
