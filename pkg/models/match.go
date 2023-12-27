package models

type Match struct {
	Id            int
	HomeTeamId    int
	HomeTeamName  string
	AwayTeamId    int
	AwayTeamName  string
	HomeTeamScore int
	AwayTeamScore int
}
