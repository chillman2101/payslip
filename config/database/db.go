package database

import (
	"github.com/payslip/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config *config.Config) (*gorm.DB, error) {
	dsn := config.DB

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	Migration(db)
	// Seeder(db)

	return db, nil
}
