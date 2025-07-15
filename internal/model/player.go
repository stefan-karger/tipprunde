package model

import (
	"time"
)

// Player represents a football player's profile data.
type Player struct {
	Name         string     `csv:"Name"`
	Position     string     `csv:"Position"`
	Birthday     time.Time  `csv:"Birthday"` // Use time.Time for dates
	Height       string     `csv:"Height"`   // Keep as string (e.g., "1,89m")
	Foot         string     `csv:"Foot"`
	JoinedAt     *time.Time `csv:"Joined At"`     // Use pointer to time.Time for nullable date
	MarketValue  int        `csv:"Market Value"`  // Store as integer in Euros
	InjuryStatus string     `csv:"Injury Status"` // Title of the injury span
}
