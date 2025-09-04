package config

import (
	"os"

	"github.com/edgar-lins/gopportunities.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		// Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating database directory: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Error creating database file: %v", err)
			return nil, err
		}
		// Close the file immediately
		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Sqlite opening error: %v", err)
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Sqlite automigrate error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}
