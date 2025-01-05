package models

type Player struct {
	Base
	Name        string `json:"name"`
	Goals       int    `json:"goals"`
	Assists     int    `json:"assists"`
	Disarms     int    `json:"disarms"`
	Dribbles    int    `json:"dribbles"`
	Matches     int    `json:"matches"`
	RedCards    int    `json:"red_cards"`
	YellowCards int    `json:"yellow_cards"`
	Position    string `json:"position"`
}
