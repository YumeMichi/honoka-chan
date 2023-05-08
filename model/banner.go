package model

// BannerList ...
type BannerList struct {
	BannerType       int    `json:"banner_type"`
	TargetID         int    `json:"target_id"`
	AssetPath        string `json:"asset_path"`
	FixedFlag        bool   `json:"fixed_flag"`
	BackSide         bool   `json:"back_side"`
	BannerID         int    `json:"banner_id"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	AddUnitStartDate string `json:"add_unit_start_date,omitempty"`
	WebviewURL       string `json:"webview_url,omitempty"`
}

// BannerListRes ...
type BannerListRes struct {
	TimeLimit  string       `json:"time_limit"`
	BannerList []BannerList `json:"banner_list"`
}

// BannerListResp ...
type BannerListResp struct {
	Result     BannerListRes `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}
