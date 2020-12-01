package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCopy(t *testing.T) {
	err:= Copy("./testdata/input.txt","./testdata/input_1.txt", 0,10)
	require.NoError(t, err)
}
