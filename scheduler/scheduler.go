package scheduler

import "time"

// ScheduleFunc is the function to be run by the scheduler.
// Note: the returned channel MUST be closed if you want the scheduler to run again.
type ScheduleFunc[OUT interface{}] func(start, end time.Time) chan OUT

func scheduleRunner[OUT interface{}](duration time.Duration, job ScheduleFunc[OUT], out chan<- OUT) {
	timer := time.NewTicker(duration)

	defer close(out)

	for {
		<-timer.C

		// call job into var
		jobResult := job(time.Now(), time.Now().Add(24*time.Hour))

		// read until channel closure and push to out
		for result := range jobResult {
			out <- result
		}
	}
}

func Schedule[OUT interface{}](duration time.Duration, job ScheduleFunc[OUT]) chan OUT {
	out := make(chan OUT)

	go scheduleRunner(duration, job, out)

	return out
}
