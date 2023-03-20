package model

// module: stamp, action: stampInfo
type StampList struct {
	Position int `json:"position"`
	StampID  int `json:"stamp_id"`
}

type SettingList struct {
	StampSettingID int         `json:"stamp_setting_id"`
	MainFlag       int         `json:"main_flag"`
	StampList      []StampList `json:"stamp_list"`
}

type StampSetting struct {
	StampType   int           `json:"stamp_type"`
	SettingList []SettingList `json:"setting_list"`
}

type StampInfoResult struct {
	OwningStampIds []int          `json:"owning_stamp_ids"`
	StampSetting   []StampSetting `json:"stamp_setting"`
}
