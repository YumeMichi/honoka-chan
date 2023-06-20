package model

// AuthKeyRes ...
type AuthKeyRes struct {
	AuthorizeToken string `json:"authorize_token"`
	DummyToken     string `json:"dummy_token"`
}

// AuthKeyResp ...
type AuthKeyResp struct {
	ResponseData AuthKeyRes `json:"response_data"`
	ReleaseInfo  []any      `json:"release_info"`
	StatusCode   int        `json:"status_code"`
}
