package model

// module: notice, action: noticeMarquee
type NoticeMarqueeResult struct {
	ItemCount   int           `json:"item_count"`
	MarqueeList []interface{} `json:"marquee_list"`
}

// module: user, action: getNavi
type User struct {
	UserID           int   `json:"user_id"`
	UnitOwningUserID int64 `json:"unit_owning_user_id"`
}

type UserNaviResult struct {
	User User `json:"user"`
}

// module: navigation, action: specialCutin
type SpecialCutinResult struct {
	SpecialCutinList []interface{} `json:"special_cutin_list"`
}

// module: award, action: awardInfo
type AwardInfo struct {
	AwardID    int    `json:"award_id"`
	IsSet      bool   `json:"is_set"`
	InsertDate string `json:"insert_date"`
}

type AwardInfoResult struct {
	AwardInfo []AwardInfo `json:"award_info"`
}

// module: background, action: backgroundInfo
type BackgroundInfo struct {
	BackgroundID int    `json:"background_id"`
	IsSet        bool   `json:"is_set"`
	InsertDate   string `json:"insert_date"`
}

type BackgroundInfoResult struct {
	BackgroundInfo []BackgroundInfo `json:"background_info"`
}

// module: exchange, action: owningPoint
type ExchangePointList struct {
	Rarity        int `json:"rarity"`
	ExchangePoint int `json:"exchange_point"`
}

type OwningPointResult struct {
	ExchangePointList []ExchangePointList `json:"exchange_point_list"`
}

// module: livese, action: liveseInfo
type LiveSeInfoResult struct {
	LiveSeList []int `json:"live_se_list"`
}

// module: liveicon, action: liveiconInfo
type LiveIconInfoResult struct {
	LiveNotesIconList []int `json:"live_notes_icon_list"`
}

// module: item, action: list

// module: marathon, action: marathonInfo

// module: challenge, action: challengeInfo
