package entities

type Match struct {
	Id         int `json:"id"`
	HomeTeamId int `json:"home_team_id"`
	AwayTeamId int `json:"away_team_id"`
}
