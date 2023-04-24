package model

// MultiUnitScenarioChapterList ...
type MultiUnitScenarioChapterList struct {
	MultiUnitScenarioID int `json:"multi_unit_scenario_id"`
	Chapter             int `json:"chapter"`
	Status              int `json:"status"`
}

// MultiUnitScenarioStatusList ...
type MultiUnitScenarioStatusList struct {
	MultiUnitID               int                            `json:"multi_unit_id"`
	Status                    int                            `json:"status"`
	MultiUnitScenarioBtnAsset string                         `json:"multi_unit_scenario_btn_asset"`
	OpenDate                  string                         `json:"open_date"`
	ChapterList               []MultiUnitScenarioChapterList `json:"chapter_list"`
}

// MultiUnitScenarioStatusRes ...
type MultiUnitScenarioStatusRes struct {
	MultiUnitScenarioStatusList  []MultiUnitScenarioStatusList `json:"multi_unit_scenario_status_list"`
	UnlockedMultiUnitScenarioIds []interface{}                 `json:"unlocked_multi_unit_scenario_ids"`
}

// MultiUnitScenarioStatusResp ...
type MultiUnitScenarioStatusResp struct {
	Result     MultiUnitScenarioStatusRes `json:"result"`
	Status     int                        `json:"status"`
	CommandNum bool                       `json:"commandNum"`
	TimeStamp  int64                      `json:"timeStamp"`
}
