package model

type MuseumResp struct {
	ResponseData MuseumRes     `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

type MuseumParameter struct {
	Smile int `json:"smile"`
	Pure  int `json:"pure"`
	Cool  int `json:"cool"`
}

type Museum struct {
	Parameter      MuseumParameter `json:"parameter"`
	ContentsIDList []int           `json:"contents_id_list"`
}

type MuseumRes struct {
	MuseumInfo      Museum `json:"museum_info"`
	ServerTimestamp int64  `json:"server_timestamp"`
}
