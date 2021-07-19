package walk

// Store ...
type Store interface {
	GetProducts(query string) ([]Products, error)
	GetProductById(query string, productID int) ProductInfo
	// GetClientPurchaseInfoByUserId(userId int) (purchase Purchase)

	// InsertNewOrder(status PurchaseStatus) // take status pointer

	// Client(w http.ResponseWriter) // Client register and client signin
}
