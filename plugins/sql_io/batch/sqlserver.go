package batch

import "strings"

type SQLServerDriver struct{}

func (s *SQLServerDriver) AddBatch(
	batch int,
	columns []Column,
) string {
	sb := strings.Builder{}

	for i, col := range columns {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(col.String(batch + 1))
	}
	return sb.String()
}
