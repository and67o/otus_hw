package storage

import "time"

type Event struct {
	ID           string        `db:"id"`
	Title        string        `db:"title"`
	Date         time.Time     `db:"date"`
	Duration     time.Duration `db:"duration"`
	Description  string        `db:"description"`
	OwnerID      string        `db:"owner_id"`
	NotifyBefore time.Duration `db:"notify_before"`
}
