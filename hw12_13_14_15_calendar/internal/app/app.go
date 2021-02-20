package app

import (
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/interfaces"
	pb "github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/server/grpc/pb"
)

type App struct { //nolint:maligned
	Logger  interfaces.Logger
	Storage interfaces.Storage

	uCS pb.UnimplementedCalendarServer
}

func New(logger interfaces.Logger, storage interfaces.Storage) *App {
	return &App{
		Logger:  logger,
		Storage: storage,
		uCS:     pb.UnimplementedCalendarServer{},
	}
}
