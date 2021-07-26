package handlers

import (
	"log"

	"github.com/redmejia/database"
)

type StoreHandler struct {
	Store  database.Store
	Errlog *log.Logger
}
