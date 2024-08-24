package util

import "fmt"

func MustWriteError(
	err error,
	errors chan<- error,
) {
	select {
	case errors <- err:
		// the error was written to the channel, nothing more to do.
	default:
		// nobody is listening to the channel, so we must panic
		fmt.Println("[FATAL] unhandled events-io exception!")
		panic(err)
	}
}
