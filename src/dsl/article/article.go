package article

import (
	"config"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Article struct {
	Id       string `json:"id"`
	Url      string `json:"url"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Text     string `json:"text"`
	MimeType string `json:"mimeType"`
	Md5      string `json:"md5"`
	Label    string `json:"label"`
	SaveDate int64  `json:"saveDate"`
}

type ArticleResult struct {
	Score float32 `json:"score"`
	Item  Article `json:"item"`
}

type ArticleResultsResponse struct {
	Results []ArticleResult `json:"results`
}

type IdResponse struct {
}

func Remove(id string) (*IdResponse, time.Duration, error) {
	start := time.Now()
	req, err := http.NewRequest("DELETE", config.SERVER_URL+"/cli/cmd/"+id, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, 0, errors.New("Resource not found.")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, 0, errors.New(fmt.Sprintf("Invalid server response:%s", resp.Status))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	var response IdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, 0, err
	}
	return &response, time.Since(start), err
}

func Get(key string, keyType string, count int) (*[]Article, time.Duration, error) {
	start := time.Now()
	url := fmt.Sprintf("%s/cli/cmd/%s?k=%s&n=%d", config.SERVER_URL, key, keyType, count)
	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, 0, errors.New(fmt.Sprintf("Invalid response status:%s", resp.Status))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	var response ArticleResultsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, 0, err
	}
	articles := make([]Article, len(response.Results))
	for i, a := range response.Results {
		articles[i] = a.Item
	}
	return &articles, time.Since(start), err

}
