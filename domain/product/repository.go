package product

var products = []Product{
	{
		ID:    1,
		Name:  "Chicha Morada",
		Info:  "Chicha morada is a beverage originated in the Andean regions of Perú but is actually consumed at a national level (wiki)",
		Price: 7.99,
	},
	{
		ID:    2,
		Name:  "Chicha de jora",
		Info:  "Chicha de jora is a corn beer chicha prepared by germinating maize, extracting the malt sugars, boiling the wort, and fermenting it in large vessels (traditionally huge earthenware vats) for several days (wiki)",
		Price: 5.95,
	},
	{
		ID:    3,
		Name:  "Pisco",
		Info:  "Pisco is a colorless or yellowish-to-amber colored brandy produced in winemaking regions of Peru and Chile (wiki)",
		Price: 9.95,
	},
}

// Repository ...
type Repository struct{}

// newRepository ...
func newRepository() *Service {
	return &Service{}
}

// GetProductByID ...
func (r *Repository) GetProductByID(id int) (Product, error) {
	for _, product := range products {
		if int(product.ID) == id {
			return product, nil
		}
	}
	return Product{}, nil
}

// GetProducts ...
func (r *Repository) GetProducts() ([]Product, error) {
	return products, nil
}

// CreateProduct ...
func (r *Repository) CreateProduct(product Product) (Product, error) {
	products = append(products, product)
	return product, nil
}

// UpdateProduct ...
func (r *Repository) UpdateProduct(product Product) (Product, error) {
	for i, p := range products {
		if p.ID == product.ID {
			products[i].Name = product.Name
			products[i].Info = product.Info
			products[i].Price = product.Price
		}
	}
	return product, nil
}

// DeleteProduct ...
func (r *Repository) DeleteProduct(id int) (Product, error) {
	product := Product{}
	for i, p := range products {
		if int64(id) == p.ID {
			product = products[i]
			products = append(products[:i], products[i+1:]...)
		}
	}
	return product, nil
}
