package main

import (
	"context"
	"flag"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/app"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/server/http"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/storage/create"
	"log"
	"os"
	"os/signal"
	"time"
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

	server := internalhttp.NewServer(calendar, config.Server)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go watchSignals(server, logg, cancel)

	logg.Info("calendar is running...")

	err = server.Start(ctx)
	if err != nil {
		logg.Error("failed to start http server: " + err.Error())
		os.Exit(1)
	}
}

func watchSignals(server *internalhttp.Server, logg *logger.Logger, cancel context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals)

	<-signals
	signal.Stop(signals)
	cancel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		logg.Error("failed to stop http server: " + err.Error())
	}
}
