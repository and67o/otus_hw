package logger

import (
	"fmt"
	"github.com/and67o/otus_hw/hw12_13_14_15_calendar/internal/configuration"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

type Interface interface {
	Info(msg string)
	Error(msg string)
	Fatal(msg string)
	Warn(msg string)
}

const outPut = "stderr"

func New(configuration configuration.LoggerConf) (*Logger, error) {
	var logger = new(Logger)

	config := getConfig(configuration.IsProd)

	config.DisableStacktrace = configuration.TraceOn

	config.OutputPaths = []string{outPut}
	config.OutputPaths = []string{configuration.File}

	l, err := config.Build()
	if err != nil {
		return logger, fmt.Errorf("config build error: %w", err)
	}

	logger.logger = l
	return logger, nil
}

func getConfig(isProd bool) zap.Config {
	var config zap.Config

	if isProd {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	return config
}

func (l Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l Logger) Error(msg string) {
	l.logger.Error(msg)
}

func (l Logger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l Logger) Warn(msg string) {
	l.logger.Warn(msg)
}
