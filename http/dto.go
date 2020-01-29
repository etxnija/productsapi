package http

// ProductDto for the http response
type ProductDto struct {
	Id    string `json:productId`
	Title string `json:title`
	Sku   string `json:sku`
}
