package model

// LbDayItem ...
type LbDayItem struct {
	ItemID  int `json:"item_id"`
	AddType int `json:"add_type"`
	Amount  int `json:"amount"`
}

// LbDays ...
type LbDays struct {
	Day               int       `json:"day"`
	DayOfTheWeek      int       `json:"day_of_the_week"`
	SpecialDay        bool      `json:"special_day"`
	SpecialImageAsset string    `json:"special_image_asset"`
	Received          bool      `json:"received"`
	AdReceived        bool      `json:"ad_received"`
	Item              LbDayItem `json:"item"`
}

// LbMonth ...
type LbMonth struct {
	Year  int      `json:"year"`
	Month int      `json:"month"`
	Days  []LbDays `json:"days"`
}

// CalendarInfo ...
type CalendarInfo struct {
	CurrentDate  string  `json:"current_date"`
	CurrentMonth LbMonth `json:"current_month"`
	NextMonth    LbMonth `json:"next_month"`
}

// Reward ...
type Reward struct {
	ItemID  int `json:"item_id"`
	AddType int `json:"add_type"`
	Amount  int `json:"amount"`
}

// TotalLoginInfo ...
type TotalLoginInfo struct {
	LoginCount     int      `json:"login_count"`
	RemainingCount int      `json:"remaining_count"`
	Reward         []Reward `json:"reward"`
}

// LbRankInfo ...
type LbRankInfo struct {
	BeforeClassRankID int    `json:"before_class_rank_id"`
	AfterClassRankID  int    `json:"after_class_rank_id"`
	RankUpDate        string `json:"rank_up_date"`
}

// LbClassSystem ...
type LbClassSystem struct {
	RankInfo     LbRankInfo `json:"rank_info"`
	CompleteFlag bool       `json:"complete_flag"`
	IsOpened     bool       `json:"is_opened"`
	IsVisible    bool       `json:"is_visible"`
}

// LbRes ...
type LbRes struct {
	Sheets            []any          `json:"sheets"`
	CalendarInfo      CalendarInfo   `json:"calendar_info"`
	TotalLoginInfo    TotalLoginInfo `json:"total_login_info"`
	LicenseLbonusList []any          `json:"license_lbonus_list"`
	ClassSystem       LbClassSystem  `json:"class_system"`
	StartDashSheets   []any          `json:"start_dash_sheets"`
	EffortPoint       []EffortPoint  `json:"effort_point"`
	LimitedEffortBox  []any          `json:"limited_effort_box"`
	MuseumInfo        Museum         `json:"museum_info"`
	ServerTimestamp   int64          `json:"server_timestamp"`
	PresentCnt        int            `json:"present_cnt"`
}

// LbResp ...
type LbResp struct {
	ResponseData LbRes `json:"response_data"`
	ReleaseInfo  []any `json:"release_info"`
	StatusCode   int   `json:"status_code"`
}
