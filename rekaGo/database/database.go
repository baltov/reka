package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"rekaGo/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database connection
func Init() {
	var err error

	var newLogger logger.Interface = logger.Default.LogMode(logger.Error)
	DB, err = gorm.Open(sqlite.Open("/mydb.sqlite"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logrus.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&models.Category{}, &models.Question{}, &models.Answer{}, &models.Image{})
	if err != nil {
		logrus.Fatalf("Failed to migrate database: %v", err)
	}

	logrus.Info("Database connected and migrated successfully.")
}
