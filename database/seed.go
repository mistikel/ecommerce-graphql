package database

import (
	"ecommerce/domain/product"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	db.AutoMigrate(&product.Variant{}, &product.Shipping{}, &product.Seo{}, &product.Product{})
	return nil
}
