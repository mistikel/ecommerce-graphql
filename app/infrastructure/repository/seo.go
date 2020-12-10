package repository

import (
	"context"
	repo "ecommerce/app/domain/repository"
	"ecommerce/app/models"
	"log"

	"gorm.io/gorm"
)

type seoRepository struct {
	db *gorm.DB
}

// NewSeoRepository ....
func NewSeoRepository(db *gorm.DB) repo.SeoRepository {
	return &seoRepository{
		db: db,
	}
}

var _ repo.SeoRepository = &seoRepository{}

func (p *seoRepository) ResolveSeoByID(ctx context.Context, id int) (*models.Seo, error) {
	seo := &models.Seo{}
	err := p.db.Find(seo, id).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return seo, nil
}

func (p *seoRepository) ResolveSeoByProductID(ctx context.Context, id int) (*models.Seo, error) {
	seo := &models.Seo{}
	err := p.db.Where("product_id = ?", id).Find(&seo).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return seo, nil
}

func (p *seoRepository) ResolveSeos(ctx context.Context) ([]*models.Seo, error) {
	var seos []*models.Seo
	err := p.db.Find(&seos).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return seos, nil
}

func (p *seoRepository) StoreSeo(ctx context.Context, seo *models.Seo) error {
	err := p.db.Create(seo).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *seoRepository) UpdateSeo(ctx context.Context, id int, seo *models.Seo) error {
	err := p.db.Where("id = ?", id).Updates(seo).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *seoRepository) DeleteSeoByID(ctx context.Context, id int) error {
	err := p.db.Where("id = ?", id).Delete(&models.Seo{}).Error
	if err != nil {
		log.Println(err)
	}
	return err
}
