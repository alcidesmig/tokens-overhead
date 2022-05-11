package gorm

import (
	"tokens-overhead/models"
	"tokens-overhead/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormDatabase struct {
	db *gorm.DB
}

func NewGormDatabase(postgresURI string) (repository.DatabaseInterface, error) {
	db, err := gorm.Open(postgres.Open(postgresURI), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(models.Request{})

	return &gormDatabase{db: db}, nil
}

func (g *gormDatabase) Save(r models.Request) error {
	return g.db.Save(&r).Error
}
