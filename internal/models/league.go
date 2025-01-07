package model

type League struct {
	Base
	Name        string   `json:"name"`
	SeasonCount int      `json:"season_count"`
	Seasons     []Season `json:"season" gorm:"many2many:user_languages;"`
}
