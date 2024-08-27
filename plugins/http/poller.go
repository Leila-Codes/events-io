package http

import (
	"net/http"
	"time"
)

// RequestExecutor receives "trigger" events to fire requests (usually a time.Ticker channel)
// When a trigger event is fired, it is constructed into a request using the given "builder"
// On successful web requests, the given "parser" is used to convert the http.Response into the OUT format.
func RequestExecutor[IN, OUT interface{}](
	trigger <-chan IN,
	builder RequestBuilder[IN],
	parser ResponseParser[OUT],
) (chan OUT, chan error) {
	var (
		output = make(chan OUT)
		errors = make(chan error)
	)

	go requestExecutor(trigger, builder, parser, errors, output)

	return output, errors
}

// PollEvery is a shortcut method to RequestExecutor, that implements simple HTTP polling.
// The given "method" is called on the given "url" param every "duration"
// The provided "parser" is called on every subsequent successful HTTP request.
func PollEvery[OUT interface{}](
	duration time.Duration,
	method, url string,
	parser ResponseParser[OUT],
) (chan OUT, chan error) {
	ticker := time.NewTicker(duration)
	return RequestExecutor(
		ticker.C,
		func(_ time.Time) (*http.Request, error) {
			return http.NewRequest(method, url, nil)
		},
		parser,
	)
}
