package model

// GdprResp ...
type GdprResp struct {
	ResponseData GdprRes       `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

// GdprRes ...
type GdprRes struct {
	EnableGdpr      bool  `json:"enable_gdpr"`
	IsEea           bool  `json:"is_eea"`
	ServerTimestamp int64 `json:"server_timestamp"`
}
