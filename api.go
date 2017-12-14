package chinese

import (
	"strings"
)

func Convert(txt string, style int, segment bool) string {
	return strings.Join((&py{style, segment}).convert(txt), "")
}

func ConvertEx(txt string, style int, segment bool) []string {
	return (&py{style, segment}).convert(txt)
}
