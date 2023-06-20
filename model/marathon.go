package model

// MarathonInfoResp ...
type MarathonInfoResp struct {
	Result     []any `json:"result"`
	Status     int   `json:"status"`
	CommandNum bool  `json:"commandNum"`
	TimeStamp  int64 `json:"timeStamp"`
}
