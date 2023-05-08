package model

// Msg ...
type Msg struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Redirect string `json:"redirect"`
}
