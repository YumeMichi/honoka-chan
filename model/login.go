package model

// LoginRes ...
type LoginRes struct {
	AuthorizeToken  string `json:"authorize_token"`
	UserId          int    `json:"user_id"`
	ReviewVersion   string `json:"review_version"`
	ServerTimestamp int64  `json:"server_timestamp"`
	IdfaEnabled     bool   `json:"idfa_enabled"`
	SkipLoginNews   bool   `json:"skip_login_news"`
	AdultFlag       int    `json:"adult_flag"`
}

// LoginResp ...
type LoginResp struct {
	ResponseData LoginRes      `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}
