package stavelessfuthark

import (
	"github.com/stscoundrel/riimut-go/internal/transform"
)

func LettersToRunes(content string) string {
	letterMapping := getLetterMapping()
	return transform.Transform(content, letterMapping)
}

func RunesToLetters(content string) string {
	runeMapping := getRuneMapping()
	return transform.Transform(content, runeMapping)
}
