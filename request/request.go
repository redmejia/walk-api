package request

import (
	"net/http"

	"github.com/redmejia/categories"
	"github.com/redmejia/clients"
	"github.com/redmejia/cors"
	"github.com/redmejia/middleware"
	"github.com/redmejia/order"
	"github.com/redmejia/product"
	"github.com/redmejia/promotion"
)

// this will change or remove
// use third party channing middleware alice
var categorie = http.HandlerFunc(categories.HandleCategories)
var Catergories = middleware.Logger(middleware.Headers(cors.Cors(categorie)))

var orders = http.HandlerFunc(order.HandleOrder)
var Order = middleware.Logger(middleware.Headers(cors.Cors(orders)))

var products = http.HandlerFunc(product.HandleProduct)
var Product = middleware.Logger(middleware.Headers(cors.Cors(products)))

var register = http.HandlerFunc(clients.HandleRegister)
var Register = middleware.Logger(middleware.Headers(cors.Cors(register)))

var signin = http.HandlerFunc(clients.HandlerSignin)
var Signin = middleware.Logger(middleware.Headers(cors.Cors(signin)))

// var promo = http.HandlerFunc(promotion.HandlerPromos)
// var Promos = middleware.Logger(middleware.Headers(cors.Cors(promo)))

var promoByid = http.HandlerFunc(promotion.HandlerPromo)
var Promo = middleware.Logger(middleware.Headers(cors.Cors(promoByid)))
