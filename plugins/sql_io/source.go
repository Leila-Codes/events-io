package sql_io

import (
	"database/sql"
	"log"
)

type SqlParamSetter[IN interface{}] func(IN) []interface{}
type SqlScanner[OUT interface{}] func(row *sql.Rows) (OUT, error)

func sqlSourceExecutor[IN, OUT interface{}](
	input <-chan IN,
	stmt *sql.Stmt,
	setter SqlParamSetter[IN],
	scanner SqlScanner[OUT],
	output chan OUT,
) {
	for event := range input {
		rows, err := stmt.Query(setter(event))
		if err != nil {
			log.Fatal("Sql Data Source Execution Error: ", err)
		}

		for rows.Next() {
			row, err := scanner(rows)
			if err != nil {
				log.Fatal("Sql Data Source Scanner Error: ", err)
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
) chan OUT {
	out := make(chan OUT, bufferSize)

	db, err := getConnection(driverName, connString)
	if err != nil {
		log.Fatal("Sql Data Source Connection Error: ", err)
	}

	stmt, err := db.Prepare(selectStmt)
	if err != nil {
		log.Fatal("Sql Data Source Statement Error: ", err)
	}

	go sqlSourceExecutor(input, stmt, setter, scanner, out)

	return out
}

func NoParams[IN interface{}](IN) []interface{} {
	return make([]interface{}, 0)
}
