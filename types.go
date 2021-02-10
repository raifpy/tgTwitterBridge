package main

type twitterResponseStruct struct {
	CreatedAt        string `json:"created_at"`
	ID               int64  `json:"id"`
	IDStr            string `json:"id_str"`
	FullText         string `json:"full_text"`
	Truncated        bool   `json:"truncated"`
	DisplayTextRange []int  `json:"display_text_range"`
	Entities         struct {
		Hashtags     []interface{} `json:"hashtags"`
		Symbols      []interface{} `json:"symbols"`
		UserMentions []interface{} `json:"user_mentions"`
		Urls         []interface{} `json:"urls"`
	} `json:"entities"`
	Source               string      `json:"source"`
	InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
	InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
	User                 struct {
		ID          int64       `json:"id"`
		IDStr       string      `json:"id_str"`
		Name        string      `json:"name"`
		ScreenName  string      `json:"screen_name"`
		Location    string      `json:"location"`
		Description string      `json:"description"`
		URL         interface{} `json:"url"`
		Entities    struct {
			Description struct {
				Urls []interface{} `json:"urls"`
			} `json:"description"`
		} `json:"entities"`
		Protected                               bool        `json:"protected"`
		FollowersCount                          int         `json:"followers_count"`
		FastFollowersCount                      int         `json:"fast_followers_count"`
		NormalFollowersCount                    int         `json:"normal_followers_count"`
		FriendsCount                            int         `json:"friends_count"`
		ListedCount                             int         `json:"listed_count"`
		CreatedAt                               string      `json:"created_at"`
		FavouritesCount                         int         `json:"favourites_count"`
		UtcOffset                               interface{} `json:"utc_offset"`
		TimeZone                                interface{} `json:"time_zone"`
		GeoEnabled                              bool        `json:"geo_enabled"`
		Verified                                bool        `json:"verified"`
		StatusesCount                           int         `json:"statuses_count"`
		MediaCount                              int         `json:"media_count"`
		Lang                                    interface{} `json:"lang"`
		ContributorsEnabled                     bool        `json:"contributors_enabled"`
		IsTranslator                            bool        `json:"is_translator"`
		IsTranslationEnabled                    bool        `json:"is_translation_enabled"`
		ProfileBackgroundColor                  string      `json:"profile_background_color"`
		ProfileBackgroundImageURL               string      `json:"profile_background_image_url"`
		ProfileBackgroundImageURLHTTPS          string      `json:"profile_background_image_url_https"`
		ProfileBackgroundTile                   bool        `json:"profile_background_tile"`
		ProfileImageURL                         string      `json:"profile_image_url"`
		ProfileImageURLHTTPS                    string      `json:"profile_image_url_https"`
		ProfileBannerURL                        string      `json:"profile_banner_url"`
		ProfileImageExtensionsMediaAvailability interface{} `json:"profile_image_extensions_media_availability"`
		ProfileImageExtensionsAltText           interface{} `json:"profile_image_extensions_alt_text"`
		ProfileImageExtensionsMediaColor        struct {
			Palette []struct {
				Rgb struct {
					Red   int `json:"red"`
					Green int `json:"green"`
					Blue  int `json:"blue"`
				} `json:"rgb"`
				Percentage float64 `json:"percentage"`
			} `json:"palette"`
		} `json:"profile_image_extensions_media_color"`
		ProfileBannerExtensionsMediaAvailability interface{} `json:"profile_banner_extensions_media_availability"`
		ProfileBannerExtensionsAltText           interface{} `json:"profile_banner_extensions_alt_text"`
		ProfileBannerExtensionsMediaColor        struct {
			Palette []struct {
				Rgb struct {
					Red   int `json:"red"`
					Green int `json:"green"`
					Blue  int `json:"blue"`
				} `json:"rgb"`
				Percentage float64 `json:"percentage"`
			} `json:"palette"`
		} `json:"profile_banner_extensions_media_color"`
		ProfileLinkColor               string        `json:"profile_link_color"`
		ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string        `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool          `json:"profile_use_background_image"`
		HasExtendedProfile             bool          `json:"has_extended_profile"`
		DefaultProfile                 bool          `json:"default_profile"`
		DefaultProfileImage            bool          `json:"default_profile_image"`
		PinnedTweetIds                 []interface{} `json:"pinned_tweet_ids"`
		PinnedTweetIdsStr              []interface{} `json:"pinned_tweet_ids_str"`
		HasCustomTimelines             bool          `json:"has_custom_timelines"`
		CanDm                          bool          `json:"can_dm"`
		CanMediaTag                    bool          `json:"can_media_tag"`
		Following                      bool          `json:"following"`
		FollowRequestSent              bool          `json:"follow_request_sent"`
		Notifications                  bool          `json:"notifications"`
		Muting                         bool          `json:"muting"`
		Blocking                       bool          `json:"blocking"`
		BlockedBy                      bool          `json:"blocked_by"`
		WantRetweets                   bool          `json:"want_retweets"`
		AdvertiserAccountType          string        `json:"advertiser_account_type"`
		AdvertiserAccountServiceLevels []interface{} `json:"advertiser_account_service_levels"`
		ProfileInterstitialType        string        `json:"profile_interstitial_type"`
		BusinessProfileState           string        `json:"business_profile_state"`
		TranslatorType                 string        `json:"translator_type"`
		FollowedBy                     bool          `json:"followed_by"`
		RequireSomeConsent             bool          `json:"require_some_consent"`
	} `json:"user"`
	Geo                  interface{} `json:"geo"`
	Coordinates          interface{} `json:"coordinates"`
	Place                interface{} `json:"place"`
	Contributors         interface{} `json:"contributors"`
	IsQuoteStatus        bool        `json:"is_quote_status"`
	RetweetCount         int         `json:"retweet_count"`
	FavoriteCount        int         `json:"favorite_count"`
	ReplyCount           int         `json:"reply_count"`
	QuoteCount           int         `json:"quote_count"`
	Favorited            bool        `json:"favorited"`
	Retweeted            bool        `json:"retweeted"`
	Lang                 string      `json:"lang"`
	SupplementalLanguage interface{} `json:"supplemental_language"`
}

type mediaStruct struct {
	MediaID          int64  `json:"media_id"`
	MediaIDString    string `json:"media_id_string"`
	ExpiresAfterSecs int    `json:"expires_after_secs"`
	MediaKey         string `json:"media_key"`
}

type twitterUploadFinalizeResponse struct {
	MediaID          int64  `json:"media_id"`
	MediaIDString    string `json:"media_id_string"`
	MediaKey         string `json:"media_key"`
	Size             int    `json:"size"`
	ExpiresAfterSecs int    `json:"expires_after_secs"`
	ProcessingInfo   struct {
		State          string `json:"state"`
		CheckAfterSecs int    `json:"check_after_secs"`
	} `json:"processing_info"`
}

type telegramConfigStruct struct {
	Token     string `json:"Token"`
	ChannelID int64  `json:"ChannelID"`
	LogID     int64  `json:"LogID"`
}
