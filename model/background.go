package model

// BackgroundSetResp ...
type BackgroundSetResp struct {
	ResponseData []interface{} `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

// BackgroundInfo ...
type BackgroundInfo struct {
	BackgroundID int    `json:"background_id"`
	IsSet        bool   `json:"is_set"`
	InsertDate   string `json:"insert_date"`
}

// BackgroundInfoRes ...
type BackgroundInfoRes struct {
	BackgroundInfo []BackgroundInfo `json:"background_info"`
}

// BackgroundInfoResp ...
type BackgroundInfoResp struct {
	Result     BackgroundInfoRes `json:"result"`
	Status     int               `json:"status"`
	CommandNum bool              `json:"commandNum"`
	TimeStamp  int64             `json:"timeStamp"`
}
