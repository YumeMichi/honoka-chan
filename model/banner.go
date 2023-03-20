package model

// module: banner, action: bannerList
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

type BannerListResult struct {
	TimeLimit  string       `json:"time_limit"`
	BannerList []BannerList `json:"banner_list"`
}
