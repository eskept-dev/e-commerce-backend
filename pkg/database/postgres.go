package database

import (
	"fmt"
	"log"

	"eskept/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(databaseConfig config.DatabaseConfig) (*gorm.DB, error) {
	// Generate the database DSN
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.DBName,
		databaseConfig.SSLMode,
	)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database successfully")

	// Migrate the database
	err = Migrate(db)
	if err != nil {
		return nil, err
	}
	log.Println("Database migrated successfully")

	return db, nil
}
