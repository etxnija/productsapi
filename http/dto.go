package http

// ProductDto for the http response
type ProductDto struct {
	id    string `json:productId`
	title string `json:title`
	sku   string `json:sku`
}
