package main

type ResponseStruct struct {
	From      int `json:"from"`
	To        int `json:"to"`
	Total     int `json:"total"`
	AuditData struct {
		ProcessTime string `json:"processTime"`
		Timestamp   string `json:"timestamp"`
		RequestHost string `json:"requestHost"`
		ServerID    string `json:"serverId"`
		Environment string `json:"environment"`
		Release     string `json:"release"`
	} `json:"auditData"`
	Hotels []struct {
		Code int `json:"code"`
		Name struct {
			Content string `json:"content"`
		} `json:"name"`
		Description struct {
			Content string `json:"content"`
		} `json:"description"`
		CountryCode     string `json:"countryCode"`
		StateCode       string `json:"stateCode"`
		DestinationCode string `json:"destinationCode"`
		ZoneCode        int    `json:"zoneCode"`
		Coordinates     struct {
			Longitude float64 `json:"longitude"`
			Latitude  float64 `json:"latitude"`
		} `json:"coordinates"`
		CategoryCode          string   `json:"categoryCode"`
		CategoryGroupCode     string   `json:"categoryGroupCode"`
		ChainCode             string   `json:"chainCode,omitempty"`
		AccommodationTypeCode string   `json:"accommodationTypeCode"`
		BoardCodes            []string `json:"boardCodes,omitempty"`
		SegmentCodes          []int    `json:"segmentCodes"`
		Address               struct {
			Content string `json:"content"`
			Street  string `json:"street"`
			Number  string `json:"number"`
		} `json:"address"`
		PostalCode string `json:"postalCode"`
		City       struct {
			Content string `json:"content"`
		} `json:"city"`
		Email     string `json:"email,omitempty"`
		License   string `json:"license,omitempty"`
		GiataCode int    `json:"giataCode"`
		Phones    []struct {
			PhoneNumber string `json:"phoneNumber"`
			PhoneType   string `json:"phoneType"`
		} `json:"phones"`
		Rooms []struct {
			RoomCode           string `json:"roomCode"`
			IsParentRoom       bool   `json:"isParentRoom"`
			MinPax             int    `json:"minPax"`
			MaxPax             int    `json:"maxPax"`
			MaxAdults          int    `json:"maxAdults"`
			MaxChildren        int    `json:"maxChildren"`
			MinAdults          int    `json:"minAdults"`
			RoomType           string `json:"roomType"`
			CharacteristicCode string `json:"characteristicCode"`
			RoomFacilities     []struct {
				FacilityCode      int  `json:"facilityCode"`
				FacilityGroupCode int  `json:"facilityGroupCode"`
				IndLogic          bool `json:"indLogic,omitempty"`
				Number            int  `json:"number"`
				Voucher           bool `json:"voucher"`
				IndYesOrNo        bool `json:"indYesOrNo,omitempty"`
			} `json:"roomFacilities,omitempty"`
			RoomStays []struct {
				StayType           string `json:"stayType"`
				Order              string `json:"order"`
				Description        string `json:"description"`
				RoomStayFacilities []struct {
					FacilityCode      int `json:"facilityCode"`
					FacilityGroupCode int `json:"facilityGroupCode"`
					Number            int `json:"number"`
				} `json:"roomStayFacilities"`
			} `json:"roomStays,omitempty"`
			PMSRoomCode string `json:"PMSRoomCode,omitempty"`
		} `json:"rooms"`
		Facilities []struct {
			FacilityCode      int     `json:"facilityCode"`
			FacilityGroupCode int     `json:"facilityGroupCode"`
			Order             int     `json:"order"`
			IndYesOrNo        bool    `json:"indYesOrNo,omitempty"`
			Number            int     `json:"number,omitempty"`
			Voucher           bool    `json:"voucher"`
			IndLogic          bool    `json:"indLogic,omitempty"`
			IndFee            bool    `json:"indFee,omitempty"`
			Distance          int     `json:"distance,omitempty"`
			Amount            float64 `json:"amount,omitempty"`
			Currency          string  `json:"currency,omitempty"`
			ApplicationType   string  `json:"applicationType,omitempty"`
			TimeFrom          string  `json:"timeFrom,omitempty"`
			TimeTo            string  `json:"timeTo,omitempty"`
			DateTo            string  `json:"dateTo,omitempty"`
		} `json:"facilities"`
		Terminals []struct {
			TerminalCode string `json:"terminalCode"`
			Distance     int    `json:"distance"`
		} `json:"terminals,omitempty"`
		InterestPoints []struct {
			FacilityCode      int    `json:"facilityCode"`
			FacilityGroupCode int    `json:"facilityGroupCode"`
			Order             int    `json:"order"`
			PoiName           string `json:"poiName"`
			Distance          string `json:"distance"`
		} `json:"interestPoints,omitempty"`
		Images []struct {
			ImageTypeCode      string   `json:"imageTypeCode"`
			Path               string   `json:"path"`
			Order              int      `json:"order"`
			VisualOrder        int      `json:"visualOrder"`
			RoomCode           string   `json:"roomCode,omitempty"`
			RoomType           string   `json:"roomType,omitempty"`
			CharacteristicCode string   `json:"characteristicCode,omitempty"`
			ImagePaths         []string `json:"imagePaths"`
		} `json:"images"`
		Wildcards []struct {
			RoomType             string `json:"roomType"`
			RoomCode             string `json:"roomCode"`
			CharacteristicCode   string `json:"characteristicCode"`
			HotelRoomDescription struct {
				Content string `json:"content"`
			} `json:"hotelRoomDescription"`
		} `json:"wildcards,omitempty"`
		Web        string `json:"web,omitempty"`
		LastUpdate string `json:"lastUpdate"`
		S2C        string `json:"S2C,omitempty"`
		Ranking    int    `json:"ranking"`
		Issues     []struct {
			IssueCode   string `json:"issueCode"`
			IssueType   string `json:"issueType"`
			DateFrom    string `json:"dateFrom"`
			DateTo      string `json:"dateTo"`
			Order       int    `json:"order"`
			Alternative bool   `json:"alternative"`
		} `json:"issues,omitempty"`
	} `json:"hotels"`
}

