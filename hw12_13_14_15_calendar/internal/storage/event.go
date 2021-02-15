package storage

import "time"

type Event struct {
	ID           string        `db:"id" json:"id"`
	Title        string        `db:"title" json:"title"`
	Date         time.Time     `db:"date" json:"date"`
	Duration     time.Duration `db:"duration" json:"duration"`
	Description  string        `db:"description" json:"description"`
	OwnerID      string        `db:"owner_id" json:"owner_id"`
	NotifyBefore time.Duration `db:"notify_before" json:"notify_before"`
}
