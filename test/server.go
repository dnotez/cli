package test

import (
	"fmt"
	"github.com/dnotez/dnotez-cli/config"
	"io/ioutil"
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

func ServeFile(filePath string) error {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	Serve(string(buf))
	return nil
}

func StopServer() {
	if ts != nil {
		ts.Close()
	}
}
