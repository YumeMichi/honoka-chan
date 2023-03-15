package model

import "encoding/json"

// module: live, action: liveStatus
type NormalLiveStatusList struct {
	LiveDifficultyID   int   `json:"live_difficulty_id"`
	Status             int   `json:"status"`
	HiScore            int   `json:"hi_score"`
	HiComboCount       int   `json:"hi_combo_count"`
	ClearCnt           int   `json:"clear_cnt"`
	AchievedGoalIDList []int `json:"achieved_goal_id_list"`
}

type SpecialLiveStatusList struct {
	LiveDifficultyID   int   `json:"live_difficulty_id"`
	Status             int   `json:"status"`
	HiScore            int   `json:"hi_score"`
	HiComboCount       int   `json:"hi_combo_count"`
	ClearCnt           int   `json:"clear_cnt"`
	AchievedGoalIDList []int `json:"achieved_goal_id_list"`
}

type TrainingLiveStatusList struct {
	LiveDifficultyID   int   `json:"live_difficulty_id"`
	Status             int   `json:"status"`
	HiScore            int   `json:"hi_score"`
	HiComboCount       int   `json:"hi_combo_count"`
	ClearCnt           int   `json:"clear_cnt"`
	AchievedGoalIDList []int `json:"achieved_goal_id_list"`
}

type LiveStatusResult struct {
	NormalLiveStatusList   []NormalLiveStatusList   `json:"normal_live_status_list"`
	SpecialLiveStatusList  []SpecialLiveStatusList  `json:"special_live_status_list"`
	TrainingLiveStatusList []TrainingLiveStatusList `json:"training_live_status_list"`
	MarathonLiveStatusList []interface{}            `json:"marathon_live_status_list"`
	FreeLiveStatusList     []interface{}            `json:"free_live_status_list"`
	CanResumeLive          bool                     `json:"can_resume_live"`
}

type LiveStatusResp struct {
	Result     LiveStatusResult `json:"result"`
	Status     int              `json:"status"`
	CommandNum bool             `json:"commandNum"`
	TimeStamp  int64            `json:"timeStamp"`
}

// module: live, action: schedule
type LiveList struct {
	LiveDifficultyID int    `json:"live_difficulty_id"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	IsRandom         bool   `json:"is_random"`
}

type LimitedBonusCommonList struct {
	LiveType          int    `json:"live_type"`
	LimitedBonusType  int    `json:"limited_bonus_type"`
	LimitedBonusValue int    `json:"limited_bonus_value"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
}

