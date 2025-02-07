package models

type Player struct {
	Base
	Email       string   `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password    string   `json:"-" binding:"required" gorm:"not null"`
	Age         int      `json:"-" binding:"required"`
	Name        string   `json:"name" binding:"required" gorm:"not null"`
	Goals       int      `json:"goals" gorm:"deafult:0"`
	Assists     int      `json:"assists" gorm:"deafult:0"`
	Disarms     int      `json:"disarms" gorm:"deafult:0"`
	Dribbles    int      `json:"dribbles" gorm:"deafult:0"`
	Matches     int      `json:"matches" gorm:"deafult:0"`
	RedCards    int      `json:"red_cards" gorm:"deafult:0"`
	YellowCards int      `json:"yellow_cards" gorm:"deafult:0"`
	Position    string   `json:"position"`
	Leagues     []League `json:"leagues" gorm:"many2many:player_leagues"`
}
