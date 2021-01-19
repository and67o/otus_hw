package sqlstorage

import (
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/configuration"
	storage2 "github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var dbConfig = configuration.DBConf{
	User:   "admin",
	Pass:   "123",
	DBName: "go_api",
	Host:   "localhost",
	Port:   3306,
}

func TestStorage(t *testing.T) {
	t.Skip(true)
	t.Run("check for dates test", func(t *testing.T) {
		db, err := New(dbConfig)
		require.Nil(t, err)

		err = db.Create(storage2.Event{
			ID:           "123",
			Title:        "123",
			Date:         time.Date(2021, 1, 12, 12, 00, 00, 0, time.UTC),
			Duration:     0,
			Description:  "123",
			OwnerID:      "123",
			NotifyBefore: 0,
		})
		require.Nil(t, err)

		err = db.Create(storage2.Event{
			ID:           "1234",
			Title:        "123",
			Date:         time.Date(2021, 2, 12, 12, 00, 00, 0, time.UTC),
			Duration:     0,
			Description:  "123",
			OwnerID:      "123",
			NotifyBefore: 0,
		})
		require.Nil(t, err)

		err = db.Create(storage2.Event{
			ID:           "12345",
			Title:        "123",
			Date:         time.Date(2021, 1, 12, 12, 00, 00, 0, time.UTC),
			Duration:     0,
			Description:  "123",
			OwnerID:      "123",
			NotifyBefore: 0,
		})
		require.Nil(t, err)

		dateTime := time.Date(2021, 1, 12, 12, 00, 00, 0, time.UTC)
		res, err := db.DayEvents(dateTime)
		require.Equal(t, 2, len(res))

	})

	t.Run("CRUD test", func(t *testing.T) {
		t.Skip(true)

		db, err := New(dbConfig)
		require.Nil(t, err)

		id := uuid.New().String()
		var exampleEvent = storage2.Event{
			ID:           id,
			Title:        "CHECK",
			Date:         time.Now(),
			Duration:     0,
			Description:  "Description",
			OwnerID:      "123",
			NotifyBefore: 0,
		}

		err = db.Create(exampleEvent)
		require.Nil(t, err)

		exampleEvent.Title = exampleEvent.Title + "_1"
		exampleEvent.Date = time.Now()
		exampleEvent.Duration = 128
		exampleEvent.OwnerID = "567"

		err = db.Update(exampleEvent)
		require.Nil(t, err)

		err = db.Delete(id)
		require.Nil(t, err)
	})

}
