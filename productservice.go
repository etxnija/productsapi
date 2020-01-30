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

// NewProductService returnm the service
func NewProductService(repo ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

// GetProducts paginated and filtered
func (s *productService) GetProducts(p Page, f Filter) ([]Product, error) {
	return s.repo.GetAll()
}

// GetProduct by id
func (s *productService) GetProduct(id int) (*Product, error) {
	return &Product{}, nil
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
