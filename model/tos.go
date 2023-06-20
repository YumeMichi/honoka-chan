package model

// TosResp ...
type TosResp struct {
	ResponseData TosRes `json:"response_data"`
	ReleaseInfo  []any  `json:"release_info"`
	StatusCode   int    `json:"status_code"`
}

// TosResp ...
type TosRes struct {
	TosID           int   `json:"tos_id"`
	TosType         int   `json:"tos_type"`
	IsAgreed        bool  `json:"is_agreed"`
	ServerTimestamp int64 `json:"server_timestamp"`
}
