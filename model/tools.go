package model

// UnitData ...
type UnitData struct {
	UnitOwningUserID            int    `xorm:"unit_owning_user_id pk autoincr" json:"unit_owning_user_id"`
	UserID                      int    `xorm:"user_id" json:"-"`
	UnitID                      int    `xorm:"unit_id" json:"unit_id"`
	Exp                         int    `xorm:"exp" json:"exp"`
	NextExp                     int    `xorm:"next_exp" json:"next_exp"`
	Level                       int    `xorm:"level" json:"level"`
	MaxLevel                    int    `xorm:"max_level" json:"max_level"`
	LevelLimitID                int    `xorm:"level_limit_id" json:"level_limit_id"`
	Rank                        int    `xorm:"rank" json:"rank"`
	MaxRank                     int    `xorm:"max_rank" json:"max_rank"`
	Love                        int    `xorm:"love" json:"love"`
	MaxLove                     int    `xorm:"max_love" json:"max_love"`
	UnitSkillExp                int    `xorm:"unit_skill_exp" json:"unit_skill_exp"`
	UnitSkillLevel              int    `xorm:"unit_skill_level" json:"unit_skill_level"`
	MaxHp                       int    `xorm:"max_hp" json:"max_hp"`
	UnitRemovableSkillCapacity  int    `xorm:"unit_removable_skill_capacity" json:"unit_removable_skill_capacity"`
	FavoriteFlag                bool   `xorm:"favorite_flag" json:"favorite_flag"`
	DisplayRank                 int    `xorm:"display_rank" json:"display_rank"`
	IsRankMax                   bool   `xorm:"is_rank_max" json:"is_rank_max"`
	IsLoveMax                   bool   `xorm:"is_love_max" json:"is_love_max"`
	IsLevelMax                  bool   `xorm:"is_level_max" json:"is_level_max"`
	IsSigned                    bool   `xorm:"is_signed" json:"is_signed"`
	IsSkillLevelMax             bool   `xorm:"is_skill_level_max" json:"is_skill_level_max"`
	IsRemovableSkillCapacityMax bool   `xorm:"is_removable_skill_capacity_max" json:"is_removable_skill_capacity_max"`
	InsertDate                  string `xorm:"insert_date" json:"insert_date"`
}

// UserDeckData ...
type UserDeckData struct {
	ID         int    `xorm:"id pk autoincr"`
	DeckID     int    `xorm:"deck_id"`
	MainFlag   int    `xorm:"main_flag"`
	DeckName   string `xorm:"deck_name"`
	UserID     int    `xorm:"user_id"`
	InsertDate int64  `xorm:"insert_date"`
}

// UnitDeckData ...
type UnitDeckData struct {
	ID               int   `xorm:"id pk autoincr" json:"-"`
	UserDeckID       int   `xorm:"user_deck_id" json:"-"`
	UnitOwningUserID int   `xorm:"unit_owning_user_id" json:"unit_owning_user_id"`
	UnitID           int   `xorm:"unit_id" json:"unit_id"`
	Position         int   `xorm:"position" json:"position"`
	Level            int   `xorm:"level" json:"level"`
	LevelLimitID     int   `xorm:"level_limit_id" json:"level_limit_id"`
	DisplayRank      int   `xorm:"display_rank" json:"display_rank"`
	Love             int   `xorm:"love" json:"love"`
	UnitSkillLevel   int   `xorm:"unit_skill_level" json:"unit_skill_level"`
	IsRankMax        bool  `xorm:"is_rank_max" json:"is_rank_max"`
	IsLoveMax        bool  `xorm:"is_love_max" json:"is_love_max"`
	IsLevelMax       bool  `xorm:"is_level_max" json:"is_level_max"`
	IsSigned         bool  `xorm:"is_signed" json:"is_signed"`
	BeforeLove       int   `xorm:"before_love" json:"before_love"`
	MaxLove          int   `xorm:"max_love" json:"max_love"`
	InsertData       int64 `xorm:"insert_date" json:"-"`
}

// DifficultyRes ...
type DifficultyRes struct {
	Difficulty int `json:"difficulty"`
	ClearCnt   int `json:"clear_cnt"`
}

// DifficultyResp ...
type DifficultyResp struct {
	Result     []DifficultyRes `json:"result"`
	Status     int             `json:"status"`
	CommandNum bool            `json:"commandNum"`
	TimeStamp  int64           `json:"timeStamp"`
}

