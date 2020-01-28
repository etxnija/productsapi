package productsapi

// ProductRepository for storing and retriving products
type ProductRepository interface {
	Create(p *Product) error
	FindById(id string) (*Product, error)
}
