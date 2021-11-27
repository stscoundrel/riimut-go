package temp

import (
	"testing"
)

func TestHDummyMethod(t *testing.T) {
	result := dummyMethod()

	if result != "Lorem ipsum" {
		t.Error("Dummy method did not return correct string")
	}
}
