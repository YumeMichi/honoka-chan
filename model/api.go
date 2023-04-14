package model

type SifApi struct {
	Module    string `json:"module"`
	Action    string `json:"action"`
	Timestamp int64  `json:"timeStamp"`
}
