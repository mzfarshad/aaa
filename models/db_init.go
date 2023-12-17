package models

import (
	"web-service-gin/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToPostgres() error {
	var err error
	psql := config.Get().Postgres()
	db, err = gorm.Open(postgres.Open(psql.DSN()), &gorm.Config{})
	if err != nil {
		return err
	}
	return migrate(db)
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Album{},
		&User{},
		// ... new models will be added
	)
}
