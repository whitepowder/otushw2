package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

var (
	// ErrUnsupportedFile is perfect.
	ErrUnsupportedFile = errors.New("unsupported file")
	// ErrOffsetExceedsFileSize is also perfect.
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	// Err404 is fine.
	Err404 = errors.New("404 Not found")
	// ErrCantCreate kinda meh.
	ErrCantCreate = errors.New("can't create file")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	srcFile, err := os.Open(fromPath)
	if err != nil {
		return Err404
	}

	fileInfo, err := os.Stat(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	fileSize := fileInfo.Size()
	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}

	if limit == 0 || limit > fileSize {
		limit = fileSize - offset
	}

	defer srcFile.Close()

	fileTransaction := io.LimitReader(srcFile, limit)

	count := 100
	bar := pb.StartNew(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}

	destFile, err := os.Create(toPath)
	if err != nil {
		return ErrCantCreate
	}
	defer destFile.Close()

	_, err = io.CopyN(destFile, fileTransaction, limit)
	if err != nil {
		log.Fatal(err)
	}

	bar.Finish()

	fmt.Printf("File copied successfully. Initial size - %v, copied - %v. From %v - to %v", fileSize, limit, srcFile, destFile)
	return nil
}