type RandomLiveList struct {
	AttributeID int    `json:"attribute_id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

type TrainingLiveList struct {
	LiveDifficultyID int    `json:"live_difficulty_id"`
	StartDate        string `json:"start_date"`
	IsRandom         bool   `json:"is_random"`
}

type LiveScheduleResult struct {
	EventList              []interface{}            `json:"event_list"`
	LiveList               []LiveList               `json:"live_list"`
	LimitedBonusList       []interface{}            `json:"limited_bonus_list"`
	LimitedBonusCommonList []LimitedBonusCommonList `json:"limited_bonus_common_list"`
	RandomLiveList         []RandomLiveList         `json:"random_live_list"`
	FreeLiveList           []interface{}            `json:"free_live_list"`
	TrainingLiveList       []TrainingLiveList       `json:"training_live_list"`
}

type LiveScheduleResp struct {
	Result     LiveScheduleResult `json:"result"`
	Status     int                `json:"status"`
	CommandNum bool               `json:"commandNum"`
	TimeStamp  int64              `json:"timeStamp"`
}

// module: unit, action: unitAll
type Costume struct {
	UnitID    int  `json:"unit_id"`
	IsRankMax bool `json:"is_rank_max"`
	IsSigned  bool `json:"is_signed"`
}

type Active struct {
	UnitOwningUserID            int     `json:"unit_owning_user_id"`
	UnitID                      int     `json:"unit_id"`
	Exp                         int     `json:"exp"`
	NextExp                     int     `json:"next_exp"`
	Level                       int     `json:"level"`
	MaxLevel                    int     `json:"max_level"`
	LevelLimitID                int     `json:"level_limit_id"`
	Rank                        int     `json:"rank"`
	MaxRank                     int     `json:"max_rank"`
	Love                        int     `json:"love"`
	MaxLove                     int     `json:"max_love"`
	UnitSkillExp                int     `json:"unit_skill_exp"`
	UnitSkillLevel              int     `json:"unit_skill_level"`
	MaxHp                       int     `json:"max_hp"`
	UnitRemovableSkillCapacity  int     `json:"unit_removable_skill_capacity"`
	FavoriteFlag                bool    `json:"favorite_flag"`
	DisplayRank                 int     `json:"display_rank"`
	IsRankMax                   bool    `json:"is_rank_max"`
	IsLoveMax                   bool    `json:"is_love_max"`
	IsLevelMax                  bool    `json:"is_level_max"`
	IsSigned                    bool    `json:"is_signed"`
	IsSkillLevelMax             bool    `json:"is_skill_level_max"`
	IsRemovableSkillCapacityMax bool    `json:"is_removable_skill_capacity_max"`
	InsertDate                  string  `json:"insert_date"`
	Costume                     Costume `json:"costume,omitempty"`
}

type Waiting struct {
	UnitOwningUserID            int64  `json:"unit_owning_user_id"`
	UnitID                      int    `json:"unit_id"`
	Exp                         int    `json:"exp"`
	NextExp                     int    `json:"next_exp"`
	Level                       int    `json:"level"`
	MaxLevel                    int    `json:"max_level"`
	LevelLimitID                int    `json:"level_limit_id"`
	Rank                        int    `json:"rank"`
	MaxRank                     int    `json:"max_rank"`
	Love                        int    `json:"love"`
	MaxLove                     int    `json:"max_love"`
	UnitSkillExp                int    `json:"unit_skill_exp"`
	UnitSkillLevel              int    `json:"unit_skill_level"`
	MaxHp                       int    `json:"max_hp"`
	UnitRemovableSkillCapacity  int    `json:"unit_removable_skill_capacity"`
	FavoriteFlag                bool   `json:"favorite_flag"`
	DisplayRank                 int    `json:"display_rank"`
	IsRankMax                   bool   `json:"is_rank_max"`
	IsLoveMax                   bool   `json:"is_love_max"`
	IsLevelMax                  bool   `json:"is_level_max"`
	IsSigned                    bool   `json:"is_signed"`
	IsSkillLevelMax             bool   `json:"is_skill_level_max"`
	IsRemovableSkillCapacityMax bool   `json:"is_removable_skill_capacity_max"`
	InsertDate                  string `json:"insert_date"`
}

type UnitAllResult struct {
	Active  []Active  `json:"active"`
	Waiting []Waiting `json:"waiting"`
}

type UnitAllResp struct {
	Result     UnitAllResult `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}

// module: unit, action: deckInfo
type UnitOwningUserIds struct {
	Position         int   `json:"position"`
	UnitOwningUserID int64 `json:"unit_owning_user_id"`
}

type UnitDeckInfo struct {
	UnitDeckID        int                 `json:"unit_deck_id"`
	MainFlag          bool                `json:"main_flag"`
	DeckName          string              `json:"deck_name"`
	UnitOwningUserIds []UnitOwningUserIds `json:"unit_owning_user_ids"`
}

type UnitDeckInfoResult struct {
	UnitDeckInfo []UnitDeckInfo
}

// module: unit, action: supporterAll
type UnitSupportList struct {
	UnitID int `json:"unit_id"`
	Amount int `json:"amount"`
}

type UnitSupporterAllResult struct {
	UnitSupportList []UnitSupportList `json:"unit_support_list"`
}

// module: unit, action: removableSkillInfo

// module: costume, action: costumeList
type CostumeList struct {
	UnitID    int  `json:"unit_id"`
	IsRankMax bool `json:"is_rank_max"`
	IsSigned  bool `json:"is_signed"`
}

type CostumeListResult struct {
	CostumeList []CostumeList `json:"costume_list"`
}

// module: album, action: albumAll
type AlbumAll struct {
	UnitID             int  `json:"unit_id"`
	RankMaxFlag        bool `json:"rank_max_flag"`
	LoveMaxFlag        bool `json:"love_max_flag"`
	RankLevelMaxFlag   bool `json:"rank_level_max_flag"`
	AllMaxFlag         bool `json:"all_max_flag"`
	HighestLovePerUnit int  `json:"highest_love_per_unit"`
	TotalLove          int  `json:"total_love"`
	FavoritePoint      int  `json:"favorite_point"`
	SignFlag           bool `json:"sign_flag"`
}

