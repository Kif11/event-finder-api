package common

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

// EventProvider represent a single event source
// such as a particular website from which events are fetched
type EventProvider interface {
	FetchEvents(conf Config, startTime time.Time, endTime time.Time, lat float64, lon float64, distance int) ([]Event, error)
}

// Event represet single generic event
// from any service such as couchserfing or meetup
type Event struct {
	ID   string
	Name string
	Time time.Time
	Lat  float64
	Lon  float64
	Link string
}

// Config for the application
type Config struct {
	Address                 string
	Port                    int
	MeetupToken             string
	CouchsurfingPrivateKey  string
	CouchsurfingAccessToken string
	CouchsurfingUID         string
	EventbriteAuthToken     string
}

// ReadConfig file and return config struct
func ReadConfig(path string) (Config, error) {
	c := Config{}

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return c, errors.New(path + " file does not exist.")
	}

	file, _ := os.Open(path)
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&c)
	if err != nil {
		return c, errors.New("can not decode json. " + err.Error())
	}

	return c, nil
}
