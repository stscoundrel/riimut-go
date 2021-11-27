package tests

import (
	"testing"

	"github.com/stscoundrel/riimut/elderfuthark"
	"github.com/stscoundrel/riimut/youngerfuthark"
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

func TestElderFutharkransformsRunesToletters(t *testing.T) {
	// From 4th century axe in Jutland
	const content = "ᚹᚨᚷᚨᚷᚨᛋᛏᛁᛉ:ᚨᛚᚢ:ᚹᛁᚻᚷᚢ:ᛋᛁᚲᛁᛃᚨᛉ:ᚨᛁᚦᚨᛚᚨᛏᚨᛉ"
	const expected = "wagagastiz alu wihgu sikijaz aiþalataz"
	result := elderfuthark.RunesToLetters(content)

	if result != expected {
		t.Error("Did not transform runes to letters. Received", result, "expected ", expected)
	}
}
