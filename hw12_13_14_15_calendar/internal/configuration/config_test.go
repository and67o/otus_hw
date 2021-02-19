package configuration

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

type test struct {
	path string
	err  error
}

func TestConfigurationErrors(t *testing.T) {
	for _, tst := range [...]test{
		{
			path: "",
			err:  errors.New("path empty"),
		},
	} {
		_, err := New(tst.path)
		require.Equal(t, tst.err, err)
	}
}

func TestConfiguration(t *testing.T) {
	for _, tst := range [...]test{
		{
			path: "../../configs/config.toml",
		},
	} {
		_, err := New(tst.path)
		require.Nil(t, err)
	}
}
