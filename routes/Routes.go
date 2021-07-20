package routes

import (
	"fmt"
	"net/http"

	"github.com/redmejia/connection"
	"github.com/redmejia/handlers"
	"github.com/redmejia/middleware"
	"github.com/redmejia/walk"
)

func Routes(base string, middlewares []middleware.Middleware) {
	// db := handlers.DBRepo{Conn: connection.DB}
	// handlerRepo := handlers.HandlerRep{DBRep: db}

	var database walk.DataBase
	database.DB = connection.DB

	var storeHandlers handlers.StoreHandlers
	storeHandlers.Store = &database

	http.HandleFunc(fmt.Sprintf("%scategorie", base), middleware.Use(storeHandlers.HandleCategories, middlewares...))
	http.HandleFunc(fmt.Sprintf("%sproduct", base), middleware.Use(storeHandlers.HandleProduct, middlewares...))
	// http.HandleFunc(fmt.Sprintf("%scategorie", base), middleware.Use(storeHandlers.HandleCategories, middlewares...))
}
