package model

// MuseumResp ...
type MuseumResp struct {
	ResponseData MuseumRes     `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

// MuseumParameter ...
type MuseumParameter struct {
	Smile int `json:"smile"`
	Pure  int `json:"pure"`
	Cool  int `json:"cool"`
}

// Museum ...
type Museum struct {
	Parameter      MuseumParameter `json:"parameter"`
	ContentsIDList []int           `json:"contents_id_list"`
}

// MuseumRes ...
type MuseumRes struct {
	MuseumInfo      Museum `json:"museum_info"`
	ServerTimestamp int64  `json:"server_timestamp"`
}

// MuseumInfoRes ...
type MuseumInfoRes struct {
	MuseumInfo Museum `json:"museum_info"`
}

// MuseumInfoResp ...
type MuseumInfoResp struct {
	Result     MuseumInfoRes `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}
