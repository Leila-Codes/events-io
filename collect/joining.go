package collect

import "strings"

func Joining(in chan string, sep rune) string {
	sb := strings.Builder{}
	i := 0

	for event := range in {
		if i > 0 {
			sb.WriteRune(sep)
		}
		sb.WriteString(event)

		i++
	}

	return sb.String()
}
