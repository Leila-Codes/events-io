module github.com/Leila-Codes/events-io/examples/file-example

replace (
	github.com/Leila-Codes/events-io => ../..
	github.com/Leila-Codes/events-io/plugins/file => ../../plugins/file
	github.com/Leila-Codes/events-io/plugins/kafka => ../../plugins/kafka
)

go 1.21.1

require (
	github.com/Leila-Codes/events-io v0.0.0
	github.com/Leila-Codes/events-io/plugins/file v0.0.0-20240824084121-fb01b121fd6a
	github.com/Leila-Codes/events-io/plugins/kafka v0.0.0-20240824084121-fb01b121fd6a
	github.com/segmentio/kafka-go v0.4.47
)

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
)
