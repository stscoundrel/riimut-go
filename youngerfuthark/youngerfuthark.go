package youngerfuthark

import (
	"github.com/stscoundrel/riimut/transform"
)

func LettersToRunes(content string) string {
	letterMapping := getLetterMapping()
	return transform.Transform(content, letterMapping)
}
