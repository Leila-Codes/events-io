package source

import (
	"database/sql"
	_ "github.com/lib/pq"
	"testing"
	"time"
)

type MyRowResult struct {
	Username string
	Action   string
	Message  string
}

func TestSqlDataSource_Postgres(t *testing.T) {
	// generate some timestamps
	timestamps := Schedule(30 * time.Second)

	// query for data in the time range
	data := SqlDataSource(
		timestamps,
		"postgres",
		"postgres://postgres:postgres@localhost:5432/events_io_test?sslmode=disable",
		"SELECT \"username\", \"action\", \"message\" FROM my_events WHERE \"timestamp\" BETWEEN $1 AND $2",
		func(tRange TimeRange) []interface{} {
			return []interface{}{tRange.From, tRange.To}
		},
		func(row *sql.Rows) (MyRowResult, error) {
			res := MyRowResult{}
			err := row.Scan(
				&res.Username,
				&res.Action,
				&res.Message,
			)

			return res, err
		})

	i := 0
	for {
		t.Logf("%+v", <-data)

		i++

		if i >= 3 {
			break
		}
	}
}
