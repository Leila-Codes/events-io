package scheduler

import (
	"fmt"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	job := func(start, end time.Time) chan string {
		out := make(chan string, 1)
		out <- fmt.Sprintf("Hello, World [@%s]", end.Format("02-01-2006 15:04"))
		close(out)
		return out
	}

	// configure a timeout
	// allow 5 seconds of padding
	timeout := time.NewTimer(35 * time.Second)
	// execute the scheduler
	out := Schedule(10*time.Second, job)

	rx := 0

Iterator:
	for {
		select {
		case <-timeout.C:
			t.Errorf("alloted timer exceeded, end condition not met.")
			t.FailNow()
		case <-out:
			t.Logf("receives message %d @%s", rx, time.Now().Format("02-01-2006 15:04"))
			rx++
			if rx == 3 {
				t.Log("successfully received 3 messages within timeout.")
				break Iterator
			}
		}
	}
}
