package model

import "encoding/json"

// response_data
type Response struct {
	ResponseData json.RawMessage `json:"response_data"`
	ReleaseInfo  []interface{}   `json:"release_info"`
	StatusCode   int             `json:"status_code"`
}
