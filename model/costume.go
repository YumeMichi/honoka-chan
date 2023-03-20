package model

// module: costume, action: costumeList
type CostumeList struct {
	UnitID    int  `json:"unit_id"`
	IsRankMax bool `json:"is_rank_max"`
	IsSigned  bool `json:"is_signed"`
}

type CostumeListResult struct {
	CostumeList []CostumeList `json:"costume_list"`
}
