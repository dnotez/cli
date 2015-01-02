package cmd

import (
	"fmt"
	"github.com/dnotez/dnotez-cli/test"
	"testing"
)

func TestSubmit(t *testing.T) {
	test.Serve("{}")
	defer test.StopServer()
	request := SaveCmdRequest{}
	response, _, err := request.Submit()
	if err != nil {
		t.Errorf("Error in submit:%s\n", err)
		return
	}
	if response == nil {
		t.Error("response should not be nil\n")
		return
	}

	fmt.Printf("Response:%v\n", *response)
}
