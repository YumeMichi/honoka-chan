package model

// UserNaviResp ...
type UserNaviChangeResp struct {
	ResponseData []interface{} `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

// UserNameChangeResp ...
type UserNameChangeResp struct {
	ResponseData UserNameChangeRes `json:"response_data"`
	ReleaseInfo  []interface{}     `json:"release_info"`
	StatusCode   int               `json:"status_code"`
}

// UserNameChangeRes ...
type UserNameChangeRes struct {
	BeforeName      string `json:"before_name"`
	AfterName       string `json:"after_name"`
	ServerTimestamp int64  `json:"server_timestamp"`
}
