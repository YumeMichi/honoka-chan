package model

// SubScenarioResp ...
type SubScenarioResp struct {
	ResponseData SubScenarioRes `json:"response_data"`
	ReleaseInfo  []any          `json:"release_info"`
	StatusCode   int            `json:"status_code"`
}

// SubScenarioRes ...
type SubScenarioRes struct {
	SubscenarioID      int   `json:"subscenario_id"`
	ScenarioAdjustment int   `json:"scenario_adjustment"`
	ServerTimestamp    int64 `json:"server_timestamp"`
}

// SubScenarioReq ...
type SubScenarioReq struct {
	Module        string `json:"module"`
	Action        string `json:"action"`
	TimeStamp     int    `json:"timeStamp"`
	SubscenarioID int    `json:"subscenario_id"`
	Mgd           int    `json:"mgd"`
	CommandNum    string `json:"commandNum"`
}
