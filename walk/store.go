package walk

// Store ...
type Store interface {
	GetProducts(query string) ([]Products, error)
	GetProductById(query string, productID int) ProductInfo
	GetClientPurchaseInfoByUserId(userId int) *Purchase

	InsertNewOrder(c *ClientOrder, status PurchaseStatus)

	// Client(w http.ResponseWriter) // Client register and client signin
}
