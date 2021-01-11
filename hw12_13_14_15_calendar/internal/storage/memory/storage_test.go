package memorystorage

import (
	storage2 "github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestStorage(t *testing.T) {
	storage := New()

	e := storage2.Event{
		ID:           "1",
		Title:        "123",
		Date:         time.Time{},
		Duration:     0,
		Description:  "",
		OwnerId:      "",
		NotifyBefore: 0,
	}

	err := storage.Create(e)
	require.Nil(t, err)

	err = storage.Create(e)
	require.Error(t, errExist)

	event := storage.Get("1")
	require.Equal(t, "1", event.ID)

	e.Title = "456"
	err = storage.Update(e)
	require.Nil(t, err)

	eventNotFoud := storage.Get("2")
	require.Nil(t, eventNotFoud)

	eventAfterUpdate := storage.Get("1")
	require.Equal(t, "456", eventAfterUpdate.Title)

	err = storage.Delete("1")
	require.Nil(t, err)

	eventAfterDelete := storage.Get("1")
	require.Nil(t, eventAfterDelete)
}
