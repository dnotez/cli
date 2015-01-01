package article

import "testing"

func TestGet(t *testing.T) {
	articles, _, err := Get("3cf96223-cdec-4eb3-876f-46117a38f8f7", "id", 1)
	if err != nil {
		t.Errorf("Error in submit:%s\n", err)
	} else {
		fmt.Printf("Articles:%v\n", *articles)
	}

}
