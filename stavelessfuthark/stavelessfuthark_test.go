package stavelessfuthark

import (
	"testing"
)

func TestTransformsLettersToRunes(t *testing.T) {
	const content = "aábcdðeéfghiíjklmnoópqrRstþuúvwxyýzåäæœöøǫþ "
	const expected = "⸝⸝ˏ╵⸍וᛁᛁᛙᛍᚽᛁᛁᛁᛍ⸌⠃⸜ˎˎˏᛍ◟◟╵⸍ו╮╮╮╮╵╮╮╵ˎ⸝⸝ˎˎˎˎו:"
	result := LettersToRunes(content)

	if result != expected {
		t.Error("Did not transform letters to runes. Received", result, "expected ", expected)
	}
}

func TestTransformsRunesToletters(t *testing.T) {
	const content = "ᛙ╮וˎ⡄ᛍᚽ⸜ᛁ⸝╵⸍ˏ⠃⸌⡄:"
	const expected = "fuþoRkhniastbmlR "
	result := RunesToLetters(content)

	if result != expected {
		t.Error("Did not transform runes to letters. Received", result, "expected ", expected)
	}
}
