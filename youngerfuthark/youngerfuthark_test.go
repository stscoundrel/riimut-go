package youngerfuthark

import (
	"testing"
)

func TestTransformsLettersToRunes(t *testing.T) {
	const content = "aábcdðeéfghiíjklmnoópqrstþuúvwxyýzåäæöøǫþ"
	const expected = "ᛅᛅᛒᛋᛏᚦᛁᛁᚠᚴᚼᛁᛁᛁᚴᛚᛘᚾᚢᚢᛒᚴᚱᛋᛏᚦᚢᚢᚢᚢᛋᚢᚢᛋᚢᛅᛅᚢᚢᚢᚦ"
	result := LettersToRunes(content)

	if result != expected {
		t.Error("Did not transform letters to runes. Received", result, "expected ", expected)
	}
}

func TestTransformsRunesToletters(t *testing.T) {
	const content = "ᚠᚢᚦᚬᚱᚴᚼᚽᚾᚿᛁᛅᛆᛋᛌᛏᛐᛒᛘᛚᛦ:"
	const expected = "fuþorkhhnniaassttbmlR "
	result := RunesToLetters(content)

	if result != expected {
		t.Error("Did not transform runes to letters. Received", result, "expected ", expected)
	}
}
