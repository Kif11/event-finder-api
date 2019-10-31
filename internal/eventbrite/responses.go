package eventbrite

import "time"

// ErrorResponse API error response
type ErrorResponse struct {
	StatusCode       int    `json:"status_code"`
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`
}

// EventsResponse object for eventbrite `events`
type EventsResponse struct {
	Pagination struct {
		ObjectCount  int  `json:"object_count"`
		PageNumber   int  `json:"page_number"`
		PageSize     int  `json:"page_size"`
		PageCount    int  `json:"page_count"`
		HasMoreItems bool `json:"has_more_items"`
	} `json:"pagination"`
	Events []struct {
		Name struct {
			Text string `json:"text"`
			HTML string `json:"html"`
		} `json:"name"`
		Description struct {
			Text string `json:"text"`
			HTML string `json:"html"`
		} `json:"description"`
		ID    string `json:"id"`
		URL   string `json:"url"`
		Start struct {
			Timezone string    `json:"timezone"`
			Local    string    `json:"local"`
			Utc      time.Time `json:"utc"`
		} `json:"start"`
		End struct {
			Timezone string    `json:"timezone"`
			Local    string    `json:"local"`
			Utc      time.Time `json:"utc"`
		} `json:"end"`
		OrganizationID               string      `json:"organization_id"`
		Created                      time.Time   `json:"created"`
		Changed                      time.Time   `json:"changed"`
		Published                    time.Time   `json:"published"`
		Capacity                     interface{} `json:"capacity"`
		CapacityIsCustom             interface{} `json:"capacity_is_custom"`
		Status                       string      `json:"status"`
		Currency                     string      `json:"currency"`
		Listed                       bool        `json:"listed"`
		Shareable                    bool        `json:"shareable"`
		OnlineEvent                  bool        `json:"online_event"`
		TxTimeLimit                  int         `json:"tx_time_limit"`
		HideStartDate                bool        `json:"hide_start_date"`
		HideEndDate                  bool        `json:"hide_end_date"`
		Locale                       string      `json:"locale"`
		IsLocked                     bool        `json:"is_locked"`
		PrivacySetting               string      `json:"privacy_setting"`
		IsSeries                     bool        `json:"is_series"`
		IsSeriesParent               bool        `json:"is_series_parent"`
		InventoryType                string      `json:"inventory_type"`
		IsReservedSeating            bool        `json:"is_reserved_seating"`
		ShowPickASeat                bool        `json:"show_pick_a_seat"`
		ShowSeatmapThumbnail         bool        `json:"show_seatmap_thumbnail"`
		ShowColorsInSeatmapThumbnail bool        `json:"show_colors_in_seatmap_thumbnail"`
		Source                       string      `json:"source"`
		IsFree                       bool        `json:"is_free"`
		Version                      string      `json:"version"`
		Summary                      string      `json:"summary"`
		LogoID                       string      `json:"logo_id"`
		OrganizerID                  string      `json:"organizer_id"`
		VenueID                      string      `json:"venue_id"`
		CategoryID                   string      `json:"category_id"`
		SubcategoryID                interface{} `json:"subcategory_id"`
		FormatID                     string      `json:"format_id"`
		ResourceURI                  string      `json:"resource_uri"`
		IsExternallyTicketed         bool        `json:"is_externally_ticketed"`
		Venue                        struct {
			Address struct {
				Address1                         string      `json:"address_1"`
				Address2                         interface{} `json:"address_2"`
				City                             string      `json:"city"`
				Region                           string      `json:"region"`
				PostalCode                       string      `json:"postal_code"`
				Country                          string      `json:"country"`
				Latitude                         string      `json:"latitude"`
				Longitude                        string      `json:"longitude"`
				LocalizedAddressDisplay          string      `json:"localized_address_display"`
				LocalizedAreaDisplay             string      `json:"localized_area_display"`
				LocalizedMultiLineAddressDisplay []string    `json:"localized_multi_line_address_display"`
			} `json:"address"`
			ResourceURI    string      `json:"resource_uri"`
			ID             string      `json:"id"`
			AgeRestriction interface{} `json:"age_restriction"`
			Capacity       interface{} `json:"capacity"`
			Name           string      `json:"name"`
			Latitude       string      `json:"latitude"`
			Longitude      string      `json:"longitude"`
		} `json:"venue"`
		Logo struct {
			CropMask struct {
				TopLeft struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"top_left"`
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"crop_mask"`
			Original struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"original"`
			ID           string `json:"id"`
			URL          string `json:"url"`
			AspectRatio  string `json:"aspect_ratio"`
			EdgeColor    string `json:"edge_color"`
			EdgeColorSet bool   `json:"edge_color_set"`
		} `json:"logo"`
	} `json:"events"`
}
