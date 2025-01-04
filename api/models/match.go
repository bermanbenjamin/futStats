package models

import (
	"time"

	"github.com/bermanbenjamin/futStats/api/commons"
)

type Match struct {
	commons.Base
	Date   time.Time `json:"date"`
	Events []Event   `json:"events" gorm:"goals"`
}
