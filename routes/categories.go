package routes

import (
	"net/http"

	"github.com/redmejia/handlers"
	"github.com/redmejia/middleware"
)

func ProductCategories(base string, middlewares []middleware.Middlewares) {
	var walk handlers.Store
	http.HandleFunc(base+"categorie", middleware.Use(walk.HandleCategories, middlewares...))
}
