package database

import (
	"rest/config"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DBURL), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}