package models

import (
	"time"

	"github.com/bermanbenjamin/futStats/commons"
)

type Match struct {
	commons.Base
	Date  time.Time `json:"date"`
	Goals []Goal    `json:"goals" gorm:"goals"`
}
