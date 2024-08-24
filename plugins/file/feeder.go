package file

import (
	"bufio"
	"os"
)

type EventFeeder interface {
	Scan() bool
	Bytes() []byte
	Close() error
	Err() error
}

type fileScanner struct {
	file    *os.File
	scanner *bufio.Scanner
}

func (f *fileScanner) Scan() bool {
	return f.scanner.Scan()
}

func (f *fileScanner) Bytes() []byte {
	return f.scanner.Bytes()
}

func (f *fileScanner) Close() error {
	return f.file.Close()
}

func (f *fileScanner) Err() error {
	return f.scanner.Err()
}
