package handlers

import (
	"log"

	"github.com/redmejia/walk"
)

type StoreHandler struct {
	Store  walk.Store
	Errlog *log.Logger
}
