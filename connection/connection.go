package connection

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/redmejia/envs"
)

func Dbconn() (*sql.DB, error) {
	p := envs.DBEnv("PORT")
	port, _ := strconv.Atoi(p)
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		envs.DBEnv("HOSTNAME"),
		port,
		envs.DBEnv("USER"),
		envs.DBEnv("PASSWORD"),
		envs.DBEnv("DBNAME"),
		envs.DBEnv("SSLMODE"),
	)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
