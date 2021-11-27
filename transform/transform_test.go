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
		t.Error("Did not transform by dictionary. Received", result, "expectet Yh, kämälä hööpö")
	}
}
