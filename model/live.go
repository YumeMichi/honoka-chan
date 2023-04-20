package model

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

// Play
type PlayReq struct {
	Module           string `json:"module"`
	PartyUserID      int    `json:"party_user_id"`
	Action           string `json:"action"`
	Mgd              int    `json:"mgd"`
	IsTraining       bool   `json:"is_training"`
	UnitDeckID       int    `json:"unit_deck_id"`
	LiveDifficultyID string `json:"live_difficulty_id"`
	TimeStamp        int    `json:"timeStamp"`
	LpFactor         int    `json:"lp_factor"`
	CommandNum       string `json:"commandNum"`
}

type RankInfo struct {
	Rank    int `json:"rank"`
	RankMin int `json:"rank_min"`
	RankMax int `json:"rank_max"`
}

type NotesList struct {
	TimingSec      float64 `json:"timing_sec"`
	NotesAttribute int     `json:"notes_attribute"`
	NotesLevel     int     `json:"notes_level"`
	Effect         int     `json:"effect"`
	EffectValue    float64 `json:"effect_value"`
	Position       int     `json:"position"`
}

type LiveInfo struct {
	LiveDifficultyID int         `json:"live_difficulty_id"`
	IsRandom         bool        `json:"is_random"`
	AcFlag           int         `json:"ac_flag"`
	SwingFlag        int         `json:"swing_flag"`
	NotesList        []NotesList `json:"notes_list"`
}

type PlayCostume struct {
	UnitID    int  `json:"unit_id"`
	IsRankMax bool `json:"is_rank_max"`
	IsSigned  bool `json:"is_signed"`
}

type UnitList struct {
	Smile int `json:"smile"`
	Cute  int `json:"cute"`
	Cool  int `json:"cool"`
	// Costume PlayCostume `json:"costume,omitempty"`
}

type DeckInfo struct {
	UnitDeckID       int        `json:"unit_deck_id"`
	TotalSmile       int        `json:"total_smile"`
	TotalCute        int        `json:"total_cute"`
	TotalCool        int        `json:"total_cool"`
	TotalHp          int        `json:"total_hp"`
	PreparedHpDamage int        `json:"prepared_hp_damage"`
	UnitList         []UnitList `json:"unit_list"`
}

type PlayLiveList struct {
	LiveInfo LiveInfo `json:"live_info"`
	DeckInfo DeckInfo `json:"deck_info"`
}

type PlayResponseData struct {
	RankInfo            []RankInfo     `json:"rank_info"`
	EnergyFullTime      string         `json:"energy_full_time"`
	OverMaxEnergy       int            `json:"over_max_energy"`
	AvailableLiveResume bool           `json:"available_live_resume"`
	LiveList            []PlayLiveList `json:"live_list"`
	IsMarathonEvent     bool           `json:"is_marathon_event"`
	MarathonEventID     interface{}    `json:"marathon_event_id"`
	NoSkill             bool           `json:"no_skill"`
	CanActivateEffect   bool           `json:"can_activate_effect"`
	ServerTimestamp     int64          `json:"server_timestamp"`
}

// preciseScore
type PlayScoreReq struct {
	Module           string `json:"module"`
	Action           string `json:"action"`
	TimeStamp        int64  `json:"timeStamp"`
	Mgd              int    `json:"mgd"`
	LiveDifficultyID string `json:"live_difficulty_id"`
	CommandNum       string `json:"commandNum"`
}

type On struct {
	HasRecord   bool        `json:"has_record"`
	LiveInfo    LiveInfo    `json:"live_info"`
	RandomSeed  interface{} `json:"random_seed"`
	MaxCombo    interface{} `json:"max_combo"`
	UpdateDate  interface{} `json:"update_date"`
	PreciseList interface{} `json:"precise_list"`
	DeckInfo    interface{} `json:"deck_info"`
	TapAdjust   interface{} `json:"tap_adjust"`
	CanReplay   bool        `json:"can_replay"`
}

type Off struct {
	HasRecord   bool        `json:"has_record"`
	LiveInfo    LiveInfo    `json:"live_info"`
	RandomSeed  interface{} `json:"random_seed"`
	MaxCombo    interface{} `json:"max_combo"`
	UpdateDate  interface{} `json:"update_date"`
	PreciseList interface{} `json:"precise_list"`
	DeckInfo    interface{} `json:"deck_info"`
	TapAdjust   interface{} `json:"tap_adjust"`
	CanReplay   bool        `json:"can_replay"`
}

