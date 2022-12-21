package db_wrapper

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

var DbConnection *Database

func Connect() error {
	db, err := gorm.Open(sqlite.Open("port.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	DbConnection = &Database{db: db}
	return nil
}
