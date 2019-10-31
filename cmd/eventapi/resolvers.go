package main

import (
	"eventapi/internal/common"

	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
)

// EventsResolver fetch events from DB
func (s *Server) EventsResolver(p graphql.ResolveParams) (interface{}, error) {
	var lat float64
	var lon float64
	var totalTime string
	var ok bool

	if lat, ok = p.Args["lat"].(float64); !ok {
		lat = 34.079701
	}
	if lon, ok = p.Args["lon"].(float64); !ok {
		lon = -118.269421
	}

	if totalTime, ok = p.Args["totalTime"].(string); !ok {
		totalTime = "168h" // 168 hours == 7 days
	}

	log.Infof("main: got event request for Lat: %f, Lon: %f, %s", lat, lon, totalTime)

	var events []common.Event
	events, err := FetchEvents(s.Config, lat, lon, totalTime)
	if err != nil {
		log.Error(err)
	}

	return events, nil
}
