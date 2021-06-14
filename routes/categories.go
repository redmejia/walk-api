package routes

import (
	"net/http"

	"github.com/redmejia/categories"
	"github.com/redmejia/middleware"
)

func ProductCategories(base string, middlewares []middleware.Middlewares) {
	http.HandleFunc(base+"categorie", middleware.Use(categories.HandleCategories, middlewares...))
}
