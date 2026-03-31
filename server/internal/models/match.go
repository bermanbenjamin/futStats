package models

import (
	"time"

	"github.com/google/uuid"
)

type Match struct {
	Base
	LeagueId uuid.UUID `json:"league_id" gorm:"type:uuid;not null"`
	League   League    `json:"league" gorm:"foreignKey:LeagueId;references:id"`
	Date     time.Time `json:"date"`
	Events   []Event   `json:"events" gorm:"foreignKey:MatchId"`
}
