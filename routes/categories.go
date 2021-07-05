package routes

import (
	"fmt"
	"net/http"

	"github.com/redmejia/handlers"
	"github.com/redmejia/middleware"
)

func ProductCategories(base string, middlewares []middleware.Middlewares) {
	http.HandleFunc(fmt.Sprintf("%scategorie", base), middleware.Use(handlers.HandleCategories, middlewares...))
}
