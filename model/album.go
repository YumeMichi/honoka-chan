package model

// module: album, action: albumAll
type AlbumAll struct {
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

type AlbumAllResult struct {
	AlbumAll []AlbumAll
}
