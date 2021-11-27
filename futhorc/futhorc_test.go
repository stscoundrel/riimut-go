package futhorc

import (
	"testing"
)

func TestTransformsLettersToRunes(t *testing.T) {
	const content = "aábcdðeéfghiíïʒjklmnŋoóǫpqrstuúvwxyýzåäæœöøþ "
	const expected = "ᚪᚪᛒᚳᛞᛞᛖᛖᚠᚷᚻᛁᛇᛇᛇᛡᚳᛚᛗᚾᛝᚩᚩᚩᛈᚳᚱᛋᛏᚢᚢᚹᚹᛉᚣᚣᛉᚩᚫᚫᛟᛟᛟᚦ:"
	result := LettersToRunes(content)

	if result != expected {
		t.Error("Did not transform letters to runes. Received", result, "expected ", expected)
	}
}

func TestTransformsRunesToletters(t *testing.T) {
	const content = "ᚠᚢᚦᚩᚱᚳᚷᚹᚻᚾᛁᛡᛄᛇᛈᛉᛋᚴᛏᛒᛖᛗᛚᛝᛟᛞᚪᚫᚣᛠ:"
	const expected = "fuþorcgwhnijjïpxsstbemlŋœdaæyea "
	result := RunesToLetters(content)

	if result != expected {
		t.Error("Did not transform runes to letters. Received", result, "expected ", expected)
	}
}
