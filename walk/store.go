package walk

// Store ...
type Store interface {
	GetProducts(query string) ([]Products, error)
	GetProductById(query string, productID int) ProductInfo

	NewOrder(status *PurchaseStatus)
}
