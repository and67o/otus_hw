package storage

import "time"

type Event struct {
	ID           string
	Title        string
	Date         time.Time
	Duration     time.Duration
	Description  string
	OwnerID      string
	NotifyBefore time.Duration
}
