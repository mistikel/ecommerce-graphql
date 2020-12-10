package models

// ToProduct ...
func (i *Input) ToProduct() *Product {
	return &Product{
		Name:           i.Name,
		Description:    i.Description,
		Category:       i.Category,
		SubCategory:    i.SubCategory,
		Images:         i.Images,
		Price:          i.Price,
		ComparePrice:   i.ComparePrice,
		CostPerItem:    i.CostPerItem,
		ShippingCost:   i.ShippingCost,
		IsChargeTax:    i.IsChargeTax,
		Sku:            i.Sku,
		Quantity:       i.Quantity,
		IsVariantExist: i.IsVariantExist,
	}
}

// ToVariant ...
func (v *VariantInput) ToVariant(productID int) *Variant {
	return &Variant{
		ProductID:  productID,
		SkuVariant: v.SkuVariant,
		Price:      v.Price,
		Quantity:   v.Quantity,
	}

}

// ToShipping ...
func (s *ShippingInput) ToShipping(productID int) *Shipping {
	return &Shipping{
		ProductID: productID,
		Weight:    s.Weight,
		Rate:      s.Rate,
	}
}

// ToSeo ...
func (s *SeoInput) ToSeo(productID int) *Seo {
	return &Seo{
		ProductID:   productID,
		Auto:        s.Auto,
		PageTitle:   s.PageTitle,
		Description: s.Description,
		Tags:        s.Tags,
	}
}
