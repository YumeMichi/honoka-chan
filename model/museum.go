package model

type MuseumInfoParameter struct {
	Smile int `json:"smile"`
	Pure  int `json:"pure"`
	Cool  int `json:"cool"`
}

type MuseumInfo struct {
	Parameter      MuseumInfoParameter `json:"parameter"`
	ContentsIDList []interface{}       `json:"contents_id_list"`
}

type MuseumInfoResult struct {
	MuseumInfo MuseumInfo `json:"museum_info"`
}
