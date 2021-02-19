package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/interfaces"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/logger"
	internalgrpc "github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/server/grpc"
	internalhttp "github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/server/http"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage/create"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config, err := configuration.New(configFile)
	if err != nil {
		log.Fatal(err)
	}

	logg, err := logger.New(config.Logger)
	if err != nil {
		log.Fatal(err)
	}

	storage, err := create.New(config)
	if err != nil {
		log.Fatal(err)
	}

	calendar := app.New(logg, storage)

	httpServer := internalhttp.New(calendar, config.Rest)
	grpcServer := internalgrpc.New(calendar, config.GRPC)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go watchSignals(httpServer, grpcServer, logg, cancel)

	logg.Info("calendar start")

	err = httpServer.Start(ctx)
	if err != nil {
		logg.Error("failed to start http server: " + err.Error())
		os.Exit(1) //nolint:gocritic
	}

	err = grpcServer.Start(ctx)
	if err != nil {
		logg.Error("failed to start grpc server: " + err.Error())
		os.Exit(1) //nolint:gocritic
	}
}

func watchSignals(httpServer interfaces.HTTPApp, grpcServer interfaces.GRPC, logg interfaces.Logger, cancel context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals
	signal.Stop(signals)
	cancel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := httpServer.Stop(ctx)
	if err != nil {
		logg.Error("failed to stop http server: " + err.Error())
	}

	err = grpcServer.Stop()
	if err != nil {
		logg.Error("failed to stop grpc server: " + err.Error())
	}
}
