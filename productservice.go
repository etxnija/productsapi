package productsapi

// ProductService for managing products
type ProductService interface {
	GetProducts(p Page, f Filter) ([]Product, error)
	GetProduct(id int) (*Product, error)
	CreateProduct(p Product) (Product, error)
	UpdateProduct(p Product) error
}

type productService struct {
	repo ProductRepository
}

// GetProducts paginated and filtered
func (s *productService) GetProducts(p Page, f Filter) []Product {
	return nil
}

// GetProduct by id
func (s *productService) GetProduct(id int) *Product {
	return &Product{}
}

// CreateProduct in storage
func (s *productService) CreateProduct(p Product) (Product, error) {
	return Product{}, nil
}

// UpdateProduct in storage
func (s *productService) UpdateProduct(p Product) error {
	return nil
}

// Page for paginatyion
type Page struct {
	start int
	num   int
}

// Filter for products
type Filter struct {
	sku     string
	barcode string
}
