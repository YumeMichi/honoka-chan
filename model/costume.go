package model

// CostumeList ...
type CostumeList struct {
	UnitID    int  `json:"unit_id"`
	IsRankMax bool `json:"is_rank_max"`
	IsSigned  bool `json:"is_signed"`
}

// CostumeListRes ...
type CostumeListRes struct {
	CostumeList []CostumeList `json:"costume_list"`
}

// CostumeListResp ...
type CostumeListResp struct {
	Result     CostumeListRes `json:"result"`
	Status     int            `json:"status"`
	CommandNum bool           `json:"commandNum"`
	TimeStamp  int64          `json:"timeStamp"`
}
