package repository

import (
	"context"
	repo "ecommerce/app/domain/repository"
	"ecommerce/app/models"
	"log"

	"gorm.io/gorm"
)

type shippingRepository struct {
	db *gorm.DB
}

// NewShippingRepository ....
func NewShippingRepository(db *gorm.DB) repo.ShippingRepository {
	return &shippingRepository{
		db: db,
	}
}

var _ repo.ShippingRepository = &shippingRepository{}

func (p *shippingRepository) ResolveShippingByID(ctx context.Context, id int) (*models.Shipping, error) {
	shipping := &models.Shipping{}
	err := p.db.Find(shipping, id).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return shipping, nil
}

func (p *shippingRepository) ResolveShippingByProductID(ctx context.Context, id int) (*models.Shipping, error) {
	shipping := &models.Shipping{}
	err := p.db.Where("product_id = ?", id).Find(&shipping).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return shipping, nil
}

func (p *shippingRepository) ResolveShippings(ctx context.Context) ([]*models.Shipping, error) {
	var shippings []*models.Shipping
	err := p.db.Find(&shippings).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return shippings, nil
}

func (p *shippingRepository) StoreShipping(ctx context.Context, shipping *models.Shipping) error {
	err := p.db.Create(shipping).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *shippingRepository) UpdateShipping(ctx context.Context, id int, shipping *models.Shipping) error {
	err := p.db.Where("id = ?", id).Updates(shipping).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *shippingRepository) DeleteShippingByID(ctx context.Context, id int) error {
	err := p.db.Where("id = ?", id).Delete(&models.Shipping{}).Error
	if err != nil {
		log.Println(err)
	}
	return err
}
