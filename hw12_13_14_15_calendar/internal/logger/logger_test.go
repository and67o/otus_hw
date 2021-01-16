package logger

import (
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogger(t *testing.T) {
	config := configuration.LoggerConf{
		File:    "../../logs/log.log",
		Level: levelDebug,
	}

	_, err := New(config)
	require.Nil(t, err)
}
