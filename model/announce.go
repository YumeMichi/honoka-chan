package model

// AnnounceResp ...
type AnnounceResp struct {
	ResponseData AnnounceRes `json:"response_data"`
	ReleaseInfo  []any       `json:"release_info"`
	StatusCode   int         `json:"status_code"`
}

// AnnounceRes ...
type AnnounceRes struct {
	HasUnreadAnnounce bool  `json:"has_unread_announce"`
	ServerTimestamp   int64 `json:"server_timestamp"`
}
