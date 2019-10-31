package database

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Event struct {
	gorm.Model
	Name string
	Link string
	Time time.Time
	Lat  float64
	Lon  float64
}
