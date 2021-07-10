package routes

import (
	"fmt"
	"net/http"

	"github.com/redmejia/handlers"
	"github.com/redmejia/middleware"
)

func Categories(base string, middlewares []middleware.Middleware) {
	http.HandleFunc(fmt.Sprintf("%scategorie", base), middleware.Use(handlers.HandleCategories, middlewares...))
}
