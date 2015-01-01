package ping

import (
	"testing"
)

func TestPing(t *testing.T) {
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
