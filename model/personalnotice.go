package model

// PersonalNoticeResp ...
type PersonalNoticeResp struct {
	ResponseData PersonalNoticeRes `json:"response_data"`
	ReleaseInfo  []any             `json:"release_info"`
	StatusCode   int               `json:"status_code"`
}

// PersonalNoticeRes ...
type PersonalNoticeRes struct {
	HasNotice       bool   `json:"has_notice"`
	NoticeID        int    `json:"notice_id"`
	Type            int    `json:"type"`
	Title           string `json:"title"`
	Contents        string `json:"contents"`
	ServerTimestamp int64  `json:"server_timestamp"`
}
