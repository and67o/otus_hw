package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrNotFoundFile          = errors.New("file not found")
	ErrWrongStat             = errors.New("stat of file is wrong")
	ErrSeek                  = errors.New("wrong start of file")
	ErrCreateFile            = errors.New("error in create file")
	ErrCopy                  = errors.New("error in copy file")
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return ErrNotFoundFile
	}
	defer fromFile.Close()

	stat, err := fromFile.Stat()
	if err != nil {
		return ErrWrongStat
	}

	if stat.IsDir() || stat.Size() == 0 {
		return ErrUnsupportedFile
	}
	if offset > stat.Size() {
		return ErrOffsetExceedsFileSize
	}

	howMuchToCopy := stat.Size() - offset
	if limit > 0 && howMuchToCopy > limit {
		howMuchToCopy = limit
	}

	bar := pb.Full.Start64(howMuchToCopy)

	fromFileReader := io.LimitReader(fromFile, howMuchToCopy)

	_, err = fromFile.Seek(offset, 0)
	if err != nil {
		return ErrSeek
	}

	toFile, err1 := os.Create(toPath)
	if err1 != nil {
		return ErrCreateFile
	}
	defer toFile.Close()

	toFileProxyWriter := bar.NewProxyWriter(toFile)

	_, err = io.CopyN(toFileProxyWriter, fromFileReader, howMuchToCopy)
	if err != nil {
		return ErrCopy
	}
	bar.Finish()
	return nil
}
