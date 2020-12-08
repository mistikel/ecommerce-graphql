package product

// Product ...
type Product struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Category       string    `json:"category"`
	SubCategory    string    `json:"subCategory"`
	Images         []string  `json:"images"`
	Price          float64   `json:"price"`
	ComparePrice   float64   `json:"comparePrice"`
	CostPerItem    float64   `json:"costPerItem"`
	ShippingCost   float64   `json:"shippingCost"`
	IsChargeTax    bool      `json:"isChargeTax"`
	SKU            string    `json:"sku"`
	Quantity       int64     `json:"quantity"`
	Shipping       Shipping  `json:"shipping"`
	IsVariantExist bool      `json:"isVariantExist"`
	Variant        []Variant `json:"variant"`
	Seo            Seo       `json:"seo"`
}

// Shipping ...
type Shipping struct {
	Weight float64 `json:"weight"`
	Rate   string  `json:"rate"`
}

// Variant ...
type Variant struct {
	ID         int64   `json:"id"`
	SKUVariant string  `json:"skuVariant"`
	Price      float64 `json:"price"`
	Quantity   int64   `json:"quantity"`
}

// Seo ...
type Seo struct {
	Auto        bool     `json:"auto"`
	PageTitle   string   `json:"pageTitle"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
