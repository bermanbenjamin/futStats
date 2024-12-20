package models

import (
	"time"
)

type Player struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Goals     int
	Assists   int
	Matches   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
