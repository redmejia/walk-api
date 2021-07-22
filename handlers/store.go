package handlers

import (
	"log"

	"github.com/redmejia/walk"
)

type StoreHandlers struct {
	Store  walk.Store
	Errlog *log.Logger
}
