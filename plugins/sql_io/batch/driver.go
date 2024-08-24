package batch

// Support MySQL, MS-SQL and PostgreSQL syntax variables

// Syntax variations
// POSTGRESQL:
//	- Unnamed, but indexed/numbered columns
//  E.g. VALUES (?1, ?2, ?3), (?4, ?5, ?6)

// MS-SQL:
// 	- Named, indexed by instance columns
//	E.g. VALUES (@Name1, @Username1, @Password1), (@Name2, @Username2, @Password2)

// MYSQL
//	- Non-named, non-indexed columns
//	E.g. VALUES (?, ?, ?), (?, ?, ?)

// BatchDriver, given a batch number and a list of columns, return the VALUES statement.
type BatchDriver interface {
	AddBatch(batch int, columns []Column) string
}
