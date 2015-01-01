package cmd

import "testing"

func TestSubmit(t *testing.T) {
	request := SaveCmdRequest{}
	response, err := request.Submit()
	if err != nil {
		t.Errorf("Error in submit:%s\n", err)
	}

	fmt.Printf("Response:%v\n", *response)
}
