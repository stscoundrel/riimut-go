package tests

import (
	"testing"

	"github.com/stscoundrel/riimut-go/elderfuthark"
	"github.com/stscoundrel/riimut-go/futhorc"
	"github.com/stscoundrel/riimut-go/medievalfuthork"
	"github.com/stscoundrel/riimut-go/stavelessfuthark"
	"github.com/stscoundrel/riimut-go/youngerfuthark"
)

func TestYoungerFutharkTransformsLettersToRunes(t *testing.T) {
	// From Old Groms runestone.
	const content = "auk tani karþi kristna"
	const expected = "ᛅᚢᚴ:ᛏᛅᚾᛁ:ᚴᛅᚱᚦᛁ:ᚴᚱᛁᛋᛏᚾᛅ"
	result := youngerfuthark.LettersToRunes(content)

	if result != expected {
		t.Error("Did not transform letters to runes. Received", result, "expected ", expected)
	}
}

func TestToungerFutharkransformsRunesToletters(t *testing.T) {
	// From Old Groms runestone.
	const content = "ᛅᚢᚴ:ᛏᛅᚾᛁ:ᚴᛅᚱᚦᛁ:ᚴᚱᛁᛋᛏᚾᛅ"
	const expected = "auk tani karþi kristna"
	result := youngerfuthark.RunesToLetters(content)

	if result != expected {
		t.Error("Did not transform runes to letters. Received", result, "expected ", expected)
	}
}

func TestElderFutharkTransformsLettersToRunes(t *testing.T) {
	// From 4th century axe in Jutland
	const content = "wagagastiz alu wihgu sikijaz aiþalataz"
	const expected = "ᚹᚨᚷᚨᚷᚨᛋᛏᛁᛉ:ᚨᛚᚢ:ᚹᛁᚻᚷᚢ:ᛋᛁᚲᛁᛃᚨᛉ:ᚨᛁᚦᚨᛚᚨᛏᚨᛉ"
	result := elderfuthark.LettersToRunes(content)

	if result != expected {
		t.Error("Did not transform letters to runes. Received", result, "expected ", expected)
	}
}

func TestElderFutharkTransformsRunesToletters(t *testing.T) {
	// From 4th century axe in Jutland
	const content = "ᚹᚨᚷᚨᚷᚨᛋᛏᛁᛉ:ᚨᛚᚢ:ᚹᛁᚻᚷᚢ:ᛋᛁᚲᛁᛃᚨᛉ:ᚨᛁᚦᚨᛚᚨᛏᚨᛉ"
	const expected = "wagagastiz alu wihgu sikijaz aiþalataz"
	result := elderfuthark.RunesToLetters(content)

	if result != expected {
		t.Error("Did not transform runes to letters. Received", result, "expected ", expected)
	}
}

func TestMedievalFuthorkTransformsLettersToRunes(t *testing.T) {
	// From Lord's Prayer, in Old Norse.
	const content = "Faðer uor som ast i himlüm, halgað warðe þit nama"
	const expected = "ᚠᛆᚦᚽᚱ:ᚢᚮᚱ:ᛋᚮᛘ:ᛆᛋᛏ:ᛁ:ᚼᛁᛘᛚᚢᛘ,:ᚼᛆᛚᚵᛆᚦ:ᚠᛆᚱᚦᚽ:ᚦᛁᛏ:ᚿᛆᛘᛆ"
	result := medievalfuthork.LettersToRunes(content)

	if result != expected {
		t.Error("Did not transform letters to runes. Received", result, "expected ", expected)
	}
}

func TestMedievalFuthorkTransformsRunesToletters(t *testing.T) {
	// From Lord's Prayer, in Old Norse.
	const content = "ᚠᛆᚦᚽᚱ:ᚢᚮᚱ:ᛋᚮᛘ:ᛆᛋᛏ:ᛁ:ᚼᛁᛘᛚᚢᛘ:ᚼᛆᛚᚵᛆᚦ:ᚠᛆᚱᚦᚽ:ᚦᛁᛏ:ᚿᛆᛘᛆ"
	const expected = "faþer uor som ast i himlum halgaþ farþe þit nama" // Wont tell apart eth & thorn in mid sentence.
	result := medievalfuthork.RunesToLetters(content)

	if result != expected {
		t.Error("Did not transform runes to letters. Received", result, "expected ", expected)
	}
}

func TestFuthorcTransformsLettersToRunes(t *testing.T) {
	// From 8th century Franks Casket, in late West Saxon.
	const content = "fisc.flodu.ahofonferg | enberig |"
	const expected = "ᚠᛁᛋᚳ.ᚠᛚᚩᛞᚢ.ᚪᚻᚩᚠᚩᚾᚠᛖᚱᚷ:|:ᛖᚾᛒᛖᚱᛁᚷ:|"
	result := futhorc.LettersToRunes(content)

	if result != expected {
		t.Error("Did not transform letters to runes. Received", result, "expected ", expected)
	}
}

func TestFuthorcTransformsRunesToletters(t *testing.T) {
	// From 8th century Franks Casket, in late West Saxon.
	const content = "ᚠᛁᛋᚳ.ᚠᛚᚩᛞᚢ.ᚪᚻᚩᚠᚩᚾᚠᛖᚱᚷ:|:ᛖᚾᛒᛖᚱᛁᚷ:|"
	const expected = "fisc.flodu.ahofonferg | enberig |"
	result := futhorc.RunesToLetters(content)

	if result != expected {
		t.Error("Did not transform runes to letters. Received", result, "expected ", expected)
	}
}

func TestStavelessFutharkTransformsLettersToRunes(t *testing.T) {
	// From Old Norse text in Hög runestone.
	const content = "kuþniutr þru sun lit rita stin þina ak bru kirþi aftiʀ bruþr sina asbiurn ak at kuþlaf"
	const expected = "ᛍ╮ו⸜ᛁ╮⸍◟:ו◟╮:╵╮⸜:⸌ᛁ⸍:◟ᛁ⸍⸝:╵⸍ᛁ⸜:וᛁ⸜⸝:⸝ᛍ:ˏ◟╮:ᛍᛁ◟וᛁ:⸝ᛙ⸍ᛁʀ:ˏ◟╮ו◟:╵ᛁ⸜⸝:⸝╵ˏᛁ╮◟⸜:⸝ᛍ:⸝⸍:ᛍ╮ו⸌⸝ᛙ"
	result := stavelessfuthark.LettersToRunes(content)

	if result != expected {
		t.Error("Did not transform letters to runes. Received", result, "expected ", expected)
	}
}

func TestStavelessFutharkTransformsRunesToletters(t *testing.T) {
	// From Old Norse text in Hög runestone.
	const content = "ᛍ╮ו⸜ᛁ╮⸍◟:ו◟╮:╵╮⸜:⸌ᛁ⸍:◟ᛁ⸍⸝:╵⸍ᛁ⸜:וᛁ⸜⸝:⸝ᛍ:ˏ◟╮:ᛍᛁ◟וᛁ:⸝ᛙ⸍ᛁʀ:ˏ◟╮ו◟:╵ᛁ⸜⸝:⸝╵ˏᛁ╮◟⸜:⸝ᛍ:⸝⸍:ᛍ╮ו⸌⸝ᛙ"
	const expected = "kuþniutr þru sun lit rita stin þina ak bru kirþi aftiʀ bruþr sina asbiurn ak at kuþlaf"
	result := stavelessfuthark.RunesToLetters(content)

	if result != expected {
		t.Error("Did not transform runes to letters. Received", result, "expected ", expected)
	}
}
