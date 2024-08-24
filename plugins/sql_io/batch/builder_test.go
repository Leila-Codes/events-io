package batch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresqlBuilder(t *testing.T) {
	bb, _ := NewBuilderFromSQL("postgres", "INSERT INTO [table] (col1, col2, col3) VALUES (?1, ?2, ?3);")

	// matches.SQL(3)
	assert.Equal(t, "INSERT INTO [table] (col1, col2, col3) VALUES (?1,?2,?3),(?4,?5,?6);", bb.SQL(2))

	bb, _ = NewBuilderFromSQL("postgres", "INSERT INTO [table] (col1, col2, col3) VALUES ($1, $2, $3);")

	// matches.SQL(3)
	assert.Equal(t, "INSERT INTO [table] (col1, col2, col3) VALUES ($1,$2,$3),($4,$5,$6),($7,$8,$9);", bb.SQL(3))
}

func TestMssqlBuilder(t *testing.T) {
	bb, _ := NewBuilderFromSQL("sqlserver", "INSERT INTO [table] (col1, col2, col3) VALUES (@Username, @Email, @Password);")
	assert.Equal(t, "INSERT INTO [table] (col1, col2, col3) VALUES (@Username1,@Email1,@Password1),(@Username2,@Email2,@Password2);", bb.SQL(2))
}

func TestMysqlBuilder(t *testing.T) {
	bb, _ := NewBuilderFromSQL("mysql", "INSERT INTO [table] (col1, col2, col3) VALUES (?, ?, ?);")
	assert.Equal(t, "INSERT INTO [table] (col1, col2, col3) VALUES (?,?,?),(?,?,?);", bb.SQL(2))
}
