package model

type Data struct {
	Version   int           `json:"version"`
	Team      []Team        `json:"team"`
	Gemstock  interface{}   `json:"gemstock"`
	Submember []interface{} `json:"submember"`
}

type Accessory struct {
	ID    string `json:"id"`
	Level int    `json:"level"`
}

type Team struct {
	Hp         int       `json:"hp"`
	Smile      int       `json:"smile"`
	Pure       int       `json:"pure"`
	Cool       int       `json:"cool"`
	Skilllevel int       `json:"skilllevel"`
	Gemlist    []string  `json:"gemlist"`
	Cardid     int       `json:"cardid"`
	Mezame     int       `json:"mezame"`
	Maxcost    int       `json:"maxcost"`
	Accessory  Accessory `json:"accessory,omitempty"`
}
