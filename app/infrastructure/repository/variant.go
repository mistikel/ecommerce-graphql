package repository

import (
	"context"
	repo "ecommerce/app/domain/repository"
	"ecommerce/app/models"
	"log"

	"gorm.io/gorm"
)

type variantRepository struct {
	db *gorm.DB
}

// NewVariantRepository ....
func NewVariantRepository(db *gorm.DB) repo.VariantRepository {
	return &variantRepository{
		db: db,
	}
}

var _ repo.VariantRepository = &variantRepository{}

func (p *variantRepository) ResolveVariantByID(ctx context.Context, id int) (*models.Variant, error) {
	variant := &models.Variant{}
	err := p.db.Find(variant, id).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return variant, nil
}

func (p *variantRepository) ResolveVariantByProductID(ctx context.Context, id int) ([]*models.Variant, error) {
	var variants []*models.Variant
	err := p.db.Where("product_id = ?", id).Find(&variants).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return variants, nil
}

func (p *variantRepository) ResolveVariants(ctx context.Context) ([]*models.Variant, error) {
	var variants []*models.Variant
	err := p.db.Find(&variants).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return variants, nil
}

func (p *variantRepository) StoreVariant(ctx context.Context, variant *models.Variant) error {
	err := p.db.Create(variant).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *variantRepository) UpdateVariant(ctx context.Context, id int, variant *models.Variant) error {
	err := p.db.Where("id = ?", id).Updates(variant).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *variantRepository) DeleteVariantByID(ctx context.Context, id int) error {
	err := p.db.Where("id = ?", id).Delete(&models.Variant{}).Error
	if err != nil {
		log.Println(err)
	}
	return err
}
