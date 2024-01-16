package source

import (
	"database/sql"
)

type SqlParamSetter[IN interface{}] func(IN) []interface{}
type SqlScanner[OUT interface{}] func(row *sql.Rows) (OUT, error)

func sqlDataSource[IN, OUT interface{}](
	db *sql.DB,
	selectStmt string,
	input <-chan IN,
	output chan<- OUT,
	setter SqlParamSetter[IN],
	scanner SqlScanner[OUT],
) {
	// read from input
	for item := range input {
		rows, err := db.Query(selectStmt, setter(item)...)
		if err != nil {
			// TODO: handle error
			panic(err)
			//continue
		}

		for rows.Next() {
			out, err := scanner(rows)
			if err != nil {
				panic(err)
			}

			output <- out
		}

		err = rows.Close()
		if err != nil {
			panic(err)
		}
	}
}

func SqlDataSource[IN, OUT interface{}](
	input <-chan IN,
	driverName,
	connString,
	selectStmt string,
	setter SqlParamSetter[IN],
	scanner SqlScanner[OUT],
) <-chan OUT {
	db, err := sql.Open(driverName, connString)
	if err != nil {
		// TODO: handle error
		panic(err)
	}

	output := make(chan OUT)
	go sqlDataSource(db, selectStmt, input, output, setter, scanner)

	return output
}
