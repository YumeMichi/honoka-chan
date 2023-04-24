package model

// ScenarioStatusList ...
type ScenarioStatusList struct {
	ScenarioID int `json:"scenario_id"`
	Status     int `json:"status"`
}

// ScenarioStatusRes ...
type ScenarioStatusRes struct {
	ScenarioStatusList []ScenarioStatusList `json:"scenario_status_list"`
}

// ScenarioStatusResp ...
type ScenarioStatusResp struct {
	Result     ScenarioStatusRes `json:"result"`
	Status     int               `json:"status"`
	CommandNum bool              `json:"commandNum"`
	TimeStamp  int64             `json:"timeStamp"`
}

// SubscenarioStatusList ...
type SubscenarioStatusList struct {
	SubscenarioID int `json:"subscenario_id"`
	Status        int `json:"status"`
}

// SubscenarioStatusRes ...
type SubscenarioStatusRes struct {
	SubscenarioStatusList  []SubscenarioStatusList `json:"subscenario_status_list"`
	UnlockedSubscenarioIds []interface{}           `json:"unlocked_subscenario_ids"`
}

// SubscenarioStatusResp ...
type SubscenarioStatusResp struct {
	Result     SubscenarioStatusRes `json:"result"`
	Status     int                  `json:"status"`
	CommandNum bool                 `json:"commandNum"`
	TimeStamp  int64                `json:"timeStamp"`
}

// EventScenarioChapterList ...
type EventScenarioChapterList struct {
	EventScenarioID int    `json:"event_scenario_id"`
	Chapter         int    `json:"chapter"`
	ChapterAsset    string `json:"chapter_asset,omitempty"`
	Status          int    `json:"status"`
	OpenFlashFlag   int    `json:"open_flash_flag"`
	IsReward        bool   `json:"is_reward"`
	CostType        int    `json:"cost_type"`
	ItemID          int    `json:"item_id"`
	Amount          int    `json:"amount"`
}

// EventScenarioList ...
type EventScenarioList struct {
	EventID               int                        `json:"event_id"`
	EventScenarioBtnAsset string                     `json:"event_scenario_btn_asset"`
	OpenDate              string                     `json:"open_date"`
	ChapterList           []EventScenarioChapterList `json:"chapter_list"`
}

// EventScenarioStatusRes ...
type EventScenarioStatusRes struct {
	EventScenarioList []EventScenarioList `json:"event_scenario_list"`
}

// EventScenarioStatusResp ...
type EventScenarioStatusResp struct {
	Result     EventScenarioStatusRes `json:"result"`
	Status     int                    `json:"status"`
	CommandNum bool                   `json:"commandNum"`
	TimeStamp  int64                  `json:"timeStamp"`
}
