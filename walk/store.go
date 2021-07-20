package walk

import "net/http"

// Store ...
type Store interface {
	GetProducts(string) ([]Products, error)
	GetProductById(string, int) ProductInfo
	GetClientPurchaseInfoByUserId(int) *Purchase

	InsertNewOrder(*ClientOrder, PurchaseStatus)

	ClientRegister(*ClientRegister, http.ResponseWriter)
	ClientSiging(*ClientSignin, http.ResponseWriter)
}
