package models

import "gorm.io/gorm"

type Input struct {
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Category       string          `json:"category"`
	SubCategory    string          `json:"subCategory"`
	Images         string          `json:"images"`
	Price          float64         `json:"price"`
	ComparePrice   float64         `json:"comparePrice"`
	CostPerItem    float64         `json:"costPerItem"`
	ShippingCost   float64         `json:"shippingCost"`
	IsChargeTax    bool            `json:"isChargeTax"`
	Sku            string          `json:"sku"`
	Quantity       int             `json:"quantity"`
	IsVariantExist bool            `json:"isVariantExist"`
	Variants       []*VariantInput `json:"variants"`
	Seo            *SeoInput       `json:"seo"`
	Shipping       *ShippingInput  `json:"shipping"`
}

type Product struct {
	gorm.Model
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Category       string     `json:"category"`
	SubCategory    string     `json:"subCategory"`
	Images         string     `json:"images"`
	Price          float64    `json:"price"`
	ComparePrice   float64    `json:"comparePrice"`
	CostPerItem    float64    `json:"costPerItem"`
	ShippingCost   float64    `json:"shippingCost"`
	IsChargeTax    bool       `json:"isChargeTax"`
	Sku            string     `json:"sku"`
	Quantity       int        `json:"quantity"`
	Shipping       *Shipping  `json:"shipping"`
	Seo            *Seo       `json:"seo"`
	IsVariantExist bool       `json:"isVariantExist"`
	Variant        []*Variant `json:"variant"`
}

type Seo struct {
	gorm.Model
	ID          int    `json:"id"`
	ProductID   int    `json:"productId"`
	Auto        bool   `json:"auto"`
	PageTitle   string `json:"pageTitle"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

type SeoInput struct {
	Auto        bool   `json:"auto"`
	PageTitle   string `json:"pageTitle"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

type Shipping struct {
	gorm.Model
	ID        int     `json:"id"`
	ProductID int     `json:"productId"`
	Weight    float64 `json:"weight"`
	Rate      string  `json:"rate"`
}

type ShippingInput struct {
	Weight float64 `json:"weight"`
	Rate   string  `json:"rate"`
}

type Variant struct {
	gorm.Model
	ID         int     `json:"id"`
	ProductID  int     `json:"productId"`
	SkuVariant string  `json:"skuVariant"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}

type VariantInput struct {
	SkuVariant string  `json:"skuVariant"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
}
