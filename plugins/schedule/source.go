package schedule

import "time"

type TimeRange struct {
	From, To time.Time
}

func scheduleGenerator(duration time.Duration, out chan<- TimeRange) {
	defer close(out)

	from := time.Now()
	ticker := time.NewTicker(duration)

	for range ticker.C {
		tr := TimeRange{From: from, To: time.Now()}
		out <- tr
		from = tr.To
	}
}

// Source is an instance of a DataSource which generates TimeRange's every X duration that passes.
func Source(duration time.Duration) <-chan TimeRange {
	out := make(chan TimeRange)

	go scheduleGenerator(duration, out)

	return out
}
