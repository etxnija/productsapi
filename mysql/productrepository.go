package mysql

import "github.com/etxnija/productsapi"

type productRepository struct {
}

// NewMySQLProdutRepository to access mySQL db
func NewMySQLProdutRepository() *productRepository {
	return &productRepository{}
}

func (r *productRepository) Create(p *productsapi.Product) error {
	return nil
}
func (r *productRepository) FindByID(id string) (*productsapi.Product, error) {
	return nil, nil
}

func (r *productRepository) GetAll() ([]productsapi.Product, error) {
	return []productsapi.Product{
		productsapi.Product{
			ID:    "1",
			Title: "Radio",
			Sku:   "112233",
		},
		productsapi.Product{
			ID:    "2",
			Title: "Tape",
			Sku:   "223344",
		},
		productsapi.Product{
			ID:    "3",
			Title: "Phone",
			Sku:   "334455",
		},
		productsapi.Product{
			ID:    "4",
			Title: "Computer",
			Sku:   "445566",
		},
	}, nil
}
