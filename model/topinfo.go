package model

type TopInfoLicenseInfo struct {
	LicenseList  []interface{} `json:"license_list"`
	LicensedInfo []interface{} `json:"licensed_info"`
	ExpiredInfo  []interface{} `json:"expired_info"`
	BadgeFlag    bool          `json:"badge_flag"`
}

type TopInfoResult struct {
	FriendActionCnt        int                `json:"friend_action_cnt"`
	FriendGreetCnt         int                `json:"friend_greet_cnt"`
	FriendVarietyCnt       int                `json:"friend_variety_cnt"`
	FriendNewCnt           int                `json:"friend_new_cnt"`
	PresentCnt             int                `json:"present_cnt"`
	SecretBoxBadgeFlag     bool               `json:"secret_box_badge_flag"`
	ServerDatetime         string             `json:"server_datetime"`
	ServerTimestamp        int64              `json:"server_timestamp"`
	NoticeFriendDatetime   string             `json:"notice_friend_datetime"`
	NoticeMailDatetime     string             `json:"notice_mail_datetime"`
	FriendsApprovalWaitCnt int                `json:"friends_approval_wait_cnt"`
	FriendsRequestCnt      int                `json:"friends_request_cnt"`
	IsTodayBirthday        bool               `json:"is_today_birthday"`
	LicenseInfo            TopInfoLicenseInfo `json:"license_info"`
	UsingBuffInfo          []interface{}      `json:"using_buff_info"`
	IsKlabIDTaskFlag       bool               `json:"is_klab_id_task_flag"`
	KlabIDTaskCanSync      bool               `json:"klab_id_task_can_sync"`
	HasUnreadAnnounce      bool               `json:"has_unread_announce"`
	ExchangeBadgeCnt       []int              `json:"exchange_badge_cnt"`
	AdFlag                 bool               `json:"ad_flag"`
	HasAdReward            bool               `json:"has_ad_reward"`
}

type TopInfoResp struct {
	Result     TopInfoResult `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}

type TopInfoOnceNotification struct {
	Push       bool `json:"push"`
	Lp         bool `json:"lp"`
	UpdateInfo bool `json:"update_info"`
	Campaign   bool `json:"campaign"`
	Live       bool `json:"live"`
	Lbonus     bool `json:"lbonus"`
	Event      bool `json:"event"`
	Secretbox  bool `json:"secretbox"`
	Birthday   bool `json:"birthday"`
}

type TopInfoOnceResult struct {
	NewAchievementCnt            int                     `json:"new_achievement_cnt"`
	UnaccomplishedAchievementCnt int                     `json:"unaccomplished_achievement_cnt"`
	LiveDailyRewardExist         bool                    `json:"live_daily_reward_exist"`
	TrainingEnergy               int                     `json:"training_energy"`
	TrainingEnergyMax            int                     `json:"training_energy_max"`
	Notification                 TopInfoOnceNotification `json:"notification"`
	OpenArena                    bool                    `json:"open_arena"`
	CostumeStatus                bool                    `json:"costume_status"`
	OpenAccessory                bool                    `json:"open_accessory"`
	ArenaSiSkillUniqueCheck      bool                    `json:"arena_si_skill_unique_check"`
	OpenV98                      bool                    `json:"open_v98"`
}

type TopInfoOnceResp struct {
	Result     TopInfoOnceResult `json:"result"`
	Status     int               `json:"status"`
	CommandNum bool              `json:"commandNum"`
	TimeStamp  int64             `json:"timeStamp"`
}
