package sql_io

import (
	"time"

	"github.com/Leila-Codes/events-io/plugins/sql_io/batch"
)

type SqlValuer[IN interface{}] func(IN) []interface{}

func sqlDataSinker[IN interface{}](
	input <-chan IN, // input channel of events
	driverName string, // e.g. "mysql", "postgres"
	connString string, // e.g. "(user:pass)@host:port/db"
	sql string, // e.g. INSERT INTO MyTable (ID, Username, Password) VALUES (@ID, @Username, @Password)
	valuer SqlValuer[IN], // function that translated your input event type to a list of sql column values
	batchSize int, // max size of each batch
	batchTimeoutDuration time.Duration, // timeout at which batches are flushed regardless of if max is reached
) error {
	db, err := getConnection(driverName, connString)
	if err != nil {
		return err
	}

	builder, err := batch.NewBuilderFromSQL(driverName, sql)
	if err != nil {
		return err
	}

	var (
		batches      = make([]interface{}, 0)
		batchTimeout = time.NewTimer(batchTimeoutDuration)
	)

	for {
	BatchCollector:
		for {
			// Append this event for batch inserts
			batches = append(batches, valuer(<-input)...)

			// check if we've reached batch size (or timeout has been exceeded)
			select {
			case <-batchTimeout.C: // batch timeout exceeded, flush all buffered
				break BatchCollector
			default:
				// batch size has been reached
				if len(batches) > batchSize {
					break BatchCollector
				}
			}
		}

		// pause timer (if not already stopped)
		batchTimeout.Stop()

		// Generate SQL and execute on the DB.
		_, err := db.Exec(
			builder.SQL(len(batches)),
			batches...,
		)

		// Return errors for handling
		if err != nil {
			return err
		}

		batchTimeout.Reset(batchTimeoutDuration)
	}
}

const (
	ErrInsertSyntax = "unrecognised or unsupported INSERT statement"
)

func DataSink[IN interface{}](
	input <-chan IN, // input channel of events
	driverName, connString, // e.g. "mysql", "postgres"   e.g. "(user:pass)@host:port/db"
	insertStmt string, // e.g. INSERT INTO MyTable (ID, Username, Password) VALUES (@ID, @Username, @Password)
	valuer SqlValuer[IN],
	batchSize int, // max # of rows to batch before inserting into DB
	batchTimeout time.Duration, // max timeout before writing buffered data into
) {
	sqlDataSinker(
		input,
		driverName, connString,
		insertStmt,
		valuer,
		batchSize,
		batchTimeout,
	)
}