type PlayScoreResponseData struct {
	On                On         `json:"on"`
	Off               Off        `json:"off"`
	RankInfo          []RankInfo `json:"rank_info"`
	CanActivateEffect bool       `json:"can_activate_effect"`
	ServerTimestamp   int        `json:"server_timestamp"`
}

// reward
type PlayRewardReq struct {
	Module           string          `json:"module"`
	Action           string          `json:"action"`
	GoodCnt          int             `json:"good_cnt"`
	MissCnt          int             `json:"miss_cnt"`
	IsTraining       bool            `json:"is_training"`
	GreatCnt         int             `json:"great_cnt"`
	CommandNum       string          `json:"commandNum"`
	LoveCnt          int             `json:"love_cnt"`
	RemainHp         int             `json:"remain_hp"`
	MaxCombo         int             `json:"max_combo"`
	ScoreSmile       int             `json:"score_smile"`
	PerfectCnt       int             `json:"perfect_cnt"`
	BadCnt           int             `json:"bad_cnt"`
	Mgd              int             `json:"mgd"`
	EventPoint       int             `json:"event_point"`
	LiveDifficultyID int             `json:"live_difficulty_id"`
	TimeStamp        int             `json:"timeStamp"`
	PreciseScoreLog  PreciseScoreLog `json:"precise_score_log"`
	ScoreCute        int             `json:"score_cute"`
	EventID          interface{}     `json:"event_id"`
	ScoreCool        int             `json:"score_cool"`
}

type Icon struct {
	SlideID  int `json:"slide_id"`
	JustID   int `json:"just_id"`
	NormalID int `json:"normal_id"`
}

type LiveSetting struct {
	StringSize                 int     `json:"string_size"`
	PreciseScoreAutoUpdateFlag bool    `json:"precise_score_auto_update_flag"`
	SeID                       int     `json:"se_id"`
	CutinBrightness            int     `json:"cutin_brightness"`
	RandomValue                int     `json:"random_value"`
	PreciseScoreUpdateType     int     `json:"precise_score_update_type"`
	EffectFlag                 bool    `json:"effect_flag"`
	NotesSpeed                 float64 `json:"notes_speed"`
	Icon                       Icon    `json:"icon"`
	CutinType                  int     `json:"cutin_type"`
}

type PreciseList struct {
	Effect     int     `json:"effect"`
	Count      int     `json:"count"`
	Tap        float64 `json:"tap"`
	NoteNumber int     `json:"note_number"`
	Position   int     `json:"position"`
	Accuracy   int     `json:"accuracy"`
	IsSame     bool    `json:"is_same"`
}

type BackgroundScore struct {
	Smile int `json:"smile"`
	Cute  int `json:"cute"`
	Cool  int `json:"cool"`
}

type TriggerLog struct {
	ActivationRate int `json:"activation_rate"`
	Position       int `json:"position"`
}
type PreciseScoreLog struct {
	LiveSetting     LiveSetting     `json:"live_setting"`
	TapAdjust       int             `json:"tap_adjust"`
	PreciseList     []PreciseList   `json:"precise_list"`
	BackgroundScore BackgroundScore `json:"background_score"`
	IsLogOn         bool            `json:"is_log_on"`
	ScoreLog        []int           `json:"score_log"`
	IsSkillOn       bool            `json:"is_skill_on"`
	TriggerLog      []TriggerLog    `json:"trigger_log"`
	RandomSeed      int             `json:"random_seed"`
}

type RewardLiveInfo struct {
	LiveDifficultyID int  `json:"live_difficulty_id"`
	IsRandom         bool `json:"is_random"`
	AcFlag           int  `json:"ac_flag"`
	SwingFlag        int  `json:"swing_flag"`
}

type PlayerExpUnitMax struct {
	Before int `json:"before"`
	After  int `json:"after"`
}

type PlayerExpFriendMax struct {
	Before int `json:"before"`
	After  int `json:"after"`
}

type PlayerExpLpMax struct {
	Before int `json:"before"`
	After  int `json:"after"`
}

type BaseRewardInfo struct {
	PlayerExp             int                `json:"player_exp"`
	PlayerExpUnitMax      PlayerExpUnitMax   `json:"player_exp_unit_max"`
	PlayerExpFriendMax    PlayerExpFriendMax `json:"player_exp_friend_max"`
	PlayerExpLpMax        PlayerExpLpMax     `json:"player_exp_lp_max"`
	GameCoin              int                `json:"game_coin"`
	GameCoinRewardBoxFlag bool               `json:"game_coin_reward_box_flag"`
	SocialPoint           int                `json:"social_point"`
}

