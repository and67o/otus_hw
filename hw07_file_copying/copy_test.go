package main

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

type test struct {
	from        string
	to          string
	offset      int64
	limit       int64
	err         error
	checkResult string
}

func TestCopyOk(t *testing.T) {
	defer os.Remove("out.txt")
	for _, tst := range [...]test{
		{
			from:        "testdata/input.txt",
			to:          "out.txt",
			checkResult: "testdata/out_offset0_limit0.txt",
		},
		{
			from:        "testdata/input.txt",
			to:          "out.txt",
			limit:       10,
			checkResult: "testdata/out_offset0_limit10.txt",
		},
		{
			from:        "testdata/input.txt",
			to:          "out.txt",
			limit:       1000,
			checkResult: "testdata/out_offset0_limit1000.txt",
		},
		{
			from:        "testdata/input.txt",
			to:          "out.txt",
			limit:       10000,
			checkResult: "testdata/out_offset0_limit10000.txt",
		},
		{
			from:        "testdata/input.txt",
			to:          "out.txt",
			limit:       1000,
			offset:      100,
			checkResult: "testdata/out_offset100_limit1000.txt",
		},
		{
			from:        "testdata/input.txt",
			to:          "out.txt",
			limit:       1000,
			offset:      6000,
			checkResult: "testdata/out_offset6000_limit1000.txt",
		},
	} {
		t.Run("ok", func(t *testing.T) {
			err := Copy(tst.from, tst.to, tst.offset, tst.limit)
			require.Nil(t,err)

			filesEqual, err := compareFiles(tst.to, tst.checkResult)
			require.Nil(t, err)

			require.True(t, filesEqual)
		})
	}
}

func TestCopyFail(t *testing.T) {
	for _, tst := range [...]test{
		{
			from: "no_file.txt",
			err:  ErrNotFoundFile,
		},
		{
			from: "testdata",
			err:  ErrUnsupportedFile,
		},
		{
			from: "testdata/input.txt",
			offset: 100000,
			err:  ErrOffsetExceedsFileSize,
		},
	} {
		t.Run("fail", func(t *testing.T) {
			err := Copy(tst.from, tst.to, tst.offset, tst.limit)
			require.Equal(t,err, tst.err)
		})
	}
}

func compareFiles(originalFilePath string, copyFilePath string) (bool, error) {
	originalFile, err := ioutil.ReadFile(originalFilePath)
	if err != nil {
		return false, err
	}

	copyFile, err := ioutil.ReadFile(copyFilePath)
	if err != nil {
		return false, err
	}

	return bytes.Equal(originalFile, copyFile), nil
}
