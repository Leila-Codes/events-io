module sql-generator-example

go 1.21

require (
	github.com/Leila-Codes/events-io v0.2.0
	github.com/lib/pq v1.10.9
)

replace github.com/Leila-Codes/events-io v0.2.0 => ../../

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/segmentio/kafka-go v0.4.42 // indirect
)