type AlbumAllResult struct {
	AlbumAll []AlbumAll
}

// module: scenario, action: scenarioStatus
type ScenarioStatusList struct {
	ScenarioID int `json:"scenario_id"`
	Status     int `json:"status"`
}

type ScenarioStatusResult struct {
	ScenarioStatusList []ScenarioStatusList `json:"scenario_status_list"`
}

// module: subscenario, action: subscenarioStatus
type SubscenarioStatusList struct {
	SubscenarioID int `json:"subscenario_id"`
	Status        int `json:"status"`
}

type SubscenarioStatusResult struct {
	SubscenarioStatusList  []SubscenarioStatusList `json:"subscenario_status_list"`
	UnlockedSubscenarioIds []interface{}           `json:"unlocked_subscenario_ids"`
}

// module: eventscenario, action: status
type EventScenarioChapterList struct {
	EventScenarioID int    `json:"event_scenario_id"`
	Chapter         int    `json:"chapter"`
	ChapterAsset    string `json:"chapter_asset"`
	Status          int    `json:"status"`
	OpenFlashFlag   int    `json:"open_flash_flag"`
	IsReward        bool   `json:"is_reward"`
	CostType        int    `json:"cost_type"`
	ItemID          int    `json:"item_id"`
	Amount          int    `json:"amount"`
}

type EventScenarioList struct {
	EventID               int                        `json:"event_id"`
	EventScenarioBtnAsset string                     `json:"event_scenario_btn_asset"`
	OpenDate              string                     `json:"open_date"`
	ChapterList           []EventScenarioChapterList `json:"chapter_list"`
}

type EventScenarioStatusResult struct {
	EventScenarioList []EventScenarioList `json:"event_scenario_list"`
}

// module: multiunit, action: multiunitscenarioStatus
type MultiUnitScenarioChapterList struct {
	MultiUnitScenarioID int `json:"multi_unit_scenario_id"`
	Chapter             int `json:"chapter"`
	Status              int `json:"status"`
}

type MultiUnitScenarioStatusList struct {
	MultiUnitID               int                            `json:"multi_unit_id"`
	Status                    int                            `json:"status"`
	MultiUnitScenarioBtnAsset string                         `json:"multi_unit_scenario_btn_asset"`
	OpenDate                  string                         `json:"open_date"`
	ChapterList               []MultiUnitScenarioChapterList `json:"chapter_list"`
}

type MultiUnitScenarioStatusResult struct {
	MultiUnitScenarioStatusList  []MultiUnitScenarioStatusList `json:"multi_unit_scenario_status_list"`
	UnlockedMultiUnitScenarioIds []interface{}                 `json:"unlocked_multi_unit_scenario_ids"`
}

// module: payment, action: productList
type RestrictionInfo struct {
	Restricted bool `json:"restricted"`
}

type UnderAgeInfo struct {
	BirthSet    bool        `json:"birth_set"`
	HasLimit    bool        `json:"has_limit"`
	LimitAmount interface{} `json:"limit_amount"`
	MonthUsed   int         `json:"month_used"`
}

type SnsProductItemList struct {
	ItemID    int  `json:"item_id"`
	AddType   int  `json:"add_type"`
	Amount    int  `json:"amount"`
	IsFreebie bool `json:"is_freebie"`
}

type SnsProductList struct {
	ProductID   string               `json:"product_id"`
	Name        string               `json:"name"`
	Price       int                  `json:"price"`
	CanBuy      bool                 `json:"can_buy"`
	ProductType int                  `json:"product_type"`
	ItemList    []SnsProductItemList `json:"item_list"`
}

type ProductItemList struct {
	ItemID    int  `json:"item_id"`
	AddType   int  `json:"add_type"`
	Amount    int  `json:"amount"`
	IsFreebie bool `json:"is_freebie"`
	IsRankMax bool `json:"is_rank_max,omitempty"`
}

