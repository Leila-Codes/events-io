package batch

import "strings"

type PostgreSQLDriver struct{}

func (s *PostgreSQLDriver) AddBatch(
	batch int,
	columns []Column,
) string {
	sb := strings.Builder{}

	for i, col := range columns {
		if i > 0 {
			sb.WriteRune(',')
		}

		colIdx := (i + 1) + (batch * len(columns))

		sb.WriteString(col.String(colIdx))
	}
	return sb.String()
}
