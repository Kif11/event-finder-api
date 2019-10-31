package main

import (
	"eventapi/internal/common"
	"eventapi/internal/database"
	"time"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func saveEvents(db *gorm.DB, events []common.Event) {
	for _, e := range events {
		var event database.Event

		if err := db.Where("link = ?", e.Link).First(&event).Error; gorm.IsRecordNotFoundError(err) {
			log.Infof("Adding new event \"%s\"\n", e.Name)
			db.Create(&database.Event{
				Name: e.Name,
				Link: e.Link,
				Lat:  e.Lat,
				Lon:  e.Lon,
				Time: e.Time,
			})
			continue
		}
		// fmt.Printf("[-] Event \"%s\" already exist in database\n", e.Name)
	}
}

func main() {
	timeInterval := 1 * time.Hour

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	db, err := database.Init()
	defer db.Close()
	if err != nil {
		log.Error("eventfetcher: failed to initialize database. " + err.Error())
	}

	ticker := time.NewTicker(timeInterval)

	for ; true; <-ticker.C {

		log.Info("Starting event fetching procedure")
		// log.Info("Starting event fetching procedure")
		// fmt.Println("[+] Starting event fetching procedure ", time.Now())

		events, err := fetchEvents()
		if err != nil {
			log.Error("eventfetcher: failed to fetch new events. " + err.Error())
		}

		saveEvents(db, events)
	}
}
