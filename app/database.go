package app

import (
	"database/sql"
	"rest-api-native/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:pw@tcp(localhost:3306)/golang")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