// LoveResp ...
type LoveResp struct {
	Result     []interface{} `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}

// AccessoryInfo ...
type AccessoryInfo struct {
	AccessoryOwningUserID int  `json:"accessory_owning_user_id"`
	AccessoryID           int  `json:"accessory_id"`
	Exp                   int  `json:"exp"`
	NextExp               int  `json:"next_exp"`
	Level                 int  `json:"level"`
	MaxLevel              int  `json:"max_level"`
	RankUpCount           int  `json:"rank_up_count"`
	FavoriteFlag          bool `json:"favorite_flag"`
}

// ProfileUserInfo ...
type ProfileUserInfo struct {
	UserID               int    `json:"user_id"`
	Name                 string `json:"name"`
	Level                int    `json:"level"`
	CostMax              int    `json:"cost_max"`
	UnitMax              int    `json:"unit_max"`
	EnergyMax            int    `json:"energy_max"`
	FriendMax            int    `json:"friend_max"`
	UnitCnt              int    `json:"unit_cnt"`
	InviteCode           string `json:"invite_code"`
	ElapsedTimeFromLogin string `json:"elapsed_time_from_login"`
	Introduction         string `json:"introduction"`
}

// CenterUnitInfo ...
type CenterUnitInfo struct {
	UnitOwningUserID           int           `json:"unit_owning_user_id"`
	UnitID                     int           `json:"unit_id"`
	Exp                        int           `json:"exp"`
	NextExp                    int           `json:"next_exp"`
	Level                      int           `json:"level"`
	LevelLimitID               int           `json:"level_limit_id"`
	MaxLevel                   int           `json:"max_level"`
	Rank                       int           `json:"rank"`
	MaxRank                    int           `json:"max_rank"`
	Love                       int           `json:"love"`
	MaxLove                    int           `json:"max_love"`
	UnitSkillLevel             int           `json:"unit_skill_level"`
	MaxHp                      int           `json:"max_hp"`
	FavoriteFlag               bool          `json:"favorite_flag"`
	DisplayRank                int           `json:"display_rank"`
	UnitSkillExp               int           `json:"unit_skill_exp"`
	UnitRemovableSkillCapacity int           `json:"unit_removable_skill_capacity"`
	Attribute                  int           `json:"attribute"`
	Smile                      int           `json:"smile"`
	Cute                       int           `json:"cute"`
	Cool                       int           `json:"cool"`
	IsLoveMax                  bool          `json:"is_love_max"`
	IsLevelMax                 bool          `json:"is_level_max"`
	IsRankMax                  bool          `json:"is_rank_max"`
	IsSigned                   bool          `json:"is_signed"`
	IsSkillLevelMax            bool          `json:"is_skill_level_max"`
	SettingAwardID             int           `json:"setting_award_id"`
	RemovableSkillIds          []int         `json:"removable_skill_ids"`
	AccessoryInfo              AccessoryInfo `json:"accessory_info"`
	Costume                    Costume       `json:"costume"`
	TotalSmile                 int           `json:"total_smile"`
	TotalCute                  int           `json:"total_cute"`
	TotalCool                  int           `json:"total_cool"`
	TotalHp                    int           `json:"total_hp"`
}

// NaviUnitInfo ...
type NaviUnitInfo struct {
	UnitOwningUserID            int    `json:"unit_owning_user_id"`
	UnitID                      int    `json:"unit_id"`
	Exp                         int    `json:"exp"`
	NextExp                     int    `json:"next_exp"`
	Level                       int    `json:"level"`
	MaxLevel                    int    `json:"max_level"`
	LevelLimitID                int    `json:"level_limit_id"`
	Rank                        int    `json:"rank"`
	MaxRank                     int    `json:"max_rank"`
	Love                        int    `json:"love"`
	MaxLove                     int    `json:"max_love"`
	UnitSkillExp                int    `json:"unit_skill_exp"`
	UnitSkillLevel              int    `json:"unit_skill_level"`
	MaxHp                       int    `json:"max_hp"`
	UnitRemovableSkillCapacity  int    `json:"unit_removable_skill_capacity"`
	FavoriteFlag                bool   `json:"favorite_flag"`
	DisplayRank                 int    `json:"display_rank"`
	IsRankMax                   bool   `json:"is_rank_max"`
	IsLoveMax                   bool   `json:"is_love_max"`
	IsLevelMax                  bool   `json:"is_level_max"`
	IsSigned                    bool   `json:"is_signed"`
	IsSkillLevelMax             bool   `json:"is_skill_level_max"`
	IsRemovableSkillCapacityMax bool   `json:"is_removable_skill_capacity_max"`
	InsertDate                  string `json:"insert_date"`
	TotalSmile                  int    `json:"total_smile"`
	TotalCute                   int    `json:"total_cute"`
	TotalCool                   int    `json:"total_cool"`
	TotalHp                     int    `json:"total_hp"`
	RemovableSkillIds           []int  `json:"removable_skill_ids"`
}

// ProfileRes ...
type ProfileRes struct {
	UserInfo            ProfileUserInfo `json:"user_info"`
	CenterUnitInfo      CenterUnitInfo  `json:"center_unit_info"`
	NaviUnitInfo        NaviUnitInfo    `json:"navi_unit_info"`
	IsAlliance          bool            `json:"is_alliance"`
	FriendStatus        int             `json:"friend_status"`
	SettingAwardID      int             `json:"setting_award_id"`
	SettingBackgroundID int             `json:"setting_background_id"`
}

// ProfileResp ...
type ProfileResp struct {
	Result     ProfileRes `json:"result"`
	Status     int        `json:"status"`
	CommandNum bool       `json:"commandNum"`
	TimeStamp  int64      `json:"timeStamp"`
}
