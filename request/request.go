package request

import (
	"net/http"

	"github.com/redmejia/categories"
	"github.com/redmejia/cors"
	"github.com/redmejia/makeorder"
	"github.com/redmejia/middleware"
	"github.com/redmejia/product"
)

// use third party channing middleware alice
var categorie = http.HandlerFunc(categories.Categories)
var Catergories = middleware.Logger(middleware.Headers(cors.Cors(categorie)))

var makeOrder = http.HandlerFunc(makeorder.Makeorder)
var MakeOrder = middleware.Logger(middleware.Headers(cors.Cors(makeOrder)))

var products = http.HandlerFunc(product.HandleProduct)
var Product = middleware.Logger(middleware.Headers(cors.Cors(products)))
