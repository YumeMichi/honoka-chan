package model

type SifApi struct {
	Module    string `json:"module"`
	Action    string `json:"action"`
	Timestamp int64  `json:"timeStamp"`
}

type MuseumInfoParameter struct {
	Smile int `json:"smile"`
	Pure  int `json:"pure"`
	Cool  int `json:"cool"`
}

type MuseumInfo struct {
	Parameter      MuseumInfoParameter `json:"parameter"`
	ContentsIDList []int               `json:"contents_id_list"`
}

type MuseumInfoResult struct {
	MuseumInfo MuseumInfo `json:"museum_info"`
}

type MuseumInfoResp struct {
	Result     MuseumInfoResult `json:"result"`
	Status     int              `json:"status"`
	CommandNum bool             `json:"commandNum"`
	TimeStamp  int64            `json:"timeStamp"`
}
