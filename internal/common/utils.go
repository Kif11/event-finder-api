package common

import (
	"fmt"
	"net/http"
	"time"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// TimeToAPIString convert time to a YYYY-MM-DDThh:mm:ss string consumed by Meetup and Eventbrite APIs
func TimeToAPIString(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", // YYYY-MM-DDThh:mm:ss
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
