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
	ServerTimestamp     int            `json:"server_timestamp"`
}
