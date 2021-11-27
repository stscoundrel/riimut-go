package transform

import (
	"strings"
)

func Transform(content string, mapping map[string]string) string {
	result := ""

	for _, letter := range []rune(content) {
		lower := strings.ToLower(string(letter))
		foundRune, runeExists := mapping[lower]

		if runeExists {
			result = result + foundRune
		} else {
			result = result + string(letter)
		}
	}

	return result
}
