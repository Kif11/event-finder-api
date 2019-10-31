package meetup

import (
	"encoding/json"
	"errors"
	"eventapi/internal/common"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
)

type MeetupAPI struct{}

func msDurationToUTC(offset int) (string, error) {
	offsetString := strconv.FormatInt(int64(offset), 10)
	durationString := fmt.Sprintf("%sms", offsetString)

	d, err := time.ParseDuration(durationString)
	if err != nil {
		return "", err
	}

	hours := -d.Hours()

	return fmt.Sprintf("-%02d:00", int(hours)), nil
}

// FetchEvents from Meetup
func (api MeetupAPI) FetchEvents(config common.Config, startTime time.Time, endTime time.Time, lat float64, lon float64, distance int) ([]common.Event, error) {
	log.Info("meetup: fetching events")

	var events []common.Event

	params := url.Values{}

	params.Add("start_date_range", common.TimeToAPIString(startTime))
	params.Add("end_date_range", common.TimeToAPIString(endTime))
	params.Add("lat", fmt.Sprintf("%f", lat))
	params.Add("lon", fmt.Sprintf("%f", lon))
	params.Add("fields", "rsvp_sample,self,event_hosts,group_key_photo,group_photo_gradient,group_self_status,saved")
	params.Add("order", "best")
	params.Add("page", "80")
	params.Add("preset_date_range", "today")
	params.Add("self_groups", "include")
	params.Add("text", "")

	url := fmt.Sprintf("https://api.meetup.com/find/upcoming_events?%s", params.Encode())

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return events, err
	}

	authString := fmt.Sprintf("Bearer %s", config.MeetupToken)

	headers := map[string]string{
		"Host":                   "api.meetup.com",
		"x-meetup-locale":        "en_US:en-US",
		"accept":                 "*/*",
		"x-meetup-request-flags": "terms_of_service_interstitial_enabled_ios",
		"accept-charset":         "utf-8",
		"authorization":          authString,
		"x-meetup-agent":         "app_name=\"Meetup-iOS\", app_version=\"8.0.20\", app_version_int=\"610427\", device_name=\"iPhone10_4\", os_version=\"12.1\", os_name=\"iOS\"",
		"x-meetup-udid":          "EFBFB7D5731C43328D5A5E7CCAEEB700",
		"if-none-match":          "\"e0d13c93de88dee31789af77d666f900-gzip\"",
		"accept-language":        "en-US;q=1.0, en;q=0.9",
		"x-meta-request-headers": "unread-updates, unread-notifications, unread-messages",
		"user-agent":             "Meetup/610427 CFNetwork/975.0.3 Darwin/18.2.0",
		"x-meta-photo-host":      "secure",
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)

	if err != nil {
		return events, errors.New("meetup request failed. " + err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return events, err
	}

	if resp.StatusCode != http.StatusOK {
		var apiError ErrorResponse
		if err := json.Unmarshal([]byte(body), &apiError); err != nil {
			return events, errors.New("meetup: can not decode response. " + err.Error())
		}
		return events, errors.New("meetup request failed with message: " + apiError.Errors[0].Message)
	}

	var eventsResponse EventsResponse

	if err := json.Unmarshal([]byte(body), &eventsResponse); err != nil {
		return events, err
	}

	for _, e := range eventsResponse.Events {
		UTCOffset, err := msDurationToUTC(e.UtcOffset)
		if err != nil {
			return events, err
		}

		timeString := fmt.Sprintf("%sT%s:00%s", e.LocalDate, e.LocalTime, UTCOffset)
		time, err := time.Parse(time.RFC3339, timeString)
		if err != nil {
			return events, errors.New("meetup: " + err.Error())
		}

		event := common.Event{
			ID:   xid.New().String(),
			Name: e.Name,
			Time: time,
			Link: e.Link,
			Lat:  e.Venue.Lat,
			Lon:  e.Venue.Lon,
		}
		events = append(events, event)
	}

	return events, nil
}
