package model

// MarathonInfoResp ...
type MarathonInfoResp struct {
	Result     []interface{} `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}
