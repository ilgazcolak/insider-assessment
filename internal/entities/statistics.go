package entities

type Statistics struct {
	Id                 int    `json:"id"`
	PlayerId           int    `json:"player_id"`
	TeamName           string `json:"team_name"`
	Position           string `json:"position"`
	MatchId            int    `json:"match_id"`
	Score              int    `json:"score"`
	AttemptCount       int    `json:"attempt_count"`
	TwoPointsAttempt   int    `json:"two_points_attempt"`
	ThreePointsAttempt int    `json:"three_points_attempt"`
	TwoPointsSuccess   int    `json:"two_points_success"`
	ThreePointsSuccess int    `json:"three_points_success"`
	Assist             int    `json:"assist"`
}
