package models

import "time"

// Demo models
type TestModelDemo struct {
	ID        uint64    `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	CreatedAt time.Time `json:"createdAt"`
}
