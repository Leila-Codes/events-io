package source

import "time"

type TimeRange struct {
	From, To time.Time
}

func scheduleRunner(duration time.Duration, out chan<- TimeRange) {
	defer close(out)

	from := time.Now()
	ticker := time.NewTicker(duration)

	for {
		<-ticker.C
		to := time.Now()
		out <- TimeRange{From: from, To: to}
		from = to
	}
}

// Schedule generates TimeRange's every X duration that passes.
func Schedule(duration time.Duration) <-chan TimeRange {
	out := make(chan TimeRange)

	go scheduleRunner(duration, out)

	return out
}
