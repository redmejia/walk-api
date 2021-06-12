package connection

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var DB *sql.DB

const (
	openConns = 10
	idleConns = 3
	lifeTime  = 60 * time.Second
)

func Dbconn() (*sql.DB, error) {
	_ = godotenv.Load()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("HOSTNAME"),
		port,
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
		os.Getenv("SSLMODE"),
	)
	DB, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	DB.SetMaxOpenConns(openConns)
	DB.SetMaxIdleConns(idleConns)
	DB.SetConnMaxLifetime(lifeTime)
	return DB, nil
}
