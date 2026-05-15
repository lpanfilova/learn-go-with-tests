package iteration

import "strings"

func Repeat(character string, x int) string {
	var repeated strings.Builder

	for i := 0; i < x; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}
