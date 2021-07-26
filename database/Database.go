package database

import "database/sql"

type DataBase struct {
	Conn *sql.DB
}
