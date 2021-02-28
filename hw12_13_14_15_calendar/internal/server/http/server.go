package internalhttp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/interfaces"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/gorilla/mux"
)

type Server struct {
	app    *app.App
	server *http.Server
	router *mux.Router
}

const (
	filterDay   = "day"
	filterWeek  = "week"
	filterMonth = "month"
)

func New(app *app.App, config configuration.HTTPConf) interfaces.HTTPApp {
	router := mux.NewRouter()

	address := net.JoinHostPort(config.Host, config.Port)

	server := &http.Server{ // nolint: exhaustivestruct
		Handler: router,
		Addr:    address,
	}

	return &Server{
		app:    app,
		server: server,
		router: router,
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.router.Use(s.loggingMiddleware)
	s.router.HandleFunc("/hello", s.HelloHandler).Methods(http.MethodGet)

	s.router.HandleFunc("/events", s.Create).Methods(http.MethodPost)
	s.router.HandleFunc("/events", s.Update).Methods(http.MethodPut)
	s.router.HandleFunc("/events", s.Delete).Methods(http.MethodDelete)
	s.router.HandleFunc("/events", s.Events).Methods(http.MethodGet)

	err := http.ListenAndServe(s.server.Addr, s.server.Handler)
	if err != nil {
		return fmt.Errorf("server start: %w", err)
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("server stop: %w", err)
	}
	return nil
}

func (s *Server) HelloHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "hello-world"})
	if err != nil {
		s.app.Logger.Error("error sending response")
	}
}

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	var event storage.Event
	var response Response

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		response.setError(err)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	err = s.app.Storage.Create(event)
	if err != nil {
		response.setError(err)
		response.JSON(w, http.StatusBadRequest)

		return
	}

	response.setData(map[string]string{"message": "Ok"})
	response.JSON(w, http.StatusOK)
}

func (s *Server) Update(w http.ResponseWriter, r *http.Request) {
	var event storage.Event
	var response Response

	params := mux.Vars(r)

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		response.setError(err)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	err = s.app.Storage.Update(params["id"], event)
	if err != nil {
		response.setError(err)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	response.setData(map[string]string{"message": "Ok"})
	response.JSON(w, http.StatusOK)
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	var response Response

	params := mux.Vars(r)

	err := s.app.Storage.Delete(params["id"])
	if err != nil {
		response.setError(err)
		respondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	response.setData(map[string]string{"message": "Ok"})
	response.JSON(w, http.StatusOK)
}

func (s *Server) Events(w http.ResponseWriter, r *http.Request) {
	var response Response
	var events []storage.Event

	params := mux.Vars(r)
	date, err := time.Parse(time.RFC3339, params["date"])
	if err != nil {
		date = time.Now()
	}

	v := r.URL.Query()
	f := v.Get("filter")

	switch f {
	case filterDay:
		events, err = s.app.Storage.DayEvents(date)
	case filterWeek:
		events, err = s.app.Storage.WeekEvents(date)
	case filterMonth:
		events, err = s.app.Storage.MonthEvents(date)
	default:
		err = errors.New("")
	}
	if err != nil {
		response.setError(err)
		response.JSON(w, http.StatusBadRequest)
		return
	}

	response.setData(events)
	response.JSON(w, http.StatusOK)
}
