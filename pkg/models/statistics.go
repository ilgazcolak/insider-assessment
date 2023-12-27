package models

type Statistics struct {
	Id                 int     `json:"id"`
	PlayerId           int     `json:"player_id"`
	PlayerName         string  `json:"player_name"`
	TeamName           string  `json:"team_name"`
	Position           string  `json:"position"`
	Score              int     `json:"score"`
	MatchId            int     `json:"match_id"`
	AttemptCount       int     `json:"attempt_count"`
	TwoPointsAttempt   int     `json:"two_points_attempt"`
	ThreePointsAttempt int     `json:"three_points_attempt"`
	TwoPointsSuccess   int     `json:"two_points_success"`
	ThreePointsSuccess int     `json:"three_points_success"`
	TwoPointsRate      float64 `json:"two_points_rate"`
	ThreePointsRate    float64 `json:"three_points_rate"`
	Assist             int     `json:"assist"`
}
