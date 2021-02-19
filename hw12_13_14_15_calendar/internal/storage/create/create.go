package create

import (
	"errors"

	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/interfaces"
	memorystorage "github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage/sql"
)

const (
	sqlType     configuration.MemoryType = "sql"
	storageType configuration.MemoryType = "memory"
)

var ErrUnknownType = errors.New("unknown type")

func New(config configuration.Config) (interfaces.Storage, error) {
	switch config.Memory.Type {
	case sqlType:
		return sqlstorage.New(config.DB)
	case storageType:
		return memorystorage.New(), nil
	default:
		return nil, ErrUnknownType
	}
}
