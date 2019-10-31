package eventbrite

import (
	"encoding/json"
	"errors"
	"eventapi/internal/common"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

type EventbriteAPI struct{}

// FetchEvents from Eventbright API
func (api EventbriteAPI) FetchEvents(config common.Config, startTime time.Time, endTime time.Time, lat float64, lon float64, distance int) ([]common.Event, error) {
	log.Info("eventbrite: fetching events")

	var events []common.Event
	authToken := config.EventbriteAuthToken

	params := url.Values{}

	params.Add("expand", "venue")
	params.Add("location.latitude", fmt.Sprintf("%f", lat))
	params.Add("location.longitude", fmt.Sprintf("%f", lon))
	params.Add("start_date.range_start", common.TimeToAPIString(startTime))
	params.Add("start_date.range_end", common.TimeToAPIString(endTime))
	params.Add("location.within", fmt.Sprintf("%dkm", distance))

	url := fmt.Sprintf("https://www.eventbriteapi.com/v3/events/search?%s", params.Encode())

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return events, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	resp, err := client.Do(req)

	if err != nil {
		return events, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return events, err
	}

	if resp.StatusCode != http.StatusOK {
		var apiError ErrorResponse
		if err := json.Unmarshal([]byte(body), &apiError); err != nil {
			return events, errors.New("eventbrite: " + err.Error())
		}
		return events, errors.New(apiError.ErrorDescription)
	}

	var eventsResponse EventsResponse

	if err := json.Unmarshal([]byte(body), &eventsResponse); err != nil {
		return events, errors.New("eventbrite: " + err.Error())
	}

	for _, e := range eventsResponse.Events {
		event := common.Event{
			ID:   xid.New().String(),
			Name: e.Name.Text,
		}
		events = append(events, event)
	}

	return events, nil
}
