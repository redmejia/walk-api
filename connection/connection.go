package connection

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	OpenConns = 10
	IdleConns = 3
	LifeTime  = 60 * time.Second
)

func Dbconn() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("HOSTNAME"), port, os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"), os.Getenv("SSLMODE"))

	DB, err = sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	DB.SetMaxOpenConns(OpenConns)
	DB.SetMaxIdleConns(IdleConns)
	DB.SetConnMaxLifetime(LifeTime)

	return DB, nil
}