type LiveClear struct {
	AddType                    int           `json:"add_type"`
	Amount                     int           `json:"amount"`
	ItemCategoryID             int           `json:"item_category_id"`
	UnitID                     int           `json:"unit_id"`
	UnitOwningUserID           int64         `json:"unit_owning_user_id"`
	IsSupportMember            bool          `json:"is_support_member"`
	Exp                        int           `json:"exp"`
	NextExp                    int           `json:"next_exp"`
	MaxHp                      int           `json:"max_hp"`
	Level                      int           `json:"level"`
	MaxLevel                   int           `json:"max_level"`
	LevelLimitID               int           `json:"level_limit_id"`
	SkillLevel                 int           `json:"skill_level"`
	Rank                       int           `json:"rank"`
	Love                       int           `json:"love"`
	IsRankMax                  bool          `json:"is_rank_max"`
	IsLevelMax                 bool          `json:"is_level_max"`
	IsLoveMax                  bool          `json:"is_love_max"`
	IsSigned                   bool          `json:"is_signed"`
	NewUnitFlag                bool          `json:"new_unit_flag"`
	RewardBoxFlag              bool          `json:"reward_box_flag"`
	UnitSkillExp               int           `json:"unit_skill_exp"`
	DisplayRank                int           `json:"display_rank"`
	UnitRemovableSkillCapacity int           `json:"unit_removable_skill_capacity"`
	RemovableSkillIds          []interface{} `json:"removable_skill_ids"`
}

type LiveRank struct {
	AddType                    int           `json:"add_type"`
	Amount                     int           `json:"amount"`
	ItemCategoryID             int           `json:"item_category_id"`
	UnitID                     int           `json:"unit_id"`
	UnitOwningUserID           int64         `json:"unit_owning_user_id"`
	IsSupportMember            bool          `json:"is_support_member"`
	Exp                        int           `json:"exp"`
	NextExp                    int           `json:"next_exp"`
	MaxHp                      int           `json:"max_hp"`
	Level                      int           `json:"level"`
	MaxLevel                   int           `json:"max_level"`
	LevelLimitID               int           `json:"level_limit_id"`
	SkillLevel                 int           `json:"skill_level"`
	Rank                       int           `json:"rank"`
	Love                       int           `json:"love"`
	IsRankMax                  bool          `json:"is_rank_max"`
	IsLevelMax                 bool          `json:"is_level_max"`
	IsLoveMax                  bool          `json:"is_love_max"`
	IsSigned                   bool          `json:"is_signed"`
	NewUnitFlag                bool          `json:"new_unit_flag"`
	RewardBoxFlag              bool          `json:"reward_box_flag"`
	UnitSkillExp               int           `json:"unit_skill_exp"`
	DisplayRank                int           `json:"display_rank"`
	UnitRemovableSkillCapacity int           `json:"unit_removable_skill_capacity"`
	RemovableSkillIds          []interface{} `json:"removable_skill_ids"`
}

type RewardUnitList struct {
	LiveClear []LiveClear   `json:"live_clear"`
	LiveRank  []LiveRank    `json:"live_rank"`
	LiveCombo []interface{} `json:"live_combo"`
}

type Rewards struct {
	Rarity         int    `json:"rarity"`
	ItemID         int    `json:"item_id"`
	AddType        int    `json:"add_type"`
	Amount         int    `json:"amount"`
	ItemCategoryID int    `json:"item_category_id"`
	RewardBoxFlag  bool   `json:"reward_box_flag"`
	InsertDate     string `json:"insert_date"`
}

type EffortPoint struct {
	LiveEffortPointBoxSpecID int       `json:"live_effort_point_box_spec_id"`
	Capacity                 int       `json:"capacity"`
	Before                   int       `json:"before"`
	After                    int       `json:"after"`
	Rewards                  []Rewards `json:"rewards"`
}

// type PlayRewardUnitList struct {
// 	UnitOwningUserID int  `json:"unit_owning_user_id"`
// 	UnitID           int  `json:"unit_id"`
// 	Position         int  `json:"position"`
// 	Level            int  `json:"level"`
// 	LevelLimitID     int  `json:"level_limit_id"`
// 	DisplayRank      int  `json:"display_rank"`
// 	Love             int  `json:"love"`
// 	UnitSkillLevel   int  `json:"unit_skill_level"`
// 	IsRankMax        bool `json:"is_rank_max"`
// 	IsLoveMax        bool `json:"is_love_max"`
// 	IsLevelMax       bool `json:"is_level_max"`
// 	IsSigned         bool `json:"is_signed"`
// 	BeforeLove       int  `json:"before_love"`
// 	MaxLove          int  `json:"max_love"`
// 	// Costume Costume `json:"costume,omitempty"`
// }

