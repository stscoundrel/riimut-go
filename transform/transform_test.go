package transform

import (
	"testing"
)

func TestTransformsByDictionary(t *testing.T) {
	var dictionary = map[string]string{
		"a": "ä",
		"o": "ö",
		"u": "y",
	}

	result := Transform("Uh, kamala hoopo", dictionary)

	if result != "yh, kämälä hööpö" {
		t.Error("Did not transform by dictionary. Received", result, "expected Yh, kämälä hööpö")
	}
}

func TestTransformsWithSpecialCharacters(t *testing.T) {
	var dictionary = map[string]string{
		"ä": "æ",
		"ö": "ø",
		"t": "þ",
	}

	result := Transform("täää töö ", dictionary)

	if result != "þæææ þøø" {
		t.Error("Did not transform by dictionary. Received", result, "expected þææ þøø")
	}
}
