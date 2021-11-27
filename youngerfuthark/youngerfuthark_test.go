package youngerfuthark

import (
	"testing"
)

func TestTransformsLettersToRunes(t *testing.T) {
	const content = "aábcdðeéfghiíjklmnoópqrstþuúvwxyýzåäæöøǫþ"
	const expected = "ᛅᛅᛒᛋᛏᚦᛁᛁᚠᚴᚼᛁᛁᛁᚴᛚᛘᚾᚢᚢᛒᚴᚱᛋᛏᚦᚢᚢᚢᚢᛋᚢᚢᛋᚢᛅᛅᚢᚢᚢᚦ"
	result := LettersToRunes(content)

	if result != expected {
		t.Error("Did not transform by dictionary. Received", result, "expected ", expected)
	}
}