type PlayRewardUnitList struct {
	ID               int   `xorm:"id pk autoincr" json:"-"`
	UserDeckID       int   `xorm:"user_deck_id" json:"-"`
	UnitOwningUserID int   `xorm:"unit_owning_user_id" json:"unit_owning_user_id"`
	UnitID           int   `xorm:"unit_id" json:"unit_id"`
	Position         int   `xorm:"position" json:"position"`
	Level            int   `xorm:"level" json:"level"`
	LevelLimitID     int   `xorm:"level_limit_id" json:"level_limit_id"`
	DisplayRank      int   `xorm:"display_rank" json:"display_rank"`
	Love             int   `xorm:"love" json:"love"`
	UnitSkillLevel   int   `xorm:"unit_skill_level" json:"unit_skill_level"`
	IsRankMax        bool  `xorm:"is_rank_max" json:"is_rank_max"`
	IsLoveMax        bool  `xorm:"is_love_max" json:"is_love_max"`
	IsLevelMax       bool  `xorm:"is_level_max" json:"is_level_max"`
	IsSigned         bool  `xorm:"is_signed" json:"is_signed"`
	BeforeLove       int   `xorm:"before_love" json:"before_love"`
	MaxLove          int   `xorm:"max_love" json:"max_love"`
	InsertData       int64 `xorm:"insert_date" json:"-"`
}

type BeforeUserInfo struct {
	Level                          int    `json:"level"`
	Exp                            int    `json:"exp"`
	PreviousExp                    int    `json:"previous_exp"`
	NextExp                        int    `json:"next_exp"`
	GameCoin                       int    `json:"game_coin"`
	SnsCoin                        int    `json:"sns_coin"`
	FreeSnsCoin                    int    `json:"free_sns_coin"`
	PaidSnsCoin                    int    `json:"paid_sns_coin"`
	SocialPoint                    int    `json:"social_point"`
	UnitMax                        int    `json:"unit_max"`
	WaitingUnitMax                 int    `json:"waiting_unit_max"`
	CurrentEnergy                  int    `json:"current_energy"`
	EnergyMax                      int    `json:"energy_max"`
	TrainingEnergy                 int    `json:"training_energy"`
	TrainingEnergyMax              int    `json:"training_energy_max"`
	EnergyFullTime                 string `json:"energy_full_time"`
	LicenseLiveEnergyRecoverlyTime int    `json:"license_live_energy_recoverly_time"`
	FriendMax                      int    `json:"friend_max"`
	TutorialState                  int    `json:"tutorial_state"`
	OverMaxEnergy                  int    `json:"over_max_energy"`
	UnlockRandomLiveMuse           int    `json:"unlock_random_live_muse"`
	UnlockRandomLiveAqours         int    `json:"unlock_random_live_aqours"`
}

type AfterUserInfo struct {
	Level                          int    `json:"level"`
	Exp                            int    `json:"exp"`
	PreviousExp                    int    `json:"previous_exp"`
	NextExp                        int    `json:"next_exp"`
	GameCoin                       int    `json:"game_coin"`
	SnsCoin                        int    `json:"sns_coin"`
	FreeSnsCoin                    int    `json:"free_sns_coin"`
	PaidSnsCoin                    int    `json:"paid_sns_coin"`
	SocialPoint                    int    `json:"social_point"`
	UnitMax                        int    `json:"unit_max"`
	WaitingUnitMax                 int    `json:"waiting_unit_max"`
	CurrentEnergy                  int    `json:"current_energy"`
	EnergyMax                      int    `json:"energy_max"`
	TrainingEnergy                 int    `json:"training_energy"`
	TrainingEnergyMax              int    `json:"training_energy_max"`
	EnergyFullTime                 string `json:"energy_full_time"`
	LicenseLiveEnergyRecoverlyTime int    `json:"license_live_energy_recoverly_time"`
	FriendMax                      int    `json:"friend_max"`
	TutorialState                  int    `json:"tutorial_state"`
	OverMaxEnergy                  int    `json:"over_max_energy"`
	UnlockRandomLiveMuse           int    `json:"unlock_random_live_muse"`
	UnlockRandomLiveAqours         int    `json:"unlock_random_live_aqours"`
}

type NextLevelInfo struct {
	Level   int `json:"level"`
	FromExp int `json:"from_exp"`
}

