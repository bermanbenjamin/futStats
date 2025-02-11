package models

type Player struct {
	Base
	Email           string   `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password        string   `json:"-" binding:"required" gorm:"not null"`
	Age             int      `json:"age" binding:"required"`
	Name            string   `json:"name" binding:"required" gorm:"not null"`
	Goals           int      `json:"goals" gorm:"default:0"`
	Assists         int      `json:"assists" gorm:"default:0"`
	Disarms         int      `json:"disarms" gorm:"default:0"`
	Dribbles        int      `json:"dribbles" gorm:"default:0"`
	Matches         int      `json:"matches" gorm:"default:0"`
	RedCards        int      `json:"red_cards" gorm:"default:0"`
	YellowCards     int      `json:"yellow_cards" gorm:"default:0"`
	Position        string   `json:"position"`
	OwnedLeagues    []League `json:"owned_leagues" gorm:"foreignKey:OwnerId;references:id"`
	MemberOfLeagues []League `json:"member_of_leagues" gorm:"many2many:league_members;constraint:OnDelete:CASCADE;"`
}
