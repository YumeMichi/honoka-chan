package model

// AdditionalReq ...
type AdditionalReq struct {
	Module      string `json:"module"`
	Mgd         int    `json:"mgd"`
	Action      string `json:"action"`
	TimeStamp   int    `json:"timeStamp"`
	PackageID   int    `json:"package_id"`
	TargetOs    string `json:"target_os"`
	PackageType int    `json:"package_type"`
	CommandNum  string `json:"commandNum"`
}

// AdditionalRes ...
type AdditionalRes struct {
	Size int    `json:"size"`
	URL  string `json:"url"`
}

// AdditionalResp ...
type AdditionalResp struct {
	ResponseData []AdditionalRes `json:"response_data"`
	ReleaseInfo  []interface{}   `json:"release_info"`
	StatusCode   int             `json:"status_code"`
}

// BatchReq ...
type BatchReq struct {
	ClientVersion      string `json:"client_version"`
	Os                 string `json:"os"`
	PackageType        int    `json:"package_type"`
	ExcludedPackageIds []int  `json:"excluded_package_ids"`
	CommandNum         string `json:"commandNum"`
}

// BatchRes ...
type BatchRes struct {
	Size int    `json:"size"`
	URL  string `json:"url"`
}

// BatchResp ...
type BatchResp struct {
	ResponseData []BatchRes    `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

// UpdateReq ...
type UpdateReq struct {
	Module          string        `json:"module"`
	TargetOs        string        `json:"target_os"`
	InstallVersion  string        `json:"install_version"`
	TimeStamp       int           `json:"timeStamp"`
	Action          string        `json:"action"`
	PackageList     []interface{} `json:"package_list"`
	CommandNum      string        `json:"commandNum"`
	ExternalVersion string        `json:"external_version"`
}

// UpdateRes ...
type UpdateRes struct {
	Size    int    `json:"size"`
	URL     string `json:"url"`
	Version string `json:"version"`
}

// UpdateResp ...
type UpdateResp struct {
	ResponseData []UpdateRes   `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

// UrlReq ...
type UrlReq struct {
	Module   string   `json:"module"`
	Os       string   `json:"os"`
	Mgd      int      `json:"mgd"`
	PathList []string `json:"path_list"`
	Action   string   `json:"action"`
}

// UrlRes ...
type UrlRes struct {
	UrlList []string `json:"url_list"`
}

// UrlResp ...
type UrlResp struct {
	ResponseData UrlRes        `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

// EventResp ...
type EventResp struct {
	ResponseData []interface{} `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}
