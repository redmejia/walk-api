package routes

import (
	"fmt"
	"net/http"

	"github.com/redmejia/handlers"
	"github.com/redmejia/middleware"
)

func Categories(base string, middlewares []middleware.Middleware) {
	// db := handlers.DBRepo{Conn: connection.DB}
	// handlerRepo := handlers.HandlerRep{DBRep: db}

	http.HandleFunc(fmt.Sprintf("%scategorie", base), middleware.Use(handlers.HandleCategories, middlewares...))
	// http.HandleFunc(fmt.Sprintf("%scategorie", base), middleware.Use(repo.HandleCategories, middlewares...))
}
