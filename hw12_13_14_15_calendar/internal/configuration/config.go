package configuration

import (
	"errors"
	"github.com/spf13/viper"
)

var (
	errEmptyPath = errors.New("path empty")
)

func New(path string) (Config, error) {
	var configuration Config

	if path == "" {
		return configuration, errEmptyPath
	}

	viper.SetConfigFile(path)

	err := viper.ReadInConfig()

	if err != nil {
		return configuration, err
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		return configuration, err
	}

	return configuration, nil
}
