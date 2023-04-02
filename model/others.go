package model

// module: notice, action: noticeMarquee
type NoticeMarqueeResult struct {
	ItemCount   int           `json:"item_count"`
	MarqueeList []interface{} `json:"marquee_list"`
}

type NoticeMarqueeResp struct {
	Result     NoticeMarqueeResult `json:"result"`
	Status     int                 `json:"status"`
	CommandNum bool                `json:"commandNum"`
	TimeStamp  int64               `json:"timeStamp"`
}

// module: user, action: getNavi
type User struct {
	UserID           int `json:"user_id"`
	UnitOwningUserID int `json:"unit_owning_user_id"`
}

type UserNaviResult struct {
	User User `json:"user"`
}

type UserNaviResp struct {
	Result     UserNaviResult `json:"result"`
	Status     int            `json:"status"`
	CommandNum bool           `json:"commandNum"`
	TimeStamp  int64          `json:"timeStamp"`
}

// module: navigation, action: specialCutin
type SpecialCutinResult struct {
	SpecialCutinList []interface{} `json:"special_cutin_list"`
}

type SpecialCutinResp struct {
	Result     SpecialCutinResult `json:"result"`
	Status     int                `json:"status"`
	CommandNum bool               `json:"commandNum"`
	TimeStamp  int64              `json:"timeStamp"`
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

type AwardInfoResp struct {
	Result     AwardInfoResult `json:"result"`
	Status     int             `json:"status"`
	CommandNum bool            `json:"commandNum"`
	TimeStamp  int64           `json:"timeStamp"`
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

type BackgroundInfoResp struct {
	Result     BackgroundInfoResult `json:"result"`
	Status     int                  `json:"status"`
	CommandNum bool                 `json:"commandNum"`
	TimeStamp  int64                `json:"timeStamp"`
}

// module: exchange, action: owningPoint
type ExchangePointList struct {
	Rarity        int `json:"rarity"`
	ExchangePoint int `json:"exchange_point"`
}

type ExchangePointResult struct {
	ExchangePointList []ExchangePointList `json:"exchange_point_list"`
}

type ExchangePointResp struct {
	Result     ExchangePointResult `json:"result"`
	Status     int                 `json:"status"`
	CommandNum bool                `json:"commandNum"`
	TimeStamp  int64               `json:"timeStamp"`
}

// module: livese, action: liveseInfo
type LiveSeInfoResult struct {
	LiveSeList []int `json:"live_se_list"`
}

type LiveSeInfoResp struct {
	Result     LiveSeInfoResult `json:"result"`
	Status     int              `json:"status"`
	CommandNum bool             `json:"commandNum"`
	TimeStamp  int64            `json:"timeStamp"`
}

// module: liveicon, action: liveiconInfo
type LiveIconInfoResult struct {
	LiveNotesIconList []int `json:"live_notes_icon_list"`
}

type LiveIconInfoResp struct {
	Result     LiveIconInfoResult `json:"result"`
	Status     int                `json:"status"`
	CommandNum bool               `json:"commandNum"`
	TimeStamp  int64              `json:"timeStamp"`
}

// module: item, action: list
type GeneralItemList struct {
	ItemID          int  `json:"item_id"`
	Amount          int  `json:"amount"`
	UseButtonFlag   bool `json:"use_button_flag"`
	GeneralItemType int  `json:"general_item_type"`
}

type BuffItemList struct {
	ItemID   int `json:"item_id"`
	Amount   int `json:"amount"`
	BuffType int `json:"buff_type"`
}

type ReinforceItemList struct {
	ItemID        int `json:"item_id"`
	ReinforceType int `json:"reinforce_type"`
	AdditionValue int `json:"addition_value"`
	Amount        int `json:"amount"`
	EventID       int `json:"event_id"`
}

type ItemResult struct {
	GeneralItemList   []GeneralItemList   `json:"general_item_list"`
	BuffItemList      []BuffItemList      `json:"buff_item_list"`
	ReinforceItemList []ReinforceItemList `json:"reinforce_item_list"`
	ReinforceInfo     interface{}         `json:"reinforce_info"`
}

type ItemResp struct {
	Result     ItemResult `json:"result"`
	Status     int        `json:"status"`
	CommandNum bool       `json:"commandNum"`
	TimeStamp  int64      `json:"timeStamp"`
}

// module: marathon, action: marathonInfo
type MarathonInfoResp struct {
	Result     []interface{} `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}

// module: challenge, action: challengeInfo
type ChallengeInfoResp struct {
	Result     []interface{} `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}
