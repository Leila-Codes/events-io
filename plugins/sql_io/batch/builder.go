package batch

import (
	"fmt"
	"regexp"
	"strings"
)

// Requirements
// Construct from a string sql query
// Extract VALUES statement
// Extend VALUES to include a configurable # of batches

type Builder struct {
	driver  BatchDriver
	query   string
	columns []Column
}

var (
	reValueMatcher = regexp.MustCompile(`[\( ]?(?P<Name>[A-z\?\$@]+)(?P<Index>\d+)?[,\)]`)
)

func NewBuilderFromSQL(driverName, sql string) (*Builder, error) {
	var driver BatchDriver
	switch driverName {
	case "mysql":
		driver = &MySQLDriver{}
	case "postgres":
		driver = &PostgreSQLDriver{}
	case "sqlserver":
		driver = &SQLServerDriver{}
	}

	// query 	= INSERT INTO [table] (col1, col2, col3)
	// values 	= (val1, val2, val3)
	query, values, ok := strings.Cut(sql, "VALUES ")
	if !ok {
		_, values, ok = strings.Cut(sql, "values ")
		if !ok {
			return nil, fmt.Errorf("unsupported sql expression, could not find values expression")
		}
	}

	matches := reValueMatcher.FindAllStringSubmatch(values, -1)
	cols := make([]Column, 0)

	for _, match := range matches {
		cols = append(cols, ColumnFromMatch(match))
	}

	return &Builder{
		driver:  driver,
		query:   query,
		columns: cols,
	}, nil
}

// SQL generates a new SQL statement from the internally stored query.
// "batches" represents the number of batches required for this insert statement.
// The function generates this many (x, y, z) values statements at the end of the query.
func (bb *Builder) SQL(batches int) string {
	sb := strings.Builder{}

	sb.WriteString(bb.query)
	sb.WriteString("VALUES ")

	for batchNo := 0; batchNo < batches; batchNo++ {
		if batchNo > 0 {
			sb.WriteRune(',')
		}

		sb.WriteRune('(')
		sb.WriteString(bb.driver.AddBatch(batchNo, bb.columns))
		sb.WriteRune(')')
	}

	sb.WriteRune(';')
	return sb.String()
}
