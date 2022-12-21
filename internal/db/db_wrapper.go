package db_wrapper

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

var dbConnection *Database

func connect() error {
	db, err := gorm.Open(sqlite.Open("port.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	dbConnection = &Database{db: db}
	return nil
}
