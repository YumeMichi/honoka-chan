package model

// SpecialCutinRes ...
type SpecialCutinRes struct {
	SpecialCutinList []any `json:"special_cutin_list"`
}

// SpecialCutinResp ...
type SpecialCutinResp struct {
	Result     SpecialCutinRes `json:"result"`
	Status     int             `json:"status"`
	CommandNum bool            `json:"commandNum"`
	TimeStamp  int64           `json:"timeStamp"`
}
