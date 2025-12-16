package database

import (
	"Users/riwandi/Documents/practice/go-restful-api/helper"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func NewDB() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" {
		dbUser = "root"
	}
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "3306"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPass, dbHost, dbPort, dbName)

	log.Println("Connecting to database with DSN:",
		fmt.Sprintf("%s:****@tcp(%s:%s)/%s", dbUser, dbHost, dbPort, dbName))

	db, err := sql.Open("mysql", dsn)
	helper.HandlePanic(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	log.Println("Database connected successfully")

	return db
}
