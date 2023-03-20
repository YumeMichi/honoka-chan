package model

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
