package model

// AccessoryList ...
type AccessoryList struct {
	AccessoryOwningUserID int  `json:"accessory_owning_user_id" xorm:"accessory_owning_user_id"`
	AccessoryID           int  `json:"accessory_id" xorm:"accessory_id"`
	Exp                   int  `json:"exp" xorm:"exp"`
	NextExp               int  `json:"next_exp" xorm:"-"`
	Level                 int  `json:"level" xorm:"-"`
	MaxLevel              int  `json:"max_level" xorm:"-"`
	RankUpCount           int  `json:"rank_up_count" xorm:"-"`
	FavoriteFlag          bool `json:"favorite_flag" xorm:"-"`
}

// WearingInfo
type WearingInfo struct {
	UnitOwningUserID      int `json:"unit_owning_user_id" xorm:"unit_owning_user_id"`
	AccessoryOwningUserID int `json:"accessory_owning_user_id" xorm:"accessory_owning_user_id"`
}

// UnitAccessoryAllResult ...
type UnitAccessoryAllResult struct {
	AccessoryList      []AccessoryList `json:"accessory_list"`
	WearingInfo        []WearingInfo   `json:"wearing_info"`
	EspecialCreateFlag bool            `json:"especial_create_flag"`
}

// UnitAccessoryAllResp ...
type UnitAccessoryAllResp struct {
	Result     UnitAccessoryAllResult `json:"result"`
	Status     int                    `json:"status"`
	CommandNum bool                   `json:"commandNum"`
	TimeStamp  int64                  `json:"timeStamp"`
}

// WearAccessoryReq ...
type WearAccessoryReq struct {
	Module     string   `json:"module"`
	Remove     []Remove `json:"remove"`
	Action     string   `json:"action"`
	TimeStamp  int      `json:"timeStamp"`
	Wear       []Wear   `json:"wear"`
	Mgd        int      `json:"mgd"`
	CommandNum string   `json:"commandNum"`
}

// Remove ...
type Remove struct {
	AccessoryOwningUserID int `json:"accessory_owning_user_id"`
	UnitOwningUserID      int `json:"unit_owning_user_id"`
}

// Wear ...
type Wear struct {
	AccessoryOwningUserID int `json:"accessory_owning_user_id"`
	UnitOwningUserID      int `json:"unit_owning_user_id"`
}

// AccessoryWearData ...
type AccessoryWearData struct {
	Id                    int    `xorm:"id pk autoincr"`
	AccessoryOwningUserID int    `xorm:"accessory_owning_user_id"`
	UnitOwningUserID      int    `xorm:"unit_owning_user_id"`
	UserId                string `xorm:"user_id"`
}

// AccessoryWearResp ...
type AccessoryWearResp struct {
	ResponseData []interface{} `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

// module: unit, action: unitAll
type Costume struct {
	UnitID    int  `json:"unit_id"`
	IsRankMax bool `json:"is_rank_max"`
	IsSigned  bool `json:"is_signed"`
}

type Active struct {
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
	// Costume                     Costume `json:"costume,omitempty"`
}

type Waiting struct {
	UnitOwningUserID            int64  `json:"unit_owning_user_id"`
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
}

type UnitAllResult struct {
	Active  []Active  `json:"active"`
	Waiting []Waiting `json:"waiting"`
}

type UnitAllResp struct {
	Result     UnitAllResult `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}

// module: unit, action: deckInfo
type UnitOwningUserIds struct {
	Position         int `json:"position"`
	UnitOwningUserID int `json:"unit_owning_user_id"`
}

type UnitDeckInfo struct {
	UnitDeckID        int                 `json:"unit_deck_id"`
	MainFlag          bool                `json:"main_flag"`
	DeckName          string              `json:"deck_name"`
	UnitOwningUserIds []UnitOwningUserIds `json:"unit_owning_user_ids"`
}

type UnitDeckInfoResp struct {
	Result     []UnitDeckInfo `json:"result"`
	Status     int            `json:"status"`
	CommandNum bool           `json:"commandNum"`
	TimeStamp  int64          `json:"timeStamp"`
}

// module: unit, action: supporterAll
type UnitSupportList struct {
	UnitID int `json:"unit_id"`
	Amount int `json:"amount"`
}

type UnitSupportResult struct {
	UnitSupportList []UnitSupportList `json:"unit_support_list"`
}

type UnitSupportResp struct {
	Result     UnitSupportResult `json:"result"`
	Status     int               `json:"status"`
	CommandNum bool              `json:"commandNum"`
	TimeStamp  int64             `json:"timeStamp"`
}

// module: unit, action: removableSkillInfo
type OwningInfo struct {
	UnitRemovableSkillID int    `json:"unit_removable_skill_id"`
	TotalAmount          int    `json:"total_amount"`
	EquippedAmount       int    `json:"equipped_amount"`
	InsertDate           string `json:"insert_date"`
}

type RemovableSkillResult struct {
	OwningInfo    []OwningInfo  `json:"owning_info"`
	EquipmentInfo []interface{} `json:"equipment_info"`
}

type RemovableSkillResp struct {
	Result     RemovableSkillResult `json:"result"`
	Status     int                  `json:"status"`
	CommandNum bool                 `json:"commandNum"`
	TimeStamp  int64                `json:"timeStamp"`
}

// module: unit, action: deck
type UnitDeckReq struct {
	Module       string         `json:"module"`
	UnitDeckList []UnitDeckList `json:"unit_deck_list"`
	Action       string         `json:"action"`
	Mgd          int            `json:"mgd"`
}

type UnitDeckDetail struct {
	Position         int `json:"position"`
	UnitOwningUserID int `json:"unit_owning_user_id"`
}

type UnitDeckList struct {
	UnitDeckDetail []UnitDeckDetail `json:"unit_deck_detail"`
	UnitDeckID     int              `json:"unit_deck_id"`
	MainFlag       int              `json:"main_flag"`
	DeckName       string           `json:"deck_name"`
}

// module: unit, action: deckName
type DeckNameReq struct {
	Module     string `json:"module"`
	UnitDeckID int    `json:"unit_deck_id"`
	Action     string `json:"action"`
	TimeStamp  int    `json:"timeStamp"`
	Mgd        int    `json:"mgd"`
	CommandNum string `json:"commandNum"`
	DeckName   string `json:"deck_name"`
}
