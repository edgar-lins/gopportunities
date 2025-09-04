package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize SQLite
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("failed to initialize SQLite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initialize logger
	logger = newLogger(p)
	return logger
}
