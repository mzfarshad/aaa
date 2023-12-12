package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("aaa.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	migrate(db)
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Album{})
}
