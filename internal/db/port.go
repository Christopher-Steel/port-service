package db_wrapper

import (
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

func AddOrUpdate(port Port) error {
	db_wrapper.DbConnection.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&users)
	result = db_wrapper.DbConnection.Create(&Port)
	if result.error != nil {
		return result.error
	}
	return nil
}
