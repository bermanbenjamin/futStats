package model

import (
	"github.com/google/uuid"
)

type Event struct {
	Base
	Type        string    `json:"type"`
	PlayerId    uuid.UUID `json:"player_id" gorm:"type:uuid; index"`
	Player      Player    `json:"player" gorm:"foreignKey:PlayerId;references:ID"`
	MatchId     uuid.UUID `json:"match_id" gorm:"type:uuid"`
	Match       Match     `json:"match" gorm:"foreignKey:MatchId;references:ID"`
	AssistentId uuid.UUID `json:"assistent_id" gorm:"type:uuid"`
	Assistent   Player    `json:"assistent" gorm:"foreignKey:AssistentId;references:ID"`
}
