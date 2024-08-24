package file

import (
	"bufio"
	"io"
	"os"

	"github.com/Leila-Codes/events-io/util"
)

func newFileScanner(filePath string) (*fileScanner, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}

	return &fileScanner{
		file:    f,
		scanner: bufio.NewScanner(f),
	}, nil
}

func fileFeeder(
	feeder EventFeeder,
	output chan<- []byte,
	errors chan<- error,
) {
	err := feeder.Err()
	for feeder.Scan() && err != nil {
		err = feeder.Err()
		output <- feeder.Bytes()
	}
	if err != nil && err != io.EOF {
		util.MustWriteError(err, errors)
	}
}

func multiFileReader(
	output chan []byte,
	errors chan error,
	filePaths ...string,
) {
	for _, filePath := range filePaths {
		feeder, err := newFileScanner(filePath)
		if err != nil {
			util.MustWriteError(err, errors)
		}

		fileFeeder(feeder, output, errors)

		err = feeder.Close()
		if err != nil {
			util.MustWriteError(err, errors)
		}
	}

	close(output)
}

// ScannerSource constructs a bufio.Scanner across each file passed in sequentially.
// Each file is opened one-by-one and read line-by-line with each line fired as an event to the returned output channel
// This channel is closed once all files have been read.
func ScannerSource(
	bufferSize uint64,
	filePaths ...string,
) (chan []byte, chan error) {
	var (
		output = make(chan []byte, bufferSize)
		errors = make(chan error)
	)

	go multiFileReader(output, errors, filePaths...)

	return output, errors
}