type GoalAccompInfo struct {
	AchievedIds []interface{} `json:"achieved_ids"`
	Rewards     []interface{} `json:"rewards"`
}

type RewardRankInfo struct {
	BeforeClassRankID int    `json:"before_class_rank_id"`
	AfterClassRankID  int    `json:"after_class_rank_id"`
	RankUpDate        string `json:"rank_up_date"`
}

type ClassSystem struct {
	RankInfo     RewardRankInfo `json:"rank_info"`
	CompleteFlag bool           `json:"complete_flag"`
	IsOpened     bool           `json:"is_opened"`
	IsVisible    bool           `json:"is_visible"`
}

type PlayRewardList struct {
	ItemID         int  `json:"item_id"`
	AddType        int  `json:"add_type"`
	Amount         int  `json:"amount"`
	ItemCategoryID int  `json:"item_category_id"`
	RewardBoxFlag  bool `json:"reward_box_flag"`
}

type AccomplishedAchievementList struct {
	AchievementID       int              `json:"achievement_id"`
	Count               int              `json:"count"`
	IsAccomplished      bool             `json:"is_accomplished"`
	InsertDate          string           `json:"insert_date"`
	EndDate             string           `json:"end_date"`
	RemainingTime       string           `json:"remaining_time"`
	IsNew               bool             `json:"is_new"`
	ForDisplay          bool             `json:"for_display"`
	IsLocked            bool             `json:"is_locked"`
	OpenConditionString string           `json:"open_condition_string"`
	AccomplishID        string           `json:"accomplish_id"`
	RewardList          []PlayRewardList `json:"reward_list"`
}

type Parameter struct {
	Smile int `json:"smile"`
	Pure  int `json:"pure"`
	Cool  int `json:"cool"`
}

type RewardMuseumInfo struct {
	Parameter      Parameter `json:"parameter"`
	ContentsIDList []int     `json:"contents_id_list"`
}

type RewardUnitSupportList struct {
	UnitID int `json:"unit_id"`
	Amount int `json:"amount"`
}

type RewardResponseData struct {
	LiveInfo                     []RewardLiveInfo              `json:"live_info"`
	Rank                         int                           `json:"rank"`
	ComboRank                    int                           `json:"combo_rank"`
	TotalLove                    int                           `json:"total_love"`
	IsHighScore                  bool                          `json:"is_high_score"`
	HiScore                      int                           `json:"hi_score"`
	BaseRewardInfo               BaseRewardInfo                `json:"base_reward_info"`
	RewardUnitList               RewardUnitList                `json:"reward_unit_list"`
	UnlockedSubscenarioIds       []interface{}                 `json:"unlocked_subscenario_ids"`
	UnlockedMultiUnitScenarioIds []interface{}                 `json:"unlocked_multi_unit_scenario_ids"`
	EffortPoint                  []EffortPoint                 `json:"effort_point"`
	IsEffortPointVisible         bool                          `json:"is_effort_point_visible"`
	LimitedEffortBox             []interface{}                 `json:"limited_effort_box"`
	UnitList                     []PlayRewardUnitList          `json:"unit_list"`
	BeforeUserInfo               BeforeUserInfo                `json:"before_user_info"`
	AfterUserInfo                AfterUserInfo                 `json:"after_user_info"`
	NextLevelInfo                []NextLevelInfo               `json:"next_level_info"`
	GoalAccompInfo               GoalAccompInfo                `json:"goal_accomp_info"`
	SpecialRewardInfo            []interface{}                 `json:"special_reward_info"`
	EventInfo                    []interface{}                 `json:"event_info"`
	DailyRewardInfo              []interface{}                 `json:"daily_reward_info"`
	CanSendFriendRequest         bool                          `json:"can_send_friend_request"`
	UsingBuffInfo                []interface{}                 `json:"using_buff_info"`
	ClassSystem                  ClassSystem                   `json:"class_system"`
	AccomplishedAchievementList  []AccomplishedAchievementList `json:"accomplished_achievement_list"`
	UnaccomplishedAchievementCnt int                           `json:"unaccomplished_achievement_cnt"`
	AddedAchievementList         []interface{}                 `json:"added_achievement_list"`
	MuseumInfo                   RewardMuseumInfo              `json:"museum_info"`
	UnitSupportList              []RewardUnitSupportList       `json:"unit_support_list"`
	ServerTimestamp              int                           `json:"server_timestamp"`
	PresentCnt                   int                           `json:"present_cnt"`
}
