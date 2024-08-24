package sql_io

import (
	"database/sql"

	"github.com/Leila-Codes/events-io/util"
)

type SqlParamSetter[IN interface{}] func(IN) []interface{}
type SqlScanner[OUT interface{}] func(row *sql.Rows) (OUT, error)

func sqlSourceExecutor[IN, OUT interface{}](
	input <-chan IN,
	stmt *sql.Stmt,
	setter SqlParamSetter[IN],
	scanner SqlScanner[OUT],
	output chan OUT,
	errors chan error,
) {
	for event := range input {
		rows, err := stmt.Query(setter(event))
		if err != nil {
			util.MustWriteError(err, errors)
		}

		for rows.Next() {
			row, err := scanner(rows)
			if err != nil {
				util.MustWriteError(err, errors)
			}

			output <- row
		}
	}
}

func DataSource[IN, OUT interface{}](
	input <-chan IN,
	driverName, connString string,
	selectStmt string,
	setter SqlParamSetter[IN],
	scanner SqlScanner[OUT],
	bufferSize uint64,
) (chan OUT, chan error) {
	var (
		out    = make(chan OUT, bufferSize)
		errors = make(chan error)
	)

	db, err := getConnection(driverName, connString)
	if err != nil {
		util.MustWriteError(err, errors)
	}

	stmt, err := db.Prepare(selectStmt)
	if err != nil {
		util.MustWriteError(err, errors)
	}

	go sqlSourceExecutor(input, stmt, setter, scanner, out, errors)

	return out, errors
}

func NoParams[IN interface{}](IN) []interface{} {
	return make([]interface{}, 0)
}
