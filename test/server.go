package test

import (
	"fmt"
	"github.com/dnotez/cli/config"
	"net/http"
	"net/http/httptest"
)

/**
Helper methods for test.
*/
var ts *httptest.Server

/**
Running a http server and serving the json.
*/
func Serve(json string) {
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, json)
	}))

	config.Server.URL = ts.URL
}

func StopServer() {
	if ts != nil {
		ts.Close()
	}
}
