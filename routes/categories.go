package routes

import (
	"net/http"

	"github.com/redmejia/handlers"
	"github.com/redmejia/middleware"
)

func ProductCategories(base string, middlewares []middleware.Middlewares) {
	// var cat categories.Cat
	// http.HandleFunc(base+"categorie", middleware.Use(categories.HandleCategories, middlewares...))
	// http.HandleFunc(base+"categorie", middleware.Use(cat.HandleCategories, middlewares...))
	var walk handlers.Store
	http.HandleFunc(base+"categorie", middleware.Use(walk.HandleCategories, middlewares...))
}
