package batch

import (
	"strconv"
	"strings"
)

type Column string

func ColumnFromMatch(match []string) Column {
	return Column(match[1])
}

func (c Column) String(index int) string {
	sb := strings.Builder{}

	sb.WriteString(string(c))

	if index > 0 {
		sb.WriteString(strconv.Itoa(index))
	}

	return sb.String()
}
