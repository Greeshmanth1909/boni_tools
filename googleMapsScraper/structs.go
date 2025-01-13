package main

type Response []struct {
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
		PlaceID         string `json:"place_id"`
		Name            string `json:"name"`
		Description     string `json:"description"`
		IsSpendingOnAds bool   `json:"is_spending_on_ads"`
		Reviews         int    `json:"reviews"`
		Competitors     []struct {
			Name         string  `json:"name"`
			Link         string  `json:"link"`
			Reviews      int     `json:"reviews"`
			Rating       float64 `json:"rating"`
			MainCategory string  `json:"main_category"`
		} `json:"competitors"`
		Website  string `json:"website"`
		Phone    string `json:"phone"`
		CanClaim bool   `json:"can_claim"`
		Owner    struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"owner"`
		FeaturedImage  string   `json:"featured_image"`
		MainCategory   string   `json:"main_category"`
		Categories     []string `json:"categories"`
		Rating         float64  `json:"rating"`
		WorkdayTiming  string   `json:"workday_timing"`
		ClosedOn       string   `json:"closed_on"`
		Address        string   `json:"address"`
		ReviewKeywords []struct {
			Keyword string `json:"keyword"`
			Count   int    `json:"count"`
		} `json:"review_keywords"`
		Link             string      `json:"link"`
		Status           string      `json:"status"`
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
		TimeZone string `json:"time_zone"`
		Cid      string `json:"cid"`
		DataID   string `json:"data_id"`
		About    []struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Options []struct {
				Name    string `json:"name"`
				Enabled bool   `json:"enabled"`
			} `json:"options"`
		} `json:"about"`
		Images []struct {
			About string `json:"about"`
			Link  string `json:"link"`
		} `json:"images"`
		Hours []struct {
			Day   string   `json:"day"`
			Times []string `json:"times"`
		} `json:"hours"`
		MostPopularTimes []struct {
			HourOfDay         int     `json:"hour_of_day"`
			AveragePopularity float64 `json:"average_popularity"`
			TimeLabel         string  `json:"time_label"`
		} `json:"most_popular_times"`
		PopularTimes struct {
			Monday []struct {
				HourOfDay             int    `json:"hour_of_day"`
				TimeLabel             string `json:"time_label"`
				PopularityPercentage  int    `json:"popularity_percentage"`
				PopularityDescription string `json:"popularity_description"`
			} `json:"Monday"`
			Tuesday []struct {
				HourOfDay             int    `json:"hour_of_day"`
				TimeLabel             string `json:"time_label"`
				PopularityPercentage  int    `json:"popularity_percentage"`
				PopularityDescription string `json:"popularity_description"`
			} `json:"Tuesday"`
			Wednesday []struct {
				HourOfDay             int    `json:"hour_of_day"`
				TimeLabel             string `json:"time_label"`
				PopularityPercentage  int    `json:"popularity_percentage"`
				PopularityDescription string `json:"popularity_description"`
			} `json:"Wednesday"`
			Thursday []struct {
				HourOfDay             int    `json:"hour_of_day"`
				TimeLabel             string `json:"time_label"`
				PopularityPercentage  int    `json:"popularity_percentage"`
				PopularityDescription string `json:"popularity_description"`
			} `json:"Thursday"`
			Friday []struct {
				HourOfDay             int    `json:"hour_of_day"`
				TimeLabel             string `json:"time_label"`
				PopularityPercentage  int    `json:"popularity_percentage"`
				PopularityDescription string `json:"popularity_description"`
			} `json:"Friday"`
			Saturday []struct {
				HourOfDay             int    `json:"hour_of_day"`
				TimeLabel             string `json:"time_label"`
				PopularityPercentage  int    `json:"popularity_percentage"`
				PopularityDescription string `json:"popularity_description"`
			} `json:"Saturday"`
			Sunday []struct {
				HourOfDay             int    `json:"hour_of_day"`
				TimeLabel             string `json:"time_label"`
				PopularityPercentage  int    `json:"popularity_percentage"`
				PopularityDescription string `json:"popularity_description"`
			} `json:"Sunday"`
		} `json:"popular_times"`
		Menu             interface{}   `json:"menu"`
		Reservations     []interface{} `json:"reservations"`
		OrderOnlineLinks []struct {
			Link   string `json:"link"`
			Source string `json:"source"`
		} `json:"order_online_links"`
		FeaturedReviews []struct {
			ReviewID                        string        `json:"review_id"`
			Rating                          int           `json:"rating"`
			ReviewText                      string        `json:"review_text"`
			PublishedAt                     string        `json:"published_at"`
			PublishedAtDate                 string        `json:"published_at_date"`
			ResponseFromOwnerText           string        `json:"response_from_owner_text"`
			ResponseFromOwnerAgo            string        `json:"response_from_owner_ago"`
			ResponseFromOwnerDate           string        `json:"response_from_owner_date"`
			ReviewLikesCount                int           `json:"review_likes_count"`
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

type Result struct {
    PlaceID         string `json:"place_id"`
    Name            string `json:"name"`
    Description     string `json:"description"`
    IsSpendingOnAds bool   `json:"is_spending_on_ads"`
    Reviews         int    `json:"reviews"`
    Competitors     []struct {
        Name         string  `json:"name"`
        Link         string  `json:"link"`
        Reviews      int     `json:"reviews"`
        Rating       float64 `json:"rating"`
        MainCategory string  `json:"main_category"`
    } `json:"competitors"`
    Website  string `json:"website"`
    Phone    string `json:"phone"`
    CanClaim bool   `json:"can_claim"`
    Owner    struct {
        ID   string `json:"id"`
        Name string `json:"name"`
        Link string `json:"link"`
    } `json:"owner"`
    FeaturedImage  string   `json:"featured_image"`
    MainCategory   string   `json:"main_category"`
    Categories     []string `json:"categories"`
    Rating         float64  `json:"rating"`
    WorkdayTiming  string   `json:"workday_timing"`
    ClosedOn       string   `json:"closed_on"`
    Address        string   `json:"address"`
    ReviewKeywords []struct {
        Keyword string `json:"keyword"`
        Count   int    `json:"count"`
    } `json:"review_keywords"`
    Link             string      `json:"link"`
    Status           string      `json:"status"`
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
    TimeZone string `json:"time_zone"`
    Cid      string `json:"cid"`
    DataID   string `json:"data_id"`
    About    []struct {
        ID      string `json:"id"`
        Name    string `json:"name"`
        Options []struct {
            Name    string `json:"name"`
            Enabled bool   `json:"enabled"`
        } `json:"options"`
    } `json:"about"`
    Images []struct {
        About string `json:"about"`
        Link  string `json:"link"`
    } `json:"images"`
    Hours []struct {
        Day   string   `json:"day"`
        Times []string `json:"times"`
    } `json:"hours"`
    MostPopularTimes []struct {
        HourOfDay         int     `json:"hour_of_day"`
        AveragePopularity float64 `json:"average_popularity"`
        TimeLabel         string  `json:"time_label"`
    } `json:"most_popular_times"`
    PopularTimes struct {
        Monday []struct {
            HourOfDay             int    `json:"hour_of_day"`
            TimeLabel             string `json:"time_label"`
            PopularityPercentage  int    `json:"popularity_percentage"`
            PopularityDescription string `json:"popularity_description"`
        } `json:"Monday"`
        Tuesday []struct {
            HourOfDay             int    `json:"hour_of_day"`
            TimeLabel             string `json:"time_label"`
            PopularityPercentage  int    `json:"popularity_percentage"`
            PopularityDescription string `json:"popularity_description"`
        } `json:"Tuesday"`
        Wednesday []struct {
            HourOfDay             int    `json:"hour_of_day"`
            TimeLabel             string `json:"time_label"`
            PopularityPercentage  int    `json:"popularity_percentage"`
            PopularityDescription string `json:"popularity_description"`
        } `json:"Wednesday"`
        Thursday []struct {
            HourOfDay             int    `json:"hour_of_day"`
            TimeLabel             string `json:"time_label"`
            PopularityPercentage  int    `json:"popularity_percentage"`
            PopularityDescription string `json:"popularity_description"`
        } `json:"Thursday"`
        Friday []struct {
            HourOfDay             int    `json:"hour_of_day"`
            TimeLabel             string `json:"time_label"`
            PopularityPercentage  int    `json:"popularity_percentage"`
            PopularityDescription string `json:"popularity_description"`
        } `json:"Friday"`
        Saturday []struct {
            HourOfDay             int    `json:"hour_of_day"`
            TimeLabel             string `json:"time_label"`
            PopularityPercentage  int    `json:"popularity_percentage"`
            PopularityDescription string `json:"popularity_description"`
        } `json:"Saturday"`
        Sunday []struct {
            HourOfDay             int    `json:"hour_of_day"`
            TimeLabel             string `json:"time_label"`
            PopularityPercentage  int    `json:"popularity_percentage"`
            PopularityDescription string `json:"popularity_description"`
        } `json:"Sunday"`
    } `json:"popular_times"`
    Menu             interface{}   `json:"menu"`
    Reservations     []interface{} `json:"reservations"`
    OrderOnlineLinks []struct {
        Link   string `json:"link"`
        Source string `json:"source"`
    } `json:"order_online_links"`
    FeaturedReviews []struct {
        ReviewID                        string        `json:"review_id"`
        Rating                          int           `json:"rating"`
        ReviewText                      string        `json:"review_text"`
        PublishedAt                     string        `json:"published_at"`
        PublishedAtDate                 string        `json:"published_at_date"`
        ResponseFromOwnerText           string        `json:"response_from_owner_text"`
        ResponseFromOwnerAgo            string        `json:"response_from_owner_ago"`
        ResponseFromOwnerDate           string        `json:"response_from_owner_date"`
        ReviewLikesCount                int           `json:"review_likes_count"`
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
}
