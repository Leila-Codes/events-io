package batch

import "strings"

type MySQLDriver struct{}

func (s *MySQLDriver) AddBatch(
	batch int,
	columns []Column,
) string {
	sb := strings.Builder{}

	for i, col := range columns {
		if i > 0 {
			sb.WriteRune(',')
		}

		sb.WriteString(col.String(0))
	}
	return sb.String()
}
