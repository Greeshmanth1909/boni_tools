package main

type RespStruct struct {
	Meta struct {
		Count       int `json:"count"`
		CurrentPage int `json:"current_page"`
	} `json:"meta"`
	Payload []struct {
		AdditionalAttributes struct {
		} `json:"additional_attributes"`
		AvailabilityStatus string      `json:"availability_status"`
		Email              interface{} `json:"email"`
		ID                 int         `json:"id"`
		Name               string      `json:"name"`
		PhoneNumber        string      `json:"phone_number"`
		Identifier         interface{} `json:"identifier"`
		Thumbnail          string      `json:"thumbnail"`
		CustomAttributes   struct {
		} `json:"custom_attributes"`
		CreatedAt      int `json:"created_at"`
		ContactInboxes []struct {
			SourceID string `json:"source_id"`
			Inbox    struct {
				ID                   int         `json:"id"`
				AvatarURL            string      `json:"avatar_url"`
				ChannelID            int         `json:"channel_id"`
				Name                 string      `json:"name"`
				ChannelType          string      `json:"channel_type"`
				GreetingEnabled      bool        `json:"greeting_enabled"`
				GreetingMessage      interface{} `json:"greeting_message"`
				WorkingHoursEnabled  bool        `json:"working_hours_enabled"`
				EnableEmailCollect   bool        `json:"enable_email_collect"`
				CsatSurveyEnabled    bool        `json:"csat_survey_enabled"`
				EnableAutoAssignment bool        `json:"enable_auto_assignment"`
				AutoAssignmentConfig struct {
					MaxAssignmentLimit interface{} `json:"max_assignment_limit"`
				} `json:"auto_assignment_config"`
				OutOfOfficeMessage interface{} `json:"out_of_office_message"`
				WorkingHours       []struct {
					DayOfWeek    int         `json:"day_of_week"`
					ClosedAllDay bool        `json:"closed_all_day"`
					OpenHour     interface{} `json:"open_hour"`
					OpenMinutes  interface{} `json:"open_minutes"`
					CloseHour    interface{} `json:"close_hour"`
					CloseMinutes interface{} `json:"close_minutes"`
					OpenAllDay   bool        `json:"open_all_day"`
				} `json:"working_hours"`
				Timezone                   string      `json:"timezone"`
				CallbackWebhookURL         string      `json:"callback_webhook_url"`
				AllowMessagesAfterResolved bool        `json:"allow_messages_after_resolved"`
				LockToSingleConversation   bool        `json:"lock_to_single_conversation"`
				SenderNameType             string      `json:"sender_name_type"`
				BusinessName               interface{} `json:"business_name"`
				WidgetColor                interface{} `json:"widget_color"`
				WebsiteURL                 interface{} `json:"website_url"`
				HmacMandatory              interface{} `json:"hmac_mandatory"`
				WelcomeTitle               interface{} `json:"welcome_title"`
				WelcomeTagline             interface{} `json:"welcome_tagline"`
				WebWidgetScript            interface{} `json:"web_widget_script"`
				WebsiteToken               interface{} `json:"website_token"`
				SelectedFeatureFlags       interface{} `json:"selected_feature_flags"`
				ReplyTime                  interface{} `json:"reply_time"`
				MessagingServiceSid        interface{} `json:"messaging_service_sid"`
				PhoneNumber                string      `json:"phone_number"`
				Provider                   string      `json:"provider"`
				MessageTemplates           []struct {
					ID         string `json:"id"`
					Name       string `json:"name"`
					Status     string `json:"status"`
					Category   string `json:"category"`
					Language   string `json:"language"`
					Components []struct {
						Type    string `json:"type"`
						Format  string `json:"format,omitempty"`
						Example struct {
							HeaderHandle []string `json:"header_handle"`
						} `json:"example,omitempty"`
						Text    string `json:"text,omitempty"`
						Buttons []struct {
							Text string `json:"text"`
							Type string `json:"type"`
						} `json:"buttons,omitempty"`
					} `json:"components"`
					SubCategory      string `json:"sub_category,omitempty"`
					ParameterFormat  string `json:"parameter_format"`
					PreviousCategory string `json:"previous_category,omitempty"`
				} `json:"message_templates"`
				ProviderConfig struct {
					APIKey             string `json:"api_key"`
					PhoneNumberID      string `json:"phone_number_id"`
					BusinessAccountID  string `json:"business_account_id"`
					WebhookVerifyToken string `json:"webhook_verify_token"`
				} `json:"provider_config"`
			} `json:"inbox"`
		} `json:"contact_inboxes"`
	} `json:"payload"`
}

