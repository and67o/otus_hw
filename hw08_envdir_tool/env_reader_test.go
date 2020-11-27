package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

type test struct {
	path    string
	expected Environment
	err      error
}

func TestReadDir(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
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
	})

	t.Run("folder not found", func(t *testing.T) {
		for _, tst := range [...]test{
			{
				path: "./testdata1/env",
				expected: map[string]string{},
			},
		} {
			result, err := ReadDir(tst.path)
			require.Equal(t, tst.expected, result)
			require.NotNil(t, err)
		}
	})

	t.Run("ok with = and new file", func(t *testing.T) {
		//new file
		f, _:= os.Create("./testdata/env/CHECK")
		defer func() {
			os.Remove("./testdata/env/CHECK")
		}()
		_, _ = f.Write([]byte("check"))

		//new file with =
		file, _:= os.Create("./testdata/env/CHECK=")
		defer func() {
			os.Remove("./testdata/env/CHECK=")
		}()
		_, _ = file.Write([]byte("check"))

		for _, tst := range [...]test{
			{
				path: "./testdata/env",
				expected: map[string]string{
					"BAR":   "bar",
					"FOO":   "   foo\nwith new line",
					"HELLO": `"hello"`,
					"UNSET": "",
					"CHECK": "check",
				},
			},
		} {
			result, err := ReadDir(tst.path)
			require.Equal(t, tst.expected, result)
			require.Nil(t, tst.err, err)
		}

	})
}

