package suggestion

import (
	"testing"
)

func TestSuggest(t *testing.T) {
	suggestion, _, err := Suggest("c", "")
	if err != nil {
		t.Errorf("Error: %s\n", err)
	}

	if suggestion == nil {
		t.Error("Must not be nil\n")
	}
}
