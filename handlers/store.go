package handlers

import (
	"github.com/redmejia/connection"
	"github.com/redmejia/walk"
)

type StoreHandlers struct {
	Store walk.Store
}

func SettingDbHandler(conn walk.Store) StoreHandlers {
	return StoreHandlers{Store: conn}
}

var DB walk.Store = &walk.DataBase{DB: connection.DB}
