package internalhttp

import (
	"context"
	"encoding/json"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/configuration"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	app    *app.App
	server *http.Server
	router *mux.Router
}

type Application interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	helloHandler(w http.ResponseWriter, r *http.Request)
}

func NewServer(app *app.App, config configuration.HTTPConf) *Server {
	router := mux.NewRouter()

	address := net.JoinHostPort(config.Host, config.Port)
	server := &http.Server{
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
	s.router.HandleFunc("/hello", s.helloHandler).Methods("GET")

	err := http.ListenAndServe(s.server.Addr, s.server.Handler)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "hello-world"})
	if err != nil {
		s.app.Logger.Error("error sending response")
	}
}
