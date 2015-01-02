package ping

import (
	"github.com/dnotez/cli/test"
	"testing"
)

func TestPing(t *testing.T) {
	test.Serve("{\"okay\":true}")
	defer test.StopServer()
	pong, err := Ping()
	if err != nil {
		t.Errorf("error: %s\n", err)
	}
	if pong == nil {
		t.Errorf("Invalid response")
	}
	if !pong.Okay {
		t.Errorf("Invalid response, check server")
	}
}
