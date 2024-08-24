module github.com/Leila-Codes/events-io/examples/sql-generator

replace (
	github.com/Leila-Codes/events-io => ../..
	github.com/Leila-Codes/events-io/plugins/sql_io => ../../plugins/sql_io
)

go 1.21.1

require (
	github.com/Leila-Codes/events-io/plugins/sql_io v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.9
)

require (
	github.com/Leila-Codes/events-io v0.0.0-00010101000000-000000000000 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
)
