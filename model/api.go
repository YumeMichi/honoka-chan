package model

import "encoding/json"

// ApiReq ...
type ApiReq struct {
	Module    string `json:"module"`
	Action    string `json:"action"`
	Timestamp int64  `json:"timeStamp"`
}

// ApiResp ...
type ApiResp struct {
	ResponseData json.RawMessage `json:"response_data"`
	ReleaseInfo  []any           `json:"release_info"`
	StatusCode   int             `json:"status_code"`
}
