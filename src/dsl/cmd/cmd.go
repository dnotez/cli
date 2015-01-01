package cmd

/**
 * A bash cmd to be saved on the remote server
 **/

import (
	"bytes"
	"config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/user"
	"time"
)

type SaveCmdRequest struct {
	Body  string `json:"body"`
	Label string `json:"label"`
	User  string `json:"user"`
}

type SaveCmdResponse struct {
	URL string `json:"url"`
}

func (request *SaveCmdRequest) Submit() (*SaveCmdResponse, time.Duration, error) {
	current, err := user.Current()
	if err != nil {
		return nil, 0, err
	}

	request.User = current.Username
	start := time.Now()
	buf, err := json.Marshal(request)
	if err != nil {
		return nil, 0, err
	}
	resp, err := http.Post(config.SERVER_URL+"/cli/cmd", "application/json", bytes.NewReader(buf))
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	var response SaveCmdResponse
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, 0, err
	}
	return &response, time.Since(start), nil
}

func (response *SaveCmdResponse) Stringer() string {
	return fmt.Sprintf("Response{url:%s}", response.URL)
}
