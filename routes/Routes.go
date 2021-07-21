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
	var database walk.DataBase
	database.DB = connection.DB

	var storeHandlers handlers.StoreHandlers
	storeHandlers.Store = &database

	http.HandleFunc(fmt.Sprintf("%scategorie", base), middleware.Use(storeHandlers.HandleCategories, middlewares...))
	http.HandleFunc(fmt.Sprintf("%sproduct", base), middleware.Use(storeHandlers.HandleProduct, middlewares...))
	http.HandleFunc(fmt.Sprintf("%spromo", base), middleware.Use(storeHandlers.HandlerPromo, middlewares...))

	http.HandleFunc(fmt.Sprintf("%sorders", base), middleware.Use(storeHandlers.HandleOrder, middlewares...))

	http.HandleFunc(fmt.Sprintf("%sregister", base), middleware.Use(storeHandlers.HandleRegister, middlewares...))
	http.HandleFunc(fmt.Sprintf("%ssignin", base), middleware.Use(storeHandlers.HandlerSignin, middlewares...))
}
