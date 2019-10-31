package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Import postgres dialect
	log "github.com/sirupsen/logrus"
)

// Init database
func Init() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")

	if err != nil {
		return db, err
	}

	log.Info("db: running migrations")
	db.AutoMigrate(&Event{})

	return db, nil
}
