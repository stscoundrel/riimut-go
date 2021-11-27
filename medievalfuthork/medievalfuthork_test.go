package medievalfuthork

import (
	"testing"
)

func TestTransformsLettersToRunes(t *testing.T) {
	const content = "aábcdðeéfghiíjklmnoóǫpqrstuúvwxyýzåäæœöøþ "
	const expected = "ᛆᛆᛒᚴᚦᚦᚽᚽᚠᚵᚼᛁᛁᛁᚴᛚᛘᚿᚮᚮᚰᛕᚴᚱᛋᛏᚢᚢᚠᚠᛋᛦᛦᛋᚮᛅᛅᚯᚯᚯᚦ:"
	result := LettersToRunes(content)

	if result != expected {
		t.Error("Did not transform letters to runes. Received", result, "expected ", expected)
	}
}

func TestTransformsRunesToletters(t *testing.T) {
	const content = "ᚠᚢᚦᚮᚱᚴᚼᚿᛁᛆᛌᛋᛐᛏᛒᛘᛚᛦᚯᛅᚰᛕᚽ:"
	const expected = "fuþorkhniassttbmlyøæǫᴘe "
	result := RunesToLetters(content)

	if result != expected {
		t.Error("Did not transform runes to letters. Received", result, "expected ", expected)
	}
}
