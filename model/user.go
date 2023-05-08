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

// User ...
type User struct {
	UserID           int `json:"user_id"`
	UnitOwningUserID int `json:"unit_owning_user_id"`
}

// UserNaviRes ...
type UserNaviRes struct {
	User User `json:"user"`
}

// UserNaviResp ...
type UserNaviResp struct {
	Result     UserNaviRes `json:"result"`
	Status     int         `json:"status"`
	CommandNum bool        `json:"commandNum"`
	TimeStamp  int64       `json:"timeStamp"`
}

// NotificationResp ...
type NotificationResp struct {
	ResponseData []interface{} `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}
