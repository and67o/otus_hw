package sqlstorage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage"
	_ "github.com/go-sql-driver/mysql" // nolint: gci
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

const driverName = "mysql"
const format = "2006-01-02 15:04:05"

func New(config configuration.DBConf) (*Storage, error) {
	db, err := sqlx.Open(driverName, dataSourceName(config))
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("check ping db: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Get(id string) *storage.Event {
	_, err := s.db.Exec("SELECT * from events where id=?", id)
	if err != nil {
		return nil
	}
	return nil
}

func dataSourceName(config configuration.DBConf) string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.DBName,
	)
}

func (s *Storage) Close() error {
	err := s.db.Close()
	if err != nil {
		return fmt.Errorf("close connect: %w", err)
	}

	return nil
}

func (s *Storage) Create(e storage.Event) error {
	_, err := s.db.Exec("INSERT INTO events (id, title, `date`, duration, description, owner_id, notify_before) VALUES (?, ?, ?, ?, ?, ?, ?)",
		e.ID,
		e.Title,
		e.Date.Format(format),
		e.Duration,
		e.Description,
		e.OwnerID,
		e.NotifyBefore,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Update(e storage.Event) error {
	_, err := s.db.Exec("UPDATE events SET title=?, `date`=?, duration=?, description=?, owner_id=?, notify_before=? WHERE id=?",
		e.Title,
		e.Date.Format("2006-01-02 15:04:05"),
		e.Duration,
		e.Description,
		e.OwnerID,
		e.NotifyBefore,
		e.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Delete(id string) error {
	_, err := s.db.Exec("DELETE FROM events WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DayEvents(t time.Time) ([]storage.Event, error) {
	y, m, d := t.Date()

	res, err := s.db.Query("SELECT * FROM events WHERE YEAR(`date`) = ? AND MONTH(`date`) = ? and DAY(`date`) = ?", y, m, d)
	if res == nil {
		return nil, err
	}

	return handleResult(res)
}

func (s *Storage) WeekEvents(t time.Time) ([]storage.Event, error) {
	y, w := t.ISOWeek()

	res, err := s.db.Query("SELECT * FROM events WHERE YEAR(`date`) = ? AND WEEK(`date`) = ?", y, w)
	if res == nil {
		return nil, err
	}

	return handleResult(res)
}

func (s *Storage) MonthEvents(t time.Time) ([]storage.Event, error) {
	y, m, _ := t.Date()

	res, err := s.db.Query("SELECT * FROM events WHERE YEAR(`date`) = ? AND MONTH(`date`) = ?", y, m)
	if res == nil {
		return nil, err
	}

	return handleResult(res)
}

func handleResult(res *sql.Rows) ([]storage.Event, error) {
	events := make([]storage.Event, 0)

	for res.Next() {
		var e storage.Event
		var dateSQLRaw string
		var durationSQLRaw int64
		var notifyBeforeSQLRaw int64

		err := res.Scan(&e.ID, &e.Title, &dateSQLRaw, &durationSQLRaw, &e.Description, &e.OwnerID, &notifyBeforeSQLRaw)
		if err != nil {
			return nil, err
		}

		e.Date, err = time.Parse(format, dateSQLRaw)
		if err != nil {
			return nil, err
		}

		e.NotifyBefore = time.Duration(notifyBeforeSQLRaw)
		e.Duration = time.Duration(durationSQLRaw)

		events = append(events, e)
	}

	return events, nil
}
