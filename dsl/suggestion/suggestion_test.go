package suggestion

import (
	"github.com/dnotez/cli/test"
	"testing"
)

func TestSuggest(t *testing.T) {
	test.Serve("{}")
	defer test.StopServer()

	suggestion, _, err := Suggest("c", "")
	if err != nil {
		t.Errorf("Error: %s\n", err)
		return
	}

	if suggestion == nil {
		t.Error("suggestion should not be nil\n")
		return
	}
}
