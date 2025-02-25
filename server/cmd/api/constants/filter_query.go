package constants

type QueryFilter string

const (
	QUERY_FILTER string = "x-api-field-type"
)

const (
	ID          QueryFilter = "id"
	EMAIL       QueryFilter = "email"
	PLAYERID    QueryFilter = "player_id"
	MATCHID     QueryFilter = "match_id"
	ASSISTENTID QueryFilter = "assistent_id"
	POSITION    QueryFilter = "position"
	EVENTID     QueryFilter = "event_id"
	SLUG        QueryFilter = "slug"
)
