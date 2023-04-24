package model

// ExchangePointList ...
type ExchangePointList struct {
	Rarity        int `json:"rarity"`
	ExchangePoint int `json:"exchange_point"`
}

// ExchangePointRes ...
type ExchangePointRes struct {
	ExchangePointList []ExchangePointList `json:"exchange_point_list"`
}

// ExchangePointResp ...
type ExchangePointResp struct {
	Result     ExchangePointRes `json:"result"`
	Status     int              `json:"status"`
	CommandNum bool             `json:"commandNum"`
	TimeStamp  int64            `json:"timeStamp"`
}
