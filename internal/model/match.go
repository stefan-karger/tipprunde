package model

import "time"

type Match struct {
	Contest  string
	GameDay  string `csv:"Game Day"`
	Date     *time.Time
	Time     string
	HomeTeam string `csv:"Home Team"`
	AwayTeam string `csv:"Away Team"`
	Score    string
}
