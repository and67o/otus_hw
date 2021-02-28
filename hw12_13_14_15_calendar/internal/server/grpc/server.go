//go:generate protoc -I ../../../api/ EventService.proto --go_out=pb --go-grpc_out=pb

package internalgrpc

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/interfaces"
	pb "github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/server/grpc/pb"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

const network = "tcp"

type Server struct { //nolint:maligned
	addr   string
	app    *app.App
	server *grpc.Server
	uCS    pb.UnimplementedCalendarServer
}

func New(app *app.App, config configuration.GRPCConf) interfaces.GRPC {
	return &Server{
		app:    app,
		addr:   net.JoinHostPort(config.Host, config.Port),
		server: nil,
		uCS:    pb.UnimplementedCalendarServer{},
	}
}

func (s *Server) Stop() error {
	if s.server == nil {
		return errors.New("grpc server is nil")
	}

	s.server.GracefulStop()
	s.app.Logger.Info("server stoped")

	return nil
}

func (s *Server) Start(ctx context.Context) error {
	l, err := net.Listen(network, s.addr)
	if err == nil {
		return fmt.Errorf("start server: %w", err)
	}

	serverGRPC := grpc.NewServer()
	s.server = serverGRPC
	pb.RegisterCalendarServer(serverGRPC, s.uCS)

	err = serverGRPC.Serve(l)
	if err != nil {
		return errors.New("error start server")
	}

	<-ctx.Done()

	return nil
}

func (s *Server) Create(_ context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	event := grpcEventToEvent(request.Event)

	err := s.app.Storage.Create(event)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return &pb.CreateResponse{}, nil
}

func (s *Server) Delete(_ context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	id := strconv.FormatInt(request.Id, 10)

	err := s.app.Storage.Delete(id)
	if err != nil {
		return nil, fmt.Errorf("delete: %w", err)
	}

	return &pb.DeleteResponse{}, nil
}
func (s *Server) Update(_ context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	id := strconv.FormatInt(request.Id, 10)
	event := grpcEventToEvent(request.Event)

	err := s.app.Storage.Update(id, event)
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}

	return &pb.UpdateResponse{}, nil
}

func (s *Server) DayEvents(_ context.Context, request *pb.DayEventsRequest) (*pb.DayEventsResponse, error) {
	date, err := ptypes.Timestamp(request.Date)
	if err != nil {
		return nil, fmt.Errorf("timestamp error: %w", err)
	}

	events, err := s.app.Storage.DayEvents(date)
	if err != nil {
		return nil, fmt.Errorf("dayevents error: %w", err)
	}

	dayEvents := make([]*pb.Event, len(events))
	for _, event := range events {
		dayEvent, err := eventToGrpcEvent(event)
		if err != nil {
			return nil, fmt.Errorf("parse event: %w", err)
		}

		dayEvents = append(dayEvents, dayEvent)
	}

	return &pb.DayEventsResponse{Events: dayEvents}, nil
}
func (s *Server) WeekEvents(_ context.Context, request *pb.WeekEventsRequest) (*pb.WeekEventsResponse, error) {
	date, err := ptypes.Timestamp(request.Date)
	if err != nil {
		return nil, fmt.Errorf("weekevents error: %w", err)
	}

	events, err := s.app.Storage.DayEvents(date)
	if err != nil {
		return nil, fmt.Errorf("weekevents error: %w", err)
	}

	weekEvents := make([]*pb.Event, len(events))
	for _, event := range events {
		weekEvent, err := eventToGrpcEvent(event)
		if err != nil {
			return nil, err
		}

		weekEvents = append(weekEvents, weekEvent)
	}

	return &pb.WeekEventsResponse{Events: weekEvents}, nil
}
func (s *Server) MonthEvents(_ context.Context, request *pb.MonthEventsRequest) (*pb.MonthEventsResponse, error) {
	date, err := ptypes.Timestamp(request.Date)
	if err != nil {
		return nil, fmt.Errorf("timestamp error: %w", err)
	}

	events, err := s.app.Storage.DayEvents(date)
	if err != nil {
		return nil, fmt.Errorf("dayevents error: %w", err)
	}

	monthEvents := make([]*pb.Event, len(events))
	for _, event := range events {
		monthEvent, err := eventToGrpcEvent(event)
		if err != nil {
			return nil, fmt.Errorf("parse event: %w", err)
		}

		monthEvents = append(monthEvents, monthEvent)
	}

	return &pb.MonthEventsResponse{Events: monthEvents}, nil
}

func grpcEventToEvent(e *pb.Event) storage.Event {
	return storage.Event{
		ID:           e.Id,
		Title:        e.Title,
		Date:         e.Date.AsTime(),
		Duration:     e.Duration.AsDuration(),
		Description:  e.Description,
		OwnerID:      e.OwnerId,
		NotifyBefore: e.NotifyBefore.AsDuration(),
	}
}

func eventToGrpcEvent(e storage.Event) (*pb.Event, error) {
	date, err := ptypes.TimestampProto(e.Date)
	if err != nil {
		return nil, fmt.Errorf("timestamp error: %w", err)
	}

	return &pb.Event{
		Id:           e.ID,
		Title:        e.Title,
		Date:         date,
		Duration:     ptypes.DurationProto(e.Duration),
		Description:  e.Description,
		OwnerId:      e.OwnerID,
		NotifyBefore: ptypes.DurationProto(e.NotifyBefore),
	}, nil
}
