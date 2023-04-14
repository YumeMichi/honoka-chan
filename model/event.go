package model

// EventsResp ...
type EventsResp struct {
	ResponseData EventsRes     `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

// EventsRes ...
type EventsRes struct {
	TargetList      []TargetList `json:"target_list"`
	ServerTimestamp int64        `json:"server_timestamp"`
}

// TargetList ...
type TargetList struct {
	Position      int  `json:"position"`
	IsDisplayable bool `json:"is_displayable"`
}
