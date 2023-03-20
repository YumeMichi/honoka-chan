package model

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
