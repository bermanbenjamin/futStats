package model

import "github.com/google/uuid"

type League struct {
	Base
	OwnerId uuid.UUID `json:"owner_id" gorm:"type:uuid; index"`
	Owner   Player    `json:"owner" gorm:"foreignKey:OwnerId;references:ID"`
	Name    string    `json:"name" gorm:"not null"`
	Slug    string    `json:"slug" gorm:"not null; unique"`
	Seasons []Season  `json:"season" gorm:"many2many:league_seasons;"`
	Players []Player  `json:"players" gorm:"many2many:league_players;"`
}
