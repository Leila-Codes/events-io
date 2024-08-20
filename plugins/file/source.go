package file

import (
	"bufio"
	"os"
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
	output chan []byte,
) {
	for feeder.Scan() {
		output <- feeder.Bytes()
	}
}

func multiFileReader(
	output chan []byte,
	filePaths ...string,
) {
	for _, filePath := range filePaths {
		feeder, err := newFileScanner(filePath)
		if err != nil {
			panic("MultiFileReader open error: " + err.Error())
		}

		fileFeeder(feeder, output)

		err = feeder.Close()
		if err != nil {
			panic("MultiFileScanner failed to close file " + filePath)
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
) chan []byte {
	output := make(chan []byte, bufferSize)

	go multiFileReader(output, filePaths...)

	return output
}
