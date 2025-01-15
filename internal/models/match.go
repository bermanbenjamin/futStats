package models

import (
	"time"
)

type Match struct {
	Base
	Date   time.Time `json:"date"`
	Events []Event   `json:"events" gorm:"goals"`
}
