package youngerfuthark

import (
	"testing"
)

func TestDefaultTransformLettersToRunes(t *testing.T) {
	const content = "aábcdðeéfghiíjklmnoópqrstþuúvwxyýzåäæöøǫþ"
	const expected = "ᛅᛅᛒᛋᛏᚦᛁᛁᚠᚴᚼᛁᛁᛁᚴᛚᛘᚾᚢᚢᛒᚴᚱᛋᛏᚦᚢᚢᚢᚢᛋᚢᚢᛋᚢᛅᛅᚢᚢᚢᚦ"
	result := LettersToRunes(content)

	if result != expected {
		t.Error("Did not transform letters to runes. Received", result, "expected ", expected)
	}
}

func TestTransformsLettersToLongBranchRunes(t *testing.T) {
	const content = "aábcdðeéfghiíjklmnoópqrstþuúvwxyýzåäæöøǫþ"
	const expected = "ᛅᛅᛒᛋᛏᚦᛁᛁᚠᚴᚼᛁᛁᛁᚴᛚᛘᚾᚢᚢᛒᚴᚱᛋᛏᚦᚢᚢᚢᚢᛋᚢᚢᛋᚢᛅᛅᚢᚢᚢᚦ"
	result := LettersToLongBranchRunes(content)

	if result != expected {
		t.Error("Did not transform letters to runes. Received", result, "expected ", expected)
	}
}

func TestTransformsLettersToShortTwigRunes(t *testing.T) {
	const content = "aábcdðeéfghiíjklmnoópqrstþuúvwxyýzåäæöøǫþ"
	const expected = "ᛆᛆᛒᛌᛐᚦᛁᛁᚠᚴᚽᛁᛁᛁᚴᛚᛘᚿᚢᚢᛒᚴᚱᛌᛐᚦᚢᚢᚢᚢᛌᚢᚢᛌᚢᛆᛆᚢᚢᚢᚦ"
	result := LettersToShortTwigRunes(content)

	if result != expected {
		t.Error("Did not transform letters to runes. Received", result, "expected ", expected)
	}
}

func TestTransformsRunesToletters(t *testing.T) {
	const expected = "fuþorkhhnniaassttbmlR "
	const longBranch = "ᚠᚢᚦᚬᚱᚴᚼᚽᚾᚿᛁᛅᛆᛋᛌᛏᛐᛒᛘᛚᛦ:"
	const shortTwig = "ᚠᚢᚦᚬᚱᚴᚽᚽᚿᚿᛁᛆᛆᛌᛌᛐᛐᛒᛘᛚᛦ:"

	// Both variants should produce same letters.
	result1 := RunesToLetters(longBranch)
	result2 := RunesToLetters(shortTwig)

	if result1 != expected {
		t.Error("Did not transform runes to letters. Received", result1, "expected ", expected)
	}

	if result2 != expected {
		t.Error("Did not transform runes to letters. Received", result2, "expected ", expected)
	}
}
