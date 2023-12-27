package entities

type Player struct {
	Id       int    `json:"id"`
	TeamId   int    `json:"team_id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}
