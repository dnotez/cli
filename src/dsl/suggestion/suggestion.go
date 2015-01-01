package suggestion

import (
	"bytes"
	"config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SuggestionResult struct {
	Suggestion string `json:"suggestion"`
	ID         string `json:"id"`
}

func (r *SuggestionResult) Stringer() string {
	return fmt.Sprintf("%s - %s", r.ID, r.Suggestion)
}

type Request struct {
	Query string `json:"query"`
}

type SuggestionResponse struct {
	Results []SuggestionResult `json:"results`
}

func (r *SuggestionResponse) Stringer() string {
	s := fmt.Sprintf("%d results", len(r.Results))
	for _, result := range r.Results {
		s += "\n"+result.Stringer()
	}
	return s
}

func (r *Request) String() string {
	//var r io.Reader
	return "request"
}

func Suggest(query string, resourceType string) (*SuggestionResponse, time.Duration, error) {
	start := time.Now()
	request := Request{Query: query}
	buf, err := json.Marshal(request)
	if err != nil {
		return nil, 0, err
	}
	resp, err := http.Post(config.SERVER_URL+"/extension/suggestion", "application/json", bytes.NewReader(buf))
	if err != nil {
		return nil, 0, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	var response SuggestionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, 0, err
	}
	return &response, time.Since(start), err
}
