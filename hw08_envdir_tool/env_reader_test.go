package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type test struct {
	path    string
	expected Environment
	err      error
}

func TestReadDir(t *testing.T) {
	for _, tst := range [...]test{
		{
			path: "./testdata/env",
			expected: map[string]string{
				"BAR":   "bar",
				"FOO":   "   foo\nwith new line",
				"HELLO": `"hello"`,
				"UNSET": "",
			},
		},
	} {
		result, err := ReadDir(tst.path)
		require.Equal(t, tst.expected, result)
		require.Nil(t, tst.err, err)
	}
}
