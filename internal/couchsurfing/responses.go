package couchsurfing

type EventsResponse []struct {
	ID               string `json:"id"`
	Priority         int    `json:"priority"`
	Title            string `json:"title"`
	StartTime        string `json:"startTime"`
	ImageURL         string `json:"imageUrl"`
	ParticipantCount int    `json:"participantCount"`
	IsFeatured       bool   `json:"isFeatured"`
	ShareLink        string `json:"shareLink"`
	Address          struct {
		ID               string      `json:"id"`
		Description      string      `json:"description"`
		StreetAddressOne string      `json:"streetAddressOne"`
		StreetAddressTwo interface{} `json:"streetAddressTwo"`
		City             string      `json:"city"`
		State            string      `json:"state"`
		Postcode         interface{} `json:"postcode"`
		Country          string      `json:"country"`
		Lat              float64     `json:"lat"`
		Lng              float64     `json:"lng"`
		PlaceID          string      `json:"placeId"`
	} `json:"address"`
}
