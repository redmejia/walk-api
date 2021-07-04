package walk

// IWalk ...
type IWalk interface {
	GetProducts(query string) ([]Products, error)
	GetProductById(query string, productID int) ProductInfo

	NewOrder(status *PurchaseStatus)
}
