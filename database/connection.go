package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// DB ... database connect to pool
type DB struct {
	DataBase *sql.DB
}

const (
	OpenConns = 10
	IdleConns = 3
	LifeTime  = 60 * time.Second
)

func DbConnect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("HOSTNAME"), port, os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"), os.Getenv("SSLMODE"))
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(OpenConns)
	db.SetMaxIdleConns(IdleConns)
	db.SetConnMaxLifetime(LifeTime)

	return db, nil
}
