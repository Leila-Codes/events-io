module github.com/Leila-Codes/events-io/examples/transform-test

replace (
	github.com/Leila-Codes/events-io => ../..
	github.com/Leila-Codes/events-io/plugins/kafka => ../../plugins/kafka
)

go 1.21.1

require (
	github.com/Leila-Codes/events-io v0.0.0-00010101000000-000000000000
	github.com/Leila-Codes/events-io/plugins/kafka v0.0.0-00010101000000-000000000000
	github.com/segmentio/kafka-go v0.4.47
)

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
)
