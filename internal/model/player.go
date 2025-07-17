package model

import (
	"time"
)

type Player struct {
	Name         string
	Club         string
	Position     string
	Birthday     time.Time
	Height       string
	Foot         string
	JoinedAt     *time.Time `csv:"Joined At"` // nullable
	MarketValue  int        `csv:"Market Value"`
	InjuryStatus string     `csv:"Injury Status"`
}
