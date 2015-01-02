package article

import (
	"fmt"
	"github.com/dnotez/cli/test"
	"testing"
)

func TestGetInvalidResponse(t *testing.T) {
	test.Serve("{}")
	defer test.StopServer()

	articles, _, err := Get("3cf96223-cdec-4eb3-876f-46117a38f8f7", "id", 1)
	if err == nil {
		t.Error("Must throw an error\n")
		return
	}

	if articles != nil {
		t.Error("results must be empty\n")
		return
	}
}

func TestGet(t *testing.T) {
	sErr := test.ServeFile("testdata/get_article.json")
	if sErr != nil {
		t.Errorf("Could not read test json file:%s\n", sErr)
		return
	}

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
