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
		BusinessName string `bson:"business_name"`
		Location     string `bson:"location"`
		PhoneNumber  string `bson:"phone_number"`
	} `bson:"business_details"`
}

type Business struct {
		BusinessName string `bson:"business_name"`
		Location     string `bson:"location"`
		PhoneNumber  string `bson:"phone_number"`
}

type UpdatedBusiness struct {
		BusinessName string `bson:"business_name"`
		Location     string `bson:"location"`
		PhoneNumber  string `bson:"phone_number"`
        ID           int
        NumLeads     int
        NumResponse  int
        NumAccepts   int
}
