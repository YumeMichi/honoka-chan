package model

// UserInfoResp ...
type UserInfoResp struct {
	ResponseData UserInfoRes `json:"response_data"`
	ReleaseInfo  []any       `json:"release_info"`
	StatusCode   int         `json:"status_code"`
}

// UserInfoRes ...
type UserInfoRes struct {
	User            UserInfo `json:"user"`
	Birth           Birth    `json:"birth"`
	ServerTimestamp int64    `json:"server_timestamp"`
}

// LpRecoveryItem ...
type LpRecoveryItem struct {
	ItemID int `json:"item_id"`
	Amount int `json:"amount"`
}

// UserInfo ...
type UserInfo struct {
	UserID                         int              `json:"user_id"`
	Name                           string           `json:"name"`
	Level                          int              `json:"level"`
	Exp                            int              `json:"exp"`
	PreviousExp                    int              `json:"previous_exp"`
	NextExp                        int              `json:"next_exp"`
	GameCoin                       int              `json:"game_coin"`
	SnsCoin                        int              `json:"sns_coin"`
	FreeSnsCoin                    int              `json:"free_sns_coin"`
	PaidSnsCoin                    int              `json:"paid_sns_coin"`
	SocialPoint                    int              `json:"social_point"`
	UnitMax                        int              `json:"unit_max"`
	WaitingUnitMax                 int              `json:"waiting_unit_max"`
	EnergyMax                      int              `json:"energy_max"`
	EnergyFullTime                 string           `json:"energy_full_time"`
	LicenseLiveEnergyRecoverlyTime int              `json:"license_live_energy_recoverly_time"`
	EnergyFullNeedTime             int              `json:"energy_full_need_time"`
	OverMaxEnergy                  int              `json:"over_max_energy"`
	TrainingEnergy                 int              `json:"training_energy"`
	TrainingEnergyMax              int              `json:"training_energy_max"`
	FriendMax                      int              `json:"friend_max"`
	InviteCode                     string           `json:"invite_code"`
	InsertDate                     string           `json:"insert_date"`
	UpdateDate                     string           `json:"update_date"`
	TutorialState                  int              `json:"tutorial_state"`
	DiamondCoin                    int              `json:"diamond_coin"`
	CrystalCoin                    int              `json:"crystal_coin"`
	LpRecoveryItem                 []LpRecoveryItem `json:"lp_recovery_item"`
}

// Birth ...
type Birth struct {
	BirthMonth int `json:"birth_month"`
	BirthDay   int `json:"birth_day"`
}