type Brodcast struct {
	ID struct {
		Oid string `bson:"$oid"`
	} `bson:"_id"`
	RequestUserID         string `bson:"request_user_id"`
	BusinessUserID        string `bson:"business_user_id"`
	BusinessSentMessageID string `bson:"business_sent_message_id"`
	BusinessReply         bool   `bson:"business_reply"`
	V                     int    `bson:"__v"`
}

type ConvReplies struct {
	ConvReply []struct {
		ReplyUserType  string `bson:"reply_user_type"`
		ReplyContent   string `bson:"reply_content"`
		ReplyMessageID string `bson:"reply_message_id"`
	} `bson:"conv_reply"`
}

type Mongo struct {
	ConvStartMsg    string `bson:"conv_start_msg"`
	BusinessDetails []struct {
		BusinessName        string   `bson:"business_name"`
		Location            string   `bson:"location"`
		PhoneNumber         string   `bson:"phone_number"`
		BusinessExamples    []string `bson:"business_examples"`
		ItemRelatedExamples []string `bson:"item_related_examples"`
	} `bson:"business_details"`
	BusinessExamples    []string `bson:"business_examples"`
	ItemRelatedExamples []string `bson:"item_related_examples"`
}

type Business struct {
	BusinessName        string   `bson:"business_name"`
	Location            string   `bson:"location"`
	PhoneNumber         string   `bson:"phone_number"`
	BusinessExamples    []string `bson:"business_examples"`
	ItemRelatedExamples []string `bson:"item_related_examples"`
}

type UpdatedBusiness struct {
	BusinessName    string `bson:"business_name"`
	Location        string `bson:"location"`
	PhoneNumber     string `bson:"phone_number"`
	ID              int
	NumLeads        int
	NumResponse     int
	NumAccepts      int
	DetailedAddress struct {
		Ward        string `json:"ward"`
		Street      string `json:"street"`
		City        string `json:"city"`
		PostalCode  string `json:"postal_code"`
		State       string `json:"state"`
		CountryCode string `json:"country_code"`
	} `json:"google_maps_detailed_address"`
	Competitors   []interface{} `json:"google_maps_competitors"`
	Description   string        `json:"google_maps_description"`
	FeaturedImage string        `json:"google_maps_featured_image"`
	Images        []struct {
		About string `json:"about"`
		Link  string `json:"link"`
	} `json:"google_maps_images"`
	Sc                 map[string]interface{} `json:"sc"`
	B_ItemTagExamples  interface{}            `json:"b_item_tag_examples"`
	B_BusinessExamples interface{}            `json:"b_business_examples"`
}

