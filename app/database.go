package app

import (
	"database/sql"
	"github.com/nandes007/employee-claim-request/helper"
	"os"
	"time"
)

func NewDB() *sql.DB {
	driver := os.Getenv("DB_DRIVER")
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := username + ":" + password + "@tcp" + "(" + host + ":" + port + ")/" + dbName + "?parseTime=true"
	db, err := sql.Open(driver, connStr)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
