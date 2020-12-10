package models

// ToProduct ...
func (i *Input) ToProduct() (product Product) {
	product.Name = i.Name
	product.Description = i.Description
	product.Category = i.Category
	product.SubCategory = i.SubCategory
	product.Images = i.Images
	product.Price = i.Price
	product.ComparePrice = i.ComparePrice
	product.CostPerItem = i.CostPerItem
	product.ShippingCost = i.ShippingCost
	product.IsChargeTax = i.IsChargeTax
	product.Sku = i.Sku
	product.Quantity = i.Quantity
	product.IsVariantExist = i.IsVariantExist
	return
}

func (v *VariantInput) toVariant(productID int) *Variant {
	return &Variant{
		ProductID:  productID,
		SkuVariant: v.SkuVariant,
		Price:      v.Price,
		Quantity:   v.Quantity,
	}

}

func (s *ShippingInput) toShipping(productID int) *Shipping {
	return &Shipping{
		ProductID: productID,
		Weight:    s.Weight,
		Rate:      s.Rate,
	}
}

func (s *SeoInput) toSeo(productID int) *Seo {
	return &Seo{
		ProductID:   productID,
		Auto:        s.Auto,
		PageTitle:   s.PageTitle,
		Description: s.Description,
		Tags:        s.Tags,
	}
}
