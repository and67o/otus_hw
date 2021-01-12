package sqlstorage

import (
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStorage(t *testing.T) {
	config:= configuration.DBConf{
		User:   "admin",
		Pass:   "123",
		DBName: "go",
		Host:   "localhost",
	}
	_, err := New(config)
	require.Nil(t, err)
}
