package ping

import (
	"encoding/json"
	"net/http"

	"github.com/dnotez/cli/config"
)

type Pong struct {
	Okay bool
}

func Ping() (pong *Pong, err error) {
	var p Pong
	r, err := http.Get(config.Server.URL + "/api/ping")
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	err = decoder.Decode(&p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
