package tests

import (
	"bytes"
	"eventapi/internal/couchsurfing"
	"eventapi/internal/meetup"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBufferString("[]")),
	}, nil
}

var startTime = time.Now()
var timeInterval, _ = time.ParseDuration("168h")
var endTime = startTime.Add(timeInterval)

var lat = 34.079701
var lon = -118.269421
var distance = 10

func TestMeetupFetchEvents(t *testing.T) {
	meetup := meetup.MeetupAPI{}
	_, err := meetup.FetchEvents(startTime, endTime, lat, lon, distance)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
}

func TestCouchsurfingpFetchEvents(t *testing.T) {
	client := &ClientMock{}

	cs := couchsurfing.CouchsurfingAPI{
		Client: client,
	}
	events, err := cs.FetchEvents(startTime, endTime, lat, lon, distance)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if len(events) != 0 {
		t.Errorf("expected 0 event got %d", len(events))
	}
}
