package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/redmejia/connection"
	"github.com/redmejia/handlers"
	"github.com/redmejia/middleware"
	"github.com/redmejia/walk"
)

func Routes(base string, middlewares []middleware.Middleware) {
	var database walk.DataBase
	database.DB = connection.DB

	var logs Logers // testing
	logs.Info = log.New(os.Stdout, "INFO ", log.Ltime|log.Ldate)
	logs.Error = log.New(os.Stdout, "ERROR ", log.Ltime|log.Ldate)

	var storeHandlers handlers.StoreHandlers
	storeHandlers.Store = &database
	storeHandlers.Errlog = logs.Error

	http.HandleFunc(fmt.Sprintf("%scategorie", base), middleware.Use(storeHandlers.HandleCategories, middlewares...))
	http.HandleFunc(fmt.Sprintf("%sproduct", base), middleware.Use(storeHandlers.HandleProduct, middlewares...))
	http.HandleFunc(fmt.Sprintf("%spromo", base), middleware.Use(storeHandlers.HandlerPromo, middlewares...))

	http.HandleFunc(fmt.Sprintf("%sorders", base), middleware.Use(storeHandlers.HandleOrder, middlewares...))

	http.HandleFunc(fmt.Sprintf("%sregister", base), middleware.Use(storeHandlers.HandleRegister, middlewares...))
	http.HandleFunc(fmt.Sprintf("%ssignin", base), middleware.Use(storeHandlers.HandlerSignin, middlewares...))
}
