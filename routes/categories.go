package routes

import (
	"net/http"

	"github.com/redmejia/handlers"
	"github.com/redmejia/middleware"
)

func ProductCategories(base string, middlewares []middleware.Middlewares) {
	http.HandleFunc(base+"categorie", middleware.Use(handlers.HandleCategories, middlewares...))
}
