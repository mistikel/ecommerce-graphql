package database

import (
	"ecommerce/domain/product"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	db.AutoMigrate(&product.Product{}, &product.Variant{}, &product.Shipping{}, &product.Seo{})
	return nil
}
