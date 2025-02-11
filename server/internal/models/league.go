package models

import "github.com/google/uuid"

type League struct {
	Base
	OwnerId uuid.UUID `json:"owner_id" gorm:"type:uuid;not null"`
	Owner   Player    `json:"owner" gorm:"foreignKey:OwnerId;references:id"`
	Name    string    `json:"name" gorm:"not null"`
	Slug    string    `json:"slug" gorm:"not null; unique"`
	Seasons []Season  `json:"seasons" gorm:"many2many:league_seasons;"`
	Members []Player  `json:"members" gorm:"many2many:league_members;constraint:OnDelete:CASCADE;"`
}

// LeagueMember represents the join table
type LeagueMember struct {
	LeagueID uuid.UUID `gorm:"type:uuid;primaryKey;"`
	PlayerID uuid.UUID `gorm:"type:uuid;primaryKey;"`
}
