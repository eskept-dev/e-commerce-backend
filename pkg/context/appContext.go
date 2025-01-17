package context

import (
	"gorm.io/gorm"
)

// AppContext holds shared resources for the application
type AppContext struct {
	DB *gorm.DB
}

// NewAppContext initializes and returns a new AppContext
func NewAppContext(db *gorm.DB) *AppContext {
	return &AppContext{
		DB: db,
	}
}
