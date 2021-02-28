package memorystorage

import (
	"errors"
	"sync"
	"time"

	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/interfaces"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	mu     sync.RWMutex
	events map[string]*storage.Event
}

var (
	errExist    = errors.New("event exist")
	errNotFound = errors.New("not found")
)

func New() interfaces.Storage {
	return &Storage{
		events: make(map[string]*storage.Event),
	}
}

func (s *Storage) Get(id string) *storage.Event {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.events[id]
	if !ok {
		return nil
	}

	return s.events[id]
}

func (s *Storage) Create(e storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.events[e.ID]
	if ok {
		return errExist
	}

	s.events[e.ID] = &e

	return nil
}

func (s *Storage) Update(id string, e storage.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.events[id] == nil {
		return errNotFound
	}

	s.events[id] = &e

	return nil
}

func (s *Storage) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.events[id] == nil {
		return errNotFound
	}

	delete(s.events, id)

	return nil
}

func (s *Storage) DayEvents(time time.Time) ([]storage.Event, error) {
	var events []storage.Event
	y, m, d := time.Date()

	for _, event := range s.events {
		eventY, eventM, eventD := event.Date.Date()

		if y == eventY && m == eventM && d == eventD {
			events = append(events, *event)
		}
	}

	return events, nil
}

func (s *Storage) WeekEvents(time time.Time) ([]storage.Event, error) {
	var events []storage.Event
	y, w := time.ISOWeek()

	for _, event := range s.events {
		eventY, eventW := event.Date.ISOWeek()

		if y == eventY && w == eventW {
			events = append(events, *event)
		}
	}

	return events, nil
}

func (s *Storage) MonthEvents(time time.Time) ([]storage.Event, error) {
	var events []storage.Event
	y, m, _ := time.Date()

	for _, event := range s.events {
		eventY, eventM, _ := event.Date.Date()

		if y == eventY && m == eventM {
			events = append(events, *event)
		}
	}

	return events, nil
}
