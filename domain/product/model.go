package product

import (
	"math/rand"

	"gorm.io/gorm"
)

// Product ...
type Product struct {
	gorm.Model
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Category       string    `json:"category"`
	SubCategory    string    `json:"subCategory"`
	Images         string    `json:"images"`
	Price          float64   `json:"price"`
	ComparePrice   float64   `json:"comparePrice"`
	CostPerItem    float64   `json:"costPerItem"`
	ShippingCost   float64   `json:"shippingCost"`
	IsChargeTax    bool      `json:"isChargeTax"`
	SKU            string    `json:"sku"`
	Quantity       int       `json:"quantity"`
	Shipping       Shipping  `json:"shipping"`
	IsVariantExist bool      `json:"isVariantExist"`
	Variants       []Variant `json:"variant"`
	Seo            Seo       `json:"seo"`
}

// Input ...
type Input struct {
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Category       string         `json:"category"`
	SubCategory    string         `json:"subCategory"`
	Images         string         `json:"images"`
	Price          float64        `json:"price"`
	ComparePrice   float64        `json:"comparePrice"`
	CostPerItem    float64        `json:"costPerItem"`
	ShippingCost   float64        `json:"shippingCost"`
	IsChargeTax    bool           `json:"isChargeTax"`
	SKU            string         `json:"sku"`
	Quantity       int            `json:"quantity"`
	Shipping       ShippingInput  `json:"shipping"`
	IsVariantExist bool           `json:"isVariantExist"`
	Variants       []VariantInput `json:"variant"`
	Seo            SeoInput       `json:"seo"`
}

// Shipping ...
type Shipping struct {
	gorm.Model
	ID        int     `json:"id"`
	ProductID int     `json:"productId"`
	Weight    float64 `json:"weight"`
	Rate      string  `json:"rate"`
}

// ShippingInput ...
type ShippingInput struct {
	Weight float64 `json:"weight"`
	Rate   string  `json:"rate"`
}

// Variant ...
type Variant struct {
	gorm.Model
	ID         int     `json:"id"`
	ProductID  int     `json:"productId"`
	SKUVariant string  `json:"skuVariant"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

// VariantInput ...
type VariantInput struct {
	SKUVariant string  `json:"skuVariant"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

// Seo ...
type Seo struct {
	gorm.Model
	ID          int    `json:"id"`
	ProductID   int    `json:"productId"`
	Auto        bool   `json:"auto"`
	PageTitle   string `json:"pageTitle"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// SeoInput ...
type SeoInput struct {
	Auto        bool   `json:"auto"`
	PageTitle   string `json:"pageTitle"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

// ToProduct ...
func (i *Input) ToProduct(id *int) Product {
	var p Product
	p.ID = int(rand.Intn(100000))
	if id != nil {
		p.ID = *id
	}
	p.Name = i.Name
	p.Description = i.Description
	p.Category = i.Category
	p.SubCategory = i.SubCategory
	p.Images = i.Images
	p.Price = i.Price
	p.ComparePrice = i.ComparePrice
	p.CostPerItem = i.CostPerItem
	p.ShippingCost = i.ShippingCost
	p.IsChargeTax = i.IsChargeTax
	p.SKU = i.SKU
	p.Quantity = i.Quantity
	p.Shipping = i.Shipping.toShipping(p.ID)
	p.IsVariantExist = i.IsVariantExist
	if i.IsVariantExist {
		p.IsVariantExist = i.IsVariantExist
		for _, v := range i.Variants {
			p.Variants = append(p.Variants, v.toVariant(p.ID))
		}
	}
	p.Seo = i.Seo.toSeo(p.ID)
	return p
}

func (v *VariantInput) toVariant(productID int) Variant {
	var variant Variant
	variant.ID = int(rand.Intn(100000))
	variant.ProductID = productID
	variant.SKUVariant = v.SKUVariant
	variant.Price = v.Price
	variant.Quantity = v.Quantity
	return variant
}

func (s *ShippingInput) toShipping(productID int) Shipping {
	var shipping Shipping
	shipping.ID = int(rand.Intn(100000))
	shipping.ProductID = productID
	shipping.Weight = s.Weight
	shipping.Rate = s.Rate
	return shipping
}

func (s *SeoInput) toSeo(productID int) Seo {
	var seo Seo
	seo.ID = int(rand.Intn(100000))
	seo.ProductID = productID
	seo.Auto = s.Auto
	seo.PageTitle = s.PageTitle
	seo.Description = s.Description
	seo.Tags = s.Tags
	return seo
}
