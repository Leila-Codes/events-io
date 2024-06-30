package sql_io

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type SqlValuer[IN interface{}] func(IN) []interface{}

var (
	reNumericStrip = regexp.MustCompile("\\d")
)

func sqlDataSinker[IN interface{}](
	db *sql.DB,
	insertStmt string,
	valuesStmt string,
	input <-chan IN,
	valuer SqlValuer[IN],
	batchSize int,
	batchTimeoutDuration time.Duration,
) {
	columnParams := strings.Split(reNumericStrip.ReplaceAllString(valuesStmt[1:len(valuesStmt)-1], ""), ",")

	for {

		var (
			rows   int
			values []interface{}
		)

		batchTimeout := time.NewTimer(batchTimeoutDuration)

	BatchCollector:
		for rows = 0; rows < batchSize; rows++ {
			select {
			// if timeout has been reached, exit batch collection phase.
			case _ = <-batchTimeout.C:
				break BatchCollector
			// otherwise, is more data available?
			case data := <-input:
				values = append(values, valuer(data)...)
			}

			time.Sleep(time.Millisecond)
		}

		totalCols := 0

		// Make sure, rows are available
		if rows > 0 {

			// Flush records to database
			query := strings.Builder{}
			query.WriteString(insertStmt)
			query.WriteString(" VALUES ")
			//query.WriteString(valuesStmt)

			for r := 0; r < rows; r++ {
				if r > 0 {
					query.WriteRune(',')
				}

				query.WriteRune('(')

				for c, col := range columnParams {
					if c > 0 {
						query.WriteRune(',')
					}
					totalCols++
					query.WriteString(col + strconv.Itoa(totalCols))
				}

				query.WriteRune(')')
			}

			queryStr := query.String()

			res, err := db.Exec(queryStr, values...)
			if err != nil {
				panic(err)
			}

			rowCount, _ := res.RowsAffected()
			fmt.Println("inserted ", rowCount, "rows")
		}

	}
}

const (
	ErrInsertSyntax = "unrecognised or unsupported INSERT statement"
)

func DataSink[IN interface{}](
	input <-chan IN,
	driverName,
	connString,
	insertStmt string,
	valuer SqlValuer[IN],
	batchSize int,
	batchTimeout time.Duration,
) {

	// 1) Connect to database
	db, err := sql.Open(driverName, connString)
	if err != nil {
		panic(err)
	}

	// 2) Partially parse the SQL insert statement (extract VALUES)

	parts := strings.Split(insertStmt, " VALUES ")
	if len(parts) != 2 {
		panic(ErrInsertSyntax)
	}

	insertStmt = parts[0]
	valuesStmt := parts[1]

	// 3) Pass over to dataSinker routine.
	sqlDataSinker[IN](
		db,
		insertStmt,
		valuesStmt,
		input,
		valuer,
		batchSize,
		batchTimeout,
	)
}
