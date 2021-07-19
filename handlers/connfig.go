package handlers

import (
	"github.com/redmejia/connection"
	"github.com/redmejia/walk"
)

var db walk.Store = &walk.DataBase{DB: connection.DB}
