package model

// Data ...
type Data struct {
	Version   int    `json:"version"`
	Team      []Team `json:"team"`
	Gemstock  any    `json:"gemstock"`
	Submember []any  `json:"submember"`
}

// Accessory ...
type Accessory struct {
	ID    string `json:"id"`
	Level int    `json:"level"`
}

// Team ...
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
