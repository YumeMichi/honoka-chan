package model

type SifApi struct {
	Module    string `json:"module"`
	Action    string `json:"action"`
	Timestamp int64  `json:"timeStamp"`
}

type AuthKeyReq struct {
	DummyToken string `json:"dummy_token"`
	AuthData   string `json:"auth_data"`
}

type AuthKeyResp struct {
	ResponseData struct {
		AuthorizeToken string `json:"authorize_token"`
		DummyToken     string `json:"dummy_token"`
	} `json:"response_data"`
	ReleaseInfo [0]interface{} `json:"release_info"`
	StatusCode  int            `json:"status_code"`
}

type LoginReq struct {
	LoginKey    string `json:"login_key"`
	LoginPasswd string `json:"login_passwd"`
	DevToken    string `json:"devtoken"`
}

type LoginResp struct {
	ResponseData struct {
		AuthorizeToken  string `json:"authorize_token"`
		UserId          int    `json:"user_id"`
		ReviewVersion   string `json:"review_version"`
		ServerTimestamp int64  `json:"server_timestamp"`
		IdfaEnabled     bool   `json:"idfa_enabled"`
		SkipLoginNews   bool   `json:"skip_login_news"`
		AdultFlag       int    `json:"adult_flag"`
	} `json:"response_data"`
	ReleaseInfo [0]interface{} `json:"release_info"`
	StatusCode  int            `json:"status_code"`
}
