package model

// ChallengeInfoResp ...
type ChallengeInfoResp struct {
	Result     []interface{} `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}
