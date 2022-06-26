package youngerfuthark

import (
	"github.com/stscoundrel/riimut-go/internal/transform"
)

func LettersToRunes(content string) string {
	return LettersToLongBranchRunes(content)
}

func LettersToLongBranchRunes(content string) string {
	letterMapping := getLettersToLongBranchRunesMapping()
	return transform.Transform(content, letterMapping)
}

func LettersToShortTwigRunes(content string) string {
	letterMapping := getLettersToShortTwigRunesMapping()
	return transform.Transform(content, letterMapping)
}

func RunesToLetters(content string) string {
	runeMapping := getRuneMapping()
	return transform.Transform(content, runeMapping)
}
