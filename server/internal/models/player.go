package models

type Player struct {
	Base
	Email           string   `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password        string   `json:"-" binding:"required" gorm:"not null"`
	Age             int      `json:"age" binding:"required"`
	Name            string   `json:"name" binding:"required" gorm:"not null"`
	Position        string   `json:"position"`
	OwnedLeagues    []League `json:"owned_leagues" gorm:"foreignKey:OwnerId;references:id"`
	MemberOfLeagues []League `json:"member_of_leagues" gorm:"many2many:league_members;constraint:OnDelete:CASCADE;"`

	// Computed stats - populated by queries, not stored in DB
	Goals       int `json:"goals" gorm:"-"`
	Assists     int `json:"assists" gorm:"-"`
	Disarms     int `json:"disarms" gorm:"-"`
	Dribbles    int `json:"dribbles" gorm:"-"`
	Matches     int `json:"matches" gorm:"-"`
	RedCards    int `json:"red_cards" gorm:"-"`
	YellowCards int `json:"yellow_cards" gorm:"-"`
}
