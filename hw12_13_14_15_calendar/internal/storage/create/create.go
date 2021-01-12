package create

import (
	"errors"
	"time"

	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage/sql"
)

const (
	sqlType    = "sql"
	memoryType = "memory"
)

var (
	ErrUnknownType = errors.New("unknown type")
)

type Storage interface {
	Get(id string) *storage.Event
	Create(e storage.Event) error
	Update(e storage.Event) error
	Delete(id string) error
	DayEvents(time time.Time) []storage.Event
	WeekEvents(time time.Time) []storage.Event
	MonthEvents(time time.Time) []storage.Event
}

func New(config configuration.Config) (Storage, error) {
	switch config.Memory.Type {
	case sqlType:
		return sqlstorage.New(config.DB)
	case memoryType:
		return memorystorage.New(), nil
	default:
		return nil, ErrUnknownType
	}
}
