package repository

import (
	"context"
	repo "ecommerce/app/domain/repository"
	"ecommerce/app/models"
	"log"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository ....
func NewProductRepository(db *gorm.DB) repo.ProductRepository {
	return &productRepository{
		db: db,
	}
}

var _ repo.ProductRepository = &productRepository{}

func (p *productRepository) ResolveProductByID(ctx context.Context, id int) (*models.Product, error) {
	product := &models.Product{}
	err := p.db.Find(product, id).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return product, nil
}

func (p *productRepository) ResolveProducts(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	err := p.db.Find(&products).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return products, nil
}

func (p *productRepository) StoreProduct(ctx context.Context, product *models.Product) error {
	err := p.db.Create(product).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *productRepository) UpdateProduct(ctx context.Context, id int, product *models.Product) error {
	err := p.db.Where("id = ?", id).Updates(product).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *productRepository) DeleteProductByID(ctx context.Context, id int) error {
	err := p.db.Where("id = ?", id).Delete(&models.Product{}).Error
	if err != nil {
		log.Println(err)
	}
	return err
}
