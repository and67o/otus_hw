package sqlstorage

import (
	"fmt"
	"time"

	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

const driverName = "pgx"
const driverFullName = "postgres"

func New(config configuration.DBConf) (*Storage, error) {
	db, err := sqlx.Open(driverName, dataSourceName(config))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
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
	return fmt.Sprintf("%s://%s:%s@%s/%s",
		driverFullName,
		config.User,
		config.Pass,
		config.Host,
		config.DBName,
	)
}

func (s *Storage) Close() error {
	err := s.db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Create(e storage.Event) error {
	_, err := s.db.Exec("INSERT INTO events SET id=?, title=?", e.ID, e.Title)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Update(e storage.Event) error {
	_, err := s.db.Exec("UPDATE events SET id=?, title=?", e.ID, e.Title)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Delete(id string) error {
	_, err := s.db.Exec("delete from events where id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DayEvents(time time.Time) []storage.Event {
	panic("implement me")
}

func (s *Storage) WeekEvents(time time.Time) []storage.Event {
	panic("implement me")
}

func (s *Storage) MonthEvents(time time.Time) []storage.Event {
	panic("implement me")
}
