package article

import (
	"fmt"
	"github.com/dnotez/cli/test"
	"testing"
)

func TestGet(t *testing.T) {
	test.Serve("{}")
	defer test.StopServer()

	articles, _, err := Get("3cf96223-cdec-4eb3-876f-46117a38f8f7", "id", 1)
	if err != nil {
		t.Errorf("Error in submit:%s\n", err)
		return
	}

	if articles == nil {
		t.Error("articles should not be nil\n")
		return
	}

	fmt.Printf("Articles:%v\n", *articles)

}
