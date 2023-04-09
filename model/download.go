package model

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

type AdditionalResult struct {
	Size int    `json:"size"`
	URL  string `json:"url"`
}

type AdditionalResp struct {
	ResponseData []AdditionalResult `json:"response_data"`
	ReleaseInfo  []interface{}      `json:"release_info"`
	StatusCode   int                `json:"status_code"`
}

type BatchReq struct {
	ClientVersion      string `json:"client_version"`
	Os                 string `json:"os"`
	PackageType        int    `json:"package_type"`
	ExcludedPackageIds []int  `json:"excluded_package_ids"`
	CommandNum         string `json:"commandNum"`
}

type BatchResult struct {
	Size int    `json:"size"`
	URL  string `json:"url"`
}

type BatchResp struct {
	ResponseData []BatchResult `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

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

type UpdateResult struct {
	Size    int    `json:"size"`
	URL     string `json:"url"`
	Version string `json:"version"`
}

type UpdateResp struct {
	ResponseData []UpdateResult `json:"response_data"`
	ReleaseInfo  []interface{}  `json:"release_info"`
	StatusCode   int            `json:"status_code"`
}

type EventResp struct {
	ResponseData []interface{} `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}