type ScraperResponse []struct {
	ID           int         `json:"id"`
	Status       string      `json:"status"`
	TaskName     string      `json:"task_name"`
	ScraperName  string      `json:"scraper_name"`
	ScraperType  string      `json:"scraper_type"`
	IsAllTask    bool        `json:"is_all_task"`
	IsSync       bool        `json:"is_sync"`
	ParentTaskID interface{} `json:"parent_task_id"`
	Duration     float64     `json:"duration"`
	StartedAt    string      `json:"started_at"`
	FinishedAt   string      `json:"finished_at"`
	Data         struct {
		Queries                 []string    `json:"queries"`
		Country                 interface{} `json:"country"`
		BusinessType            string      `json:"business_type"`
		MaxCities               interface{} `json:"max_cities"`
		RandomizeCities         bool        `json:"randomize_cities"`
		APIKey                  string      `json:"api_key"`
		EnableReviewsExtraction bool        `json:"enable_reviews_extraction"`
		MaxReviews              int         `json:"max_reviews"`
		ReviewsSort             string      `json:"reviews_sort"`
		Lang                    string      `json:"lang"`
		MaxResults              int         `json:"max_results"`
		Coordinates             string      `json:"coordinates"`
		ZoomLevel               int         `json:"zoom_level"`
	} `json:"data"`
	Result []struct {
		PlaceID     string `json:"place_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		// IsSpendingOnAds bool          `json:"is_spending_on_ads"`
		Reviews     int           `json:"reviews"`
		Competitors []interface{} `json:"competitors"`
		Website     string        `json:"website"`
		Phone       string        `json:"phone"`
		// CanClaim        bool          `json:"can_claim"`
		// Owner           struct {
		// 	ID   string `json:"id"`
		// 	Name string `json:"name"`
		// 	Link string `json:"link"`
		// } `json:"owner"`
		FeaturedImage string `json:"featured_image"`
		// MainCategory   string      `json:"main_category"`
		// Categories     []string    `json:"categories"`
		// Rating         float64     `json:"rating"`
		// WorkdayTiming  interface{} `json:"workday_timing"`
		// Address        string      `json:"address"`
		// ReviewKeywords []struct {
		// 	Keyword string `json:"keyword"`
		// 	Count   int    `json:"count"`
		// } `json:"review_keywords"`
		// ReviewsPerRating struct {
		// 	Num1 int `json:"1"`
		// 	Num2 int `json:"2"`
		// 	Num3 int `json:"3"`
		// 	Num4 int `json:"4"`
		// 	Num5 int `json:"5"`
		// } `json:"reviews_per_rating"`
		// FeaturedQuestion interface{} `json:"featured_question"`
		// ReviewsLink      string      `json:"reviews_link"`
		Coordinates struct {
			Latitude  interface{} `json:"latitude"`
			Longitude interface{} `json:"longitude"`
		} `json:"coordinates"`
		// PlusCode        string `json:"plus_code"`
		DetailedAddress struct {
			Ward        string `json:"ward"`
			Street      string `json:"street"`
			City        string `json:"city"`
			PostalCode  string `json:"postal_code"`
			State       string `json:"state"`
			CountryCode string `json:"country_code"`
		} `json:"detailed_address"`
		Images []struct {
			About string `json:"about"`
			Link  string `json:"link"`
		} `json:"images"`
		// Query string `json:"query"`
	} `json:"result"`
	// ResultCount int    `json:"result_count"`
	// CreatedAt   string `json:"created_at"`
	// UpdatedAt   string `json:"updated_at"`
}

type Scraperbody struct {
	ScraperName string `json:"scraper_name"`
	Data        struct {
		Queries                 []string `json:"queries"`
		Country                 *string  `json:"country"`
		MaxCities               *int     `json:"max_cities"`
		RandomizeCities         bool     `json:"randomize_cities"`
		APIKey                  string   `json:"api_key"`
		EnableReviewsExtraction bool     `json:"enable_reviews_extraction"`
		MaxReviews              int      `json:"max_reviews"`
		ReviewsSort             string   `json:"reviews_sort"`
		Lang                    string   `json:"lang"`
		MaxResults              int      `json:"max_results"`
		Coordinates             string   `json:"coordinates"`
		ZoomLevel               int      `json:"zoom_level"`
	} `json:"data"`
}

type AutoGenerated []struct {
	ID           int         `json:"id"`
	Status       string      `json:"status"`
	TaskName     string      `json:"task_name"`
	ScraperName  string      `json:"scraper_name"`
	ScraperType  string      `json:"scraper_type"`
	IsAllTask    bool        `json:"is_all_task"`
	IsSync       bool        `json:"is_sync"`
	ParentTaskID interface{} `json:"parent_task_id"`
	Duration     float64     `json:"duration"`
	StartedAt    string      `json:"started_at"`
	FinishedAt   string      `json:"finished_at"`
	Data         struct {
		Queries                 []string    `json:"queries"`
		Country                 interface{} `json:"country"`
		BusinessType            string      `json:"business_type"`
		MaxCities               interface{} `json:"max_cities"`
		RandomizeCities         bool        `json:"randomize_cities"`
		APIKey                  string      `json:"api_key"`
		EnableReviewsExtraction bool        `json:"enable_reviews_extraction"`
		MaxReviews              int         `json:"max_reviews"`
		ReviewsSort             string      `json:"reviews_sort"`
		Lang                    string      `json:"lang"`
		MaxResults              int         `json:"max_results"`
		Coordinates             string      `json:"coordinates"`
		ZoomLevel               int         `json:"zoom_level"`
	} `json:"data"`
	Metadata struct {
	} `json:"metadata"`
	Result []struct {
		PlaceID         string        `json:"place_id"`
		Name            string        `json:"name"`
		Description     interface{}   `json:"description"`
		IsSpendingOnAds bool          `json:"is_spending_on_ads"`
		Reviews         int           `json:"reviews"`
		Competitors     []interface{} `json:"competitors"`
		Website         string        `json:"website"`
		Phone           string        `json:"phone"`
		CanClaim        bool          `json:"can_claim"`
		Owner           struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"owner"`
		FeaturedImage  string      `json:"featured_image"`
		MainCategory   string      `json:"main_category"`
		Categories     []string    `json:"categories"`
		Rating         float64     `json:"rating"`
		WorkdayTiming  interface{} `json:"workday_timing"`
		ClosedOn       string      `json:"closed_on"`
		Address        string      `json:"address"`
		ReviewKeywords []struct {
			Keyword string `json:"keyword"`
			Count   int    `json:"count"`
		} `json:"review_keywords"`
		Link             string      `json:"link"`
		Status           interface{} `json:"status"`
		PriceRange       interface{} `json:"price_range"`
		ReviewsPerRating struct {
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
			Num4 int `json:"4"`
			Num5 int `json:"5"`
		} `json:"reviews_per_rating"`
		FeaturedQuestion interface{} `json:"featured_question"`
		ReviewsLink      string      `json:"reviews_link"`
		Coordinates      struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`
		PlusCode        string `json:"plus_code"`
		DetailedAddress struct {
			Ward        string `json:"ward"`
			Street      string `json:"street"`
			City        string `json:"city"`
			PostalCode  string `json:"postal_code"`
			State       string `json:"state"`
			CountryCode string `json:"country_code"`
		} `json:"detailed_address"`
		TimeZone string        `json:"time_zone"`
		Cid      string        `json:"cid"`
		DataID   string        `json:"data_id"`
		About    []interface{} `json:"about"`
		Images   []struct {
			About string `json:"about"`
			Link  string `json:"link"`
		} `json:"images"`
		Hours            []interface{} `json:"hours"`
		MostPopularTimes string        `json:"most_popular_times"`
		PopularTimes     string        `json:"popular_times"`
		Menu             interface{}   `json:"menu"`
		Reservations     []interface{} `json:"reservations"`
		OrderOnlineLinks []interface{} `json:"order_online_links"`
		FeaturedReviews  []struct {
			ReviewID                        string        `json:"review_id"`
			Rating                          int           `json:"rating"`
			ReviewText                      string        `json:"review_text"`
			PublishedAt                     string        `json:"published_at"`
			PublishedAtDate                 string        `json:"published_at_date"`
			ResponseFromOwnerText           string        `json:"response_from_owner_text"`
			ResponseFromOwnerAgo            string        `json:"response_from_owner_ago"`
			ResponseFromOwnerDate           string        `json:"response_from_owner_date"`
			ReviewLikesCount                interface{}   `json:"review_likes_count"`
			TotalNumberOfReviewsByReviewer  interface{}   `json:"total_number_of_reviews_by_reviewer"`
			TotalNumberOfPhotosByReviewer   interface{}   `json:"total_number_of_photos_by_reviewer"`
			IsLocalGuide                    bool          `json:"is_local_guide"`
			ReviewTranslatedText            interface{}   `json:"review_translated_text"`
			ResponseFromOwnerTranslatedText interface{}   `json:"response_from_owner_translated_text"`
			ReviewPhotos                    []interface{} `json:"review_photos"`
		} `json:"featured_reviews"`
		DetailedReviews []struct {
			ReviewID                        string      `json:"review_id"`
			Rating                          int         `json:"rating"`
			ReviewText                      string      `json:"review_text"`
			PublishedAt                     string      `json:"published_at"`
			PublishedAtDate                 string      `json:"published_at_date"`
			ResponseFromOwnerText           string      `json:"response_from_owner_text"`
			ResponseFromOwnerAgo            string      `json:"response_from_owner_ago"`
			ResponseFromOwnerDate           string      `json:"response_from_owner_date"`
			ReviewLikesCount                int         `json:"review_likes_count"`
			TotalNumberOfReviewsByReviewer  int         `json:"total_number_of_reviews_by_reviewer"`
			TotalNumberOfPhotosByReviewer   interface{} `json:"total_number_of_photos_by_reviewer"`
			IsLocalGuide                    bool        `json:"is_local_guide"`
			ReviewTranslatedText            interface{} `json:"review_translated_text"`
			ResponseFromOwnerTranslatedText interface{} `json:"response_from_owner_translated_text"`
		} `json:"detailed_reviews"`
		Query string `json:"query"`
	} `json:"result"`
	ResultCount int    `json:"result_count"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
