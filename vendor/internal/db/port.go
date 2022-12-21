package db_wrapper

import (
	pb "api"
	db_wrapper "internal/db"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Port struct {
	gorm.Model
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates *Coordinates
	Province    string
	Timezone    string
	Unlocs      []string `gorm:"primaryKey"`
	Code        *string
}

db.Clauses(clause.OnConflict{
	UpdateAll: true,
}).Create(&users)

func AddOrUpdate(port Port) error {
	result = DbConnection.Create(&Port)
	if result.error != nil {
		return result.error
	}
	return nil
}
