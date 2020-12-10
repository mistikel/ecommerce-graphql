package database

import (
	"fmt"
	"log"

	"ecommerce/app/config"
	"ecommerce/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() *gorm.DB {
	conf := config.Get()
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s DB.name=gorm database=%s sslmode=disable", conf.PostgresHost, conf.PostgresUsername, conf.PostgresPassword, conf.PostgresDatabase),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't initiate db")
	}
	autoMigrate(db)
	return db
}
func autoMigrate(db *gorm.DB) error {
	db.AutoMigrate(&models.Product{}, &models.Variant{}, &models.Shipping{}, &models.Seo{})
	return nil
}
