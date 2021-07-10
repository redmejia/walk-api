package walk

import "net/http"

// Store ...
type Store interface {
	GetProducts(query string) ([]Products, error)
	GetProductById(query string, productID int) ProductInfo
	GetClientPurchaseInfoByUserId(userId int) (purchase Purchase)

	InsertNewOrder(status *PurchaseStatus)

	NewClient(w http.ResponseWriter) // for regiter new client
}
