package model

// AlbumResult ...
type AlbumResult struct {
	UnitID             int  `json:"unit_id"`
	RankMaxFlag        bool `json:"rank_max_flag"`
	LoveMaxFlag        bool `json:"love_max_flag"`
	RankLevelMaxFlag   bool `json:"rank_level_max_flag"`
	AllMaxFlag         bool `json:"all_max_flag"`
	HighestLovePerUnit int  `json:"highest_love_per_unit"`
	TotalLove          int  `json:"total_love"`
	FavoritePoint      int  `json:"favorite_point"`
	SignFlag           bool `json:"sign_flag"`
}

// AlbumResp ...
type AlbumResp struct {
	Result     []AlbumResult `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}

// AlbumSeriesRes ...
type AlbumSeriesRes struct {
	SeriesID int           `json:"series_id"`
	UnitList []AlbumResult `json:"unit_list"`
}

// AlbumSeriesResp ...
type AlbumSeriesResp struct {
	ResponseData []AlbumSeriesRes `json:"response_data"`
	ReleaseInfo  []interface{}    `json:"release_info"`
	StatusCode   int              `json:"status_code"`
}
