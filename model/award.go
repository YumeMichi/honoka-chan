package model

// AwardSetResp ...
type AwardSetResp struct {
	ResponseData []interface{} `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

// AwardInfo ...
type AwardInfo struct {
	AwardID    int    `json:"award_id"`
	IsSet      bool   `json:"is_set"`
	InsertDate string `json:"insert_date"`
}

// AwardInfoRes ...
type AwardInfoRes struct {
	AwardInfo []AwardInfo `json:"award_info"`
}

// AwardInfoResp ...
type AwardInfoResp struct {
	Result     AwardInfoRes `json:"result"`
	Status     int          `json:"status"`
	CommandNum bool         `json:"commandNum"`
	TimeStamp  int64        `json:"timeStamp"`
}
