package meetup

// ErrorResponse from Meetup API
type ErrorResponse struct {
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

// EventsResponse object for Meetup `events`
type EventsResponse struct {
	City struct {
		ID          int     `json:"id"`
		City        string  `json:"city"`
		Lat         float64 `json:"lat"`
		Lon         float64 `json:"lon"`
		State       string  `json:"state"`
		Country     string  `json:"country"`
		Zip         string  `json:"zip"`
		MemberCount int     `json:"member_count"`
	} `json:"city"`
	Events []struct {
		Created    int64  `json:"created"`
		Duration   int    `json:"duration"`
		ID         string `json:"id"`
		Name       string `json:"name"`
		RsvpSample []struct {
			ID      int   `json:"id"`
			Created int64 `json:"created"`
			Updated int64 `json:"updated"`
			Member  struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Bio   string `json:"bio"`
				Photo struct {
					ID          int    `json:"id"`
					HighresLink string `json:"highres_link"`
					PhotoLink   string `json:"photo_link"`
					ThumbLink   string `json:"thumb_link"`
					Type        string `json:"type"`
					BaseURL     string `json:"base_url"`
				} `json:"photo"`
				Role string `json:"role"`
				Self struct {
					Blocks  bool     `json:"blocks"`
					Actions []string `json:"actions"`
					Common  struct {
						Groups []interface{} `json:"groups"`
					} `json:"common"`
					Friends bool `json:"friends"`
				} `json:"self"`
				EventContext struct {
					Host bool `json:"host"`
				} `json:"event_context"`
			} `json:"member"`
		} `json:"rsvp_sample"`
		Self struct {
			Actions []string `json:"actions"`
		} `json:"self"`
		DateInSeriesPattern bool   `json:"date_in_series_pattern"`
		Status              string `json:"status"`
		Time                int64  `json:"time"`
		LocalDate           string `json:"local_date"`
		LocalTime           string `json:"local_time"`
		Updated             int64  `json:"updated"`
		UtcOffset           int    `json:"utc_offset"`
		WaitlistCount       int    `json:"waitlist_count"`
		YesRsvpCount        int    `json:"yes_rsvp_count"`
		Venue               struct {
			ID                   int     `json:"id"`
			Name                 string  `json:"name"`
			Lat                  float64 `json:"lat"`
			Lon                  float64 `json:"lon"`
			Repinned             bool    `json:"repinned"`
			Address1             string  `json:"address_1"`
			City                 string  `json:"city"`
			Country              string  `json:"country"`
			LocalizedCountryName string  `json:"localized_country_name"`
			Zip                  string  `json:"zip"`
			State                string  `json:"state"`
		} `json:"venue"`
		Group struct {
			Created           int64   `json:"created"`
			Name              string  `json:"name"`
			ID                int     `json:"id"`
			JoinMode          string  `json:"join_mode"`
			Lat               float64 `json:"lat"`
			Lon               float64 `json:"lon"`
			Urlname           string  `json:"urlname"`
			Who               string  `json:"who"`
			LocalizedLocation string  `json:"localized_location"`
			State             string  `json:"state"`
			Country           string  `json:"country"`
			Region            string  `json:"region"`
			Timezone          string  `json:"timezone"`
			KeyPhoto          struct {
				ID          int    `json:"id"`
				HighresLink string `json:"highres_link"`
				PhotoLink   string `json:"photo_link"`
				ThumbLink   string `json:"thumb_link"`
				Type        string `json:"type"`
				BaseURL     string `json:"base_url"`
			} `json:"key_photo"`
			PhotoGradient struct {
				ID             int    `json:"id"`
				LightColor     string `json:"light_color"`
				DarkColor      string `json:"dark_color"`
				CompositeColor string `json:"composite_color"`
			} `json:"photo_gradient"`
			Self struct {
				Status string `json:"status"`
			} `json:"self"`
		} `json:"group"`
		Link         string `json:"link"`
		Description  string `json:"description"`
		HowToFindUs  string `json:"how_to_find_us"`
		Visibility   string `json:"visibility"`
		MemberPayFee bool   `json:"member_pay_fee"`
		EventHosts   []struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Intro string `json:"intro"`
			Photo struct {
				ID          int    `json:"id"`
				HighresLink string `json:"highres_link"`
				PhotoLink   string `json:"photo_link"`
				ThumbLink   string `json:"thumb_link"`
				Type        string `json:"type"`
				BaseURL     string `json:"base_url"`
			} `json:"photo"`
			Role      string `json:"role"`
			HostCount int    `json:"host_count"`
			JoinDate  int64  `json:"join_date"`
		} `json:"event_hosts"`
		Saved bool `json:"saved"`
	} `json:"events"`
}
