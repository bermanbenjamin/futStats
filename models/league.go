package models

import (
	"github.com/bermanbenjamin/futStats/commons"
	"github.com/google/uuid"
)

type League struct {
	commons.Base
	Name         string    `json:"name"`
	SeasonCount  int       `json:"season_count"`
	Seasons      []Season  `json:"season" gorm:"many2many:user_languages;"`
	BestPlayerId uuid.UUID `json:"best_player_id" gorm:"type:uuid"`
}
