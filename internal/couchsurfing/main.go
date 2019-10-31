package couchsurfing

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
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

type CouchsurfingAPI struct {
	Client common.HttpClient
}

func makeURLSignature(secret string, msg string) string {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write([]byte(msg))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

// FetchEvents from Couchsurfing
func (api CouchsurfingAPI) FetchEvents(config common.Config, startTime time.Time, endTime time.Time, lat float64, lon float64, distance int) ([]common.Event, error) {
	log.Info("couchsurfing: fetching events")

	var events []common.Event
	csURL := "https://hapi.couchsurfing.com"
	privateKey := config.CouchsurfingPrivateKey
	accessToken := config.CouchsurfingAccessToken
	UID := config.CouchsurfingUID
	secret := fmt.Sprintf("%s.%s", privateKey, UID)

	params := url.Values{}

	params.Add("page", "1")
	params.Add("perPage", "100")
	params.Add("latLng", fmt.Sprintf("%f,%f", lat, lon))

	// url := fmt.Sprintf("https://hapi.couchsurfing.com/api/v3.2/events/search?%s", params.Encode())
	path := fmt.Sprintf("/api/v3.2/events/search?%s", params.Encode())
	url := fmt.Sprintf("%s%s", csURL, path)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return events, err
	}

	signature := makeURLSignature(secret, path)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-CS-Url-Signature", signature)
	req.Header.Add("Accept-Language", "en;q=1")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("User-Agent", "Dalvik/2.1.0 (Linux; U; Android 5.0.1; Android SDK built for x86 Build/LSX66B) Couchsurfing/android/20141121013910661/Couchsurfing/3.0.1/ee6a1da")
	req.Header.Add("X-Access-Token", accessToken)

	client := api.Client

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return events, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return events, err
	}

	if resp.StatusCode != http.StatusOK {
		return events, errors.New("couchsurfing: request failed with code " + string(resp.StatusCode))
	}

	var eventsResponse EventsResponse

	if err := json.Unmarshal([]byte(body), &eventsResponse); err != nil {
		return events, err
	}

	for _, e := range eventsResponse {
		time, err := time.Parse(time.RFC3339, e.StartTime)
		if err != nil {
			return events, errors.New("couchsurfing: can not parse time string in response. " + err.Error())
		}

		event := common.Event{
			ID:   xid.New().String(),
			Name: e.Title,
			Time: time,
			Link: e.ShareLink,
			Lat:  e.Address.Lat,
			Lon:  e.Address.Lng,
		}
		events = append(events, event)
	}

	return events, nil
}
