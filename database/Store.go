package database

import (
	"net/http"

	"github.com/redmejia/walk"
)

// Store ...
type Store interface {
	GetProducts(string) ([]walk.Products, error)
	GetProductById(string, int) walk.ProductInfo
	GetClientPurchaseInfoByUserId(int) *walk.Purchase

	InsertNewOrder(*walk.ClientOrder, walk.PurchaseStatus)

	ClientRegister(*walk.ClientRegister, http.ResponseWriter)
	ClientSiging(*walk.ClientSignin, http.ResponseWriter)

	DeleteAndRefound(purchaseId string) *walk.OrderRefound
}
