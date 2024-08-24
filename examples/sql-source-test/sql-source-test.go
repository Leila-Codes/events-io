package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Leila-Codes/events-io/plugins/schedule"
	"github.com/Leila-Codes/events-io/plugins/sql_io"
	"github.com/Leila-Codes/events-io/util"

	_ "github.com/lib/pq"
)

type MyEvent struct {
	Username, Action, Message string
	Timestamp                 time.Time
}

func main() {
	events, err := sql_io.DataSource(
		schedule.Source(30*time.Second),
		"postgres",
		"postgres://postgres:postgres@localhost:5432/events_test?sslmode=disable",
		"SELECT \"username\", \"action\", \"message\", \"timestamp\" FROM my_events WHERE \"timestamp\" BETWEEN $1 AND $2",
		func(tRange schedule.TimeRange) []interface{} {
			return []interface{}{tRange.From, tRange.To}
		},
		func(row *sql.Rows) (MyEvent, error) {
			event := MyEvent{}
			err := row.Scan(
				&event.Username,
				&event.Action,
				&event.Message,
				&event.Timestamp,
			)

			return event, err
		},
		1_000,
	)

	go util.PanicHandler(err)

	for {
		fmt.Printf("%+v\n", <-events)
	}

}
