package interfaces

import (
	"context"
	"net/http"
	"time"

	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage interface {
	Get(id string) *storage.Event
	Create(e storage.Event) error
	Update(id string, e storage.Event) error
	Delete(id string) error
	DayEvents(time time.Time) ([]storage.Event, error)
	WeekEvents(time time.Time) ([]storage.Event, error)
	MonthEvents(time time.Time) ([]storage.Event, error)
}

type Logger interface {
	Info(msg string)
	Error(msg string)
	Fatal(msg string)
	Warn(msg string)
}

type GRPC interface {
	Stop() error
	Start(ctx context.Context) error
}

type HTTPApp interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	HelloHandler(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Events(w http.ResponseWriter, r *http.Request)
}
