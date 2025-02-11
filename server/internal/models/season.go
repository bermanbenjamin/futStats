package models

import (
	"github.com/google/uuid"
)

type Season struct {
	Base
	Year          string    `json:"year"`
	EndDate       string    `json:"end"`
	InitDate      string    `json:"init"`
	GoalsAmount   int       `json:"goals_amount"`
	AssistsAmount int       `json:"assists_amount"`
	StrikerId     uuid.UUID `json:"striker_id" gorm:"type:uuid"`
	Striker       Player    `json:"striker" gorm:"foreignKey:StrikerId;references:id"`
	WaiterId      uuid.UUID `json:"waiter_id" gorm:"type:uuid"`
	Waiter        Player    `json:"waiter" gorm:"foreignKey:WaiterId;references:id"`
	BestPlayerId  uuid.UUID `json:"best_player_id" gorm:"type:uuid"`
	BestPlayer    Player    `json:"best_player" gorm:"foreignKey:BestPlayerId;references:id"`
}
