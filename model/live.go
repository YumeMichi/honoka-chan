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
