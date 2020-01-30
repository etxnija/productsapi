package productsapi

// ProductRepository for storing and retriving products
type ProductRepository interface {
	Create(p *Product) error
	FindByID(id string) (*Product, error)
	GetAll() ([]Product, error)
}
