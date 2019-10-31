package main

import (
	"eventapi/internal/common"
	"eventapi/internal/couchsurfing"
	"eventapi/internal/eventbrite"
	"eventapi/internal/meetup"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/graphql-go/handler"
	log "github.com/sirupsen/logrus"
)

type result struct {
	Events []common.Event
	Error  error
}

func FetchEvents(config common.Config, lat float64, lon float64, totalTime string) ([]common.Event, error) {
	startTime := time.Now()
	timeInterval, _ := time.ParseDuration(totalTime)
	endTime := startTime.Add(timeInterval)

	distance := 10

	client := &http.Client{}

	providers := []common.EventProvider{
		&eventbrite.EventbriteAPI{},
		&meetup.MeetupAPI{},
		&couchsurfing.CouchsurfingAPI{
			Client: client,
		},
	}

	c := make(chan result)
	var wg sync.WaitGroup

	for _, p := range providers {
		wg.Add(1)
		go func(p common.EventProvider) {
			defer wg.Done()
			var cr result
			events, err := p.FetchEvents(config, startTime, endTime, lat, lon, distance)
			cr.Error = err
			cr.Events = events
			c <- cr
		}(p)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	var allEvents []common.Event

	for r := range c {
		if r.Error != nil {
			log.Warn(r.Error.Error())
			continue
		}
		allEvents = append(allEvents, r.Events...)
	}

	return allEvents, nil
}

func main() {
	server := Server{}
	server.InitServer()

	config := server.Config

	schema, err := BuildSchema(&server)
	if err != nil {
		log.Error(err.Error())
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", CorsMiddleware(h))

	addressString := fmt.Sprintf("%s:%s", config.Address, strconv.Itoa(config.Port))

	log.Infof("main: starting server on %s", addressString)
	log.Fatal(http.ListenAndServe(addressString, nil))
}