type HotelsDB struct {
	Hotels []struct {
		Code int `json:"code"`
		Name struct {
			Content string `json:"content"`
		} `json:"name"`
		Description struct {
			Content string `json:"content"`
		} `json:"description"`
		CountryCode     string `json:"countryCode"`
		StateCode       string `json:"stateCode"`
		DestinationCode string `json:"destinationCode"`
		ZoneCode        int    `json:"zoneCode"`
		Coordinates     struct {
			Longitude float64 `json:"longitude"`
			Latitude  float64 `json:"latitude"`
		} `json:"coordinates"`
		CategoryCode          string   `json:"categoryCode"`
		CategoryGroupCode     string   `json:"categoryGroupCode"`
		ChainCode             string   `json:"chainCode,omitempty"`
		AccommodationTypeCode string   `json:"accommodationTypeCode"`
		BoardCodes            []string `json:"boardCodes,omitempty"`
		SegmentCodes          []int    `json:"segmentCodes"`
		Address               struct {
			Content string `json:"content"`
			Street  string `json:"street"`
			Number  string `json:"number"`
		} `json:"address"`
		PostalCode string `json:"postalCode"`
		City       struct {
			Content string `json:"content"`
		} `json:"city"`
		Email     string `json:"email,omitempty"`
		License   string `json:"license,omitempty"`
		GiataCode int    `json:"giataCode"`
		Phones    []struct {
			PhoneNumber string `json:"phoneNumber"`
			PhoneType   string `json:"phoneType"`
		} `json:"phones"`
		Rooms []struct {
			RoomCode           string `json:"roomCode"`
			IsParentRoom       bool   `json:"isParentRoom"`
			MinPax             int    `json:"minPax"`
			MaxPax             int    `json:"maxPax"`
			MaxAdults          int    `json:"maxAdults"`
			MaxChildren        int    `json:"maxChildren"`
			MinAdults          int    `json:"minAdults"`
			RoomType           string `json:"roomType"`
			CharacteristicCode string `json:"characteristicCode"`
			RoomFacilities     []struct {
				FacilityCode      int  `json:"facilityCode"`
				FacilityGroupCode int  `json:"facilityGroupCode"`
				IndLogic          bool `json:"indLogic,omitempty"`
				Number            int  `json:"number"`
				Voucher           bool `json:"voucher"`
				IndYesOrNo        bool `json:"indYesOrNo,omitempty"`
			} `json:"roomFacilities,omitempty"`
			RoomStays []struct {
				StayType           string `json:"stayType"`
				Order              string `json:"order"`
				Description        string `json:"description"`
				RoomStayFacilities []struct {
					FacilityCode      int `json:"facilityCode"`
					FacilityGroupCode int `json:"facilityGroupCode"`
					Number            int `json:"number"`
				} `json:"roomStayFacilities"`
			} `json:"roomStays,omitempty"`
			PMSRoomCode string `json:"PMSRoomCode,omitempty"`
		} `json:"rooms"`
		Facilities []struct {
			FacilityCode      int     `json:"facilityCode"`
			FacilityGroupCode int     `json:"facilityGroupCode"`
			Order             int     `json:"order"`
			IndYesOrNo        bool    `json:"indYesOrNo,omitempty"`
			Number            int     `json:"number,omitempty"`
			Voucher           bool    `json:"voucher"`
			IndLogic          bool    `json:"indLogic,omitempty"`
			IndFee            bool    `json:"indFee,omitempty"`
			Distance          int     `json:"distance,omitempty"`
			Amount            float64 `json:"amount,omitempty"`
			Currency          string  `json:"currency,omitempty"`
			ApplicationType   string  `json:"applicationType,omitempty"`
			TimeFrom          string  `json:"timeFrom,omitempty"`
			TimeTo            string  `json:"timeTo,omitempty"`
			DateTo            string  `json:"dateTo,omitempty"`
		} `json:"facilities"`
		Terminals []struct {
			TerminalCode string `json:"terminalCode"`
			Distance     int    `json:"distance"`
		} `json:"terminals,omitempty"`
		InterestPoints []struct {
			FacilityCode      int    `json:"facilityCode"`
			FacilityGroupCode int    `json:"facilityGroupCode"`
			Order             int    `json:"order"`
			PoiName           string `json:"poiName"`
			Distance          string `json:"distance"`
		} `json:"interestPoints,omitempty"`
		Images []struct {
			ImageTypeCode      string   `json:"imageTypeCode"`
			Path               string   `json:"path"`
			Order              int      `json:"order"`
			VisualOrder        int      `json:"visualOrder"`
			RoomCode           string   `json:"roomCode,omitempty"`
			RoomType           string   `json:"roomType,omitempty"`
			CharacteristicCode string   `json:"characteristicCode,omitempty"`
			ImagePaths         []string `json:"imagePaths"`
		} `json:"images"`
		Wildcards []struct {
			RoomType             string `json:"roomType"`
			RoomCode             string `json:"roomCode"`
			CharacteristicCode   string `json:"characteristicCode"`
			HotelRoomDescription struct {
				Content string `json:"content"`
			} `json:"hotelRoomDescription"`
		} `json:"wildcards,omitempty"`
		Web        string `json:"web,omitempty"`
		LastUpdate string `json:"lastUpdate"`
		S2C        string `json:"S2C,omitempty"`
		Ranking    int    `json:"ranking"`
		Issues     []struct {
			IssueCode   string `json:"issueCode"`
			IssueType   string `json:"issueType"`
			DateFrom    string `json:"dateFrom"`
			DateTo      string `json:"dateTo"`
			Order       int    `json:"order"`
			Alternative bool   `json:"alternative"`
		} `json:"issues,omitempty"`
	} `json:"hotels"`
}
