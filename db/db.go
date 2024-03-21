package db

import (
	"github.com/Nishad4140/SkillSync_UserService/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(connectTo string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectTo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(
		&entities.Client{},
		&entities.ClientProfile{},
		&entities.Freelancer{},
		&entities.Category{},
	)
	return db, nil
}
