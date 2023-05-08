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

// MultiUnitStartUpResp ...
type MultiUnitStartUpResp struct {
	ResponseData MultiUnitStartUpRes `json:"response_data"`
	ReleaseInfo  []interface{}       `json:"release_info"`
	StatusCode   int                 `json:"status_code"`
}

// MultiUnitStartUpRes ...
type MultiUnitStartUpRes struct {
	MultiUnitScenarioID int   `json:"multi_unit_scenario_id"`
	ScenarioAdjustment  int   `json:"scenario_adjustment"`
	ServerTimestamp     int64 `json:"server_timestamp"`
}

// MultiUnitStartUpReq ...
type MultiUnitStartUpReq struct {
	Module              string `json:"module"`
	Action              string `json:"action"`
	TimeStamp           int    `json:"timeStamp"`
	Mgd                 int    `json:"mgd"`
	MultiUnitScenarioID int    `json:"multi_unit_scenario_id"`
	CommandNum          string `json:"commandNum"`
}
