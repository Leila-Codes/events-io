package sql_io

import "database/sql"

func getConnection(driverName, connString string) (*sql.DB, error) {
	return sql.Open(driverName, connString)
}
