package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/redmejia/cors"
	"github.com/redmejia/database"
	"github.com/redmejia/handlers"
	"github.com/redmejia/logs"
	"github.com/redmejia/middleware"
)

func Routes() {

	const base string = "/v1/"

	middlewares := middleware.SetMiddleware(middleware.Headers, middleware.Logger, cors.Cors)

	var db database.DataBase
	db.Conn = database.DB

	var logs logs.Logers
	logs.Info = log.New(os.Stdout, "INFO ", log.Ltime|log.Ldate)
	logs.Error = log.New(os.Stdout, "ERROR ", log.Ltime|log.Ldate)

	var storeHandlers handlers.StoreHandler
	storeHandlers.Store = &db
	storeHandlers.Errlog = logs.Error

	http.HandleFunc(fmt.Sprintf("%scategorie", base), middleware.Use(storeHandlers.HandleCategories, middlewares.Middle...))
	http.HandleFunc(fmt.Sprintf("%sproduct", base), middleware.Use(storeHandlers.HandleProduct, middlewares.Middle...))
	http.HandleFunc(fmt.Sprintf("%spromo", base), middleware.Use(storeHandlers.HandlerPromo, middlewares.Middle...))

	http.HandleFunc(fmt.Sprintf("%sorders", base), middleware.Use(storeHandlers.HandleOrder, middlewares.Middle...))

	http.HandleFunc(fmt.Sprintf("%sregister", base), middleware.Use(storeHandlers.HandleRegister, middlewares.Middle...))
	http.HandleFunc(fmt.Sprintf("%ssignin", base), middleware.Use(storeHandlers.HandlerSignin, middlewares.Middle...))

	// file serv
	var fs = http.FileServer(http.Dir(os.Getenv("PIC_PATH_DIR")))
	http.Handle(base+"img/", http.StripPrefix(base+"img/", fs))

}
