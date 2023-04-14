package model

// NoticeFriendVarietyResp ...
type NoticeFriendVarietyResp struct {
	ResponseData NoticeFriendVarietyRes `json:"response_data"`
	ReleaseInfo  []interface{}          `json:"release_info"`
	StatusCode   int                    `json:"status_code"`
}

// NoticeFriendVarietyRes ...
type NoticeFriendVarietyRes struct {
	ItemCount       int           `json:"item_count"`
	NoticeList      []interface{} `json:"notice_list"`
	ServerTimestamp int64         `json:"server_timestamp"`
}

// NoticeFriendGreetingResp ...
type NoticeFriendGreetingResp struct {
	ResponseData NoticeFriendGreetingRes `json:"response_data"`
	ReleaseInfo  []interface{}           `json:"release_info"`
	StatusCode   int                     `json:"status_code"`
}

// NoticeFriendGreetingRes ...
type NoticeFriendGreetingRes struct {
	NextId          int           `json:"next_id"`
	NoticeList      []interface{} `json:"notice_list"`
	ServerTimestamp int64         `json:"server_timestamp"`
}

// NoticeUserGreetingResp ...
type NoticeUserGreetingResp struct {
	ResponseData NoticeUserGreetingRes `json:"response_data"`
	ReleaseInfo  []interface{}         `json:"release_info"`
	StatusCode   int                   `json:"status_code"`
}

// NoticeUserGreetingRes ...
type NoticeUserGreetingRes struct {
	ItemCount       int           `json:"item_count"`
	HasNext         bool          `json:"has_next"`
	NoticeList      []interface{} `json:"notice_list"`
	ServerTimestamp int64         `json:"server_timestamp"`
}