type LimitStatus struct {
	TermStartDate  string `json:"term_start_date"`
	RemainingTime  string `json:"remaining_time"`
	RemainingCount int    `json:"remaining_count"`
}

type ProductList struct {
	ProductID      string            `json:"product_id"`
	Name           string            `json:"name"`
	BannerImgAsset string            `json:"banner_img_asset"`
	Price          int               `json:"price"`
	CanBuy         bool              `json:"can_buy"`
	ProductType    int               `json:"product_type"`
	AnnounceURL    string            `json:"announce_url"`
	ConfirmURL     string            `json:"confirm_url"`
	ItemList       []ProductItemList `json:"item_list"`
	LimitStatus    LimitStatus       `json:"limit_status"`
}

type SubscriptionItemList struct {
	ItemID    int  `json:"item_id"`
	AddType   int  `json:"add_type"`
	Amount    int  `json:"amount"`
	IsFreebie bool `json:"is_freebie"`
}

type RewardList struct {
	ItemID  int `json:"item_id"`
	AddType int `json:"add_type"`
	Amount  int `json:"amount"`
}

type Items struct {
	Seq        int          `json:"seq"`
	RewardList []RewardList `json:"reward_list"`
}

type LicenseInfo struct {
	Name  string  `json:"name"`
	Items []Items `json:"items"`
}

type UserStatus struct {
	IsLicensed bool `json:"is_licensed"`
}

type SubscriptionStatus struct {
	LicenseID     int         `json:"license_id"`
	LicenseType   int         `json:"license_type"`
	LicenseInfo   LicenseInfo `json:"license_info,omitempty"`
	UserStatus    UserStatus  `json:"user_status"`
	PurchaseCount int         `json:"purchase_count"`
	BadgeFlag     bool        `json:"badge_flag"`
}

type SubscriptionList struct {
	ProductID          string                 `json:"product_id"`
	Name               string                 `json:"name"`
	BannerImgAsset     string                 `json:"banner_img_asset"`
	Price              int                    `json:"price"`
	CanBuy             bool                   `json:"can_buy"`
	ProductType        int                    `json:"product_type"`
	ProductURL         string                 `json:"product_url"`
	ItemList           []SubscriptionItemList `json:"item_list"`
	LimitStatus        LimitStatus            `json:"limit_status"`
	SubscriptionStatus SubscriptionStatus     `json:"subscription_status"`
}

type ProductListResult struct {
	RestrictionInfo  RestrictionInfo    `json:"restriction_info"`
	UnderAgeInfo     UnderAgeInfo       `json:"under_age_info"`
	SnsProductList   []SnsProductList   `json:"sns_product_list"`
	ProductList      []ProductList      `json:"product_list"`
	SubscriptionList []SubscriptionList `json:"subscription_list"`
	ShowPointShop    bool               `json:"show_point_shop"`
}

// module: banner, action: bannerList
type BannerList struct {
	BannerType       int    `json:"banner_type"`
	TargetID         int    `json:"target_id"`
	AssetPath        string `json:"asset_path"`
	FixedFlag        bool   `json:"fixed_flag"`
	BackSide         bool   `json:"back_side"`
	BannerID         int    `json:"banner_id"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	AddUnitStartDate string `json:"add_unit_start_date,omitempty"`
	WebviewURL       string `json:"webview_url,omitempty"`
}

type BannerListResult struct {
	TimeLimit  string       `json:"time_limit"`
	BannerList []BannerList `json:"banner_list"`
}

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

// module: stamp, action: stampInfo
type StampList struct {
	Position int `json:"position"`
	StampID  int `json:"stamp_id"`
}

type SettingList struct {
	StampSettingID int         `json:"stamp_setting_id"`
	MainFlag       int         `json:"main_flag"`
	StampList      []StampList `json:"stamp_list"`
}

type StampSetting struct {
	StampType   int           `json:"stamp_type"`
	SettingList []SettingList `json:"setting_list"`
}

type StampInfoResult struct {
	OwningStampIds []int          `json:"owning_stamp_ids"`
	StampSetting   []StampSetting `json:"stamp_setting"`
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

// response_data
type Response struct {
	ResponseData json.RawMessage `json:"response_data"`
	ReleaseInfo  []interface{}   `json:"release_info"`
	StatusCode   int             `json:"status_code"`
}
