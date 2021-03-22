package request

import (
	"net/http"

	"github.com/redmejia/categories"
	"github.com/redmejia/cors"
	"github.com/redmejia/middleware"
	"github.com/redmejia/order"
	"github.com/redmejia/product"
)

// use third party channing middleware alice
var categorie = http.HandlerFunc(categories.HandleCategories)
var Catergories = middleware.Logger(middleware.Headers(cors.Cors(categorie)))

var orders = http.HandlerFunc(order.HandleOrder)
var Order = middleware.Logger(middleware.Headers(cors.Cors(orders)))

var products = http.HandlerFunc(product.HandleProducts)
var Product = middleware.Logger(middleware.Headers(cors.Cors(products)))
