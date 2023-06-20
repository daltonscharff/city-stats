package sources

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/daltonscharff/city-stats/internal/utils"
	"golang.org/x/exp/maps"
)

type ClimateRecord struct {
	Name    string
	Records [12]float32
	Average float32
}

type Stats struct {
	City       string
	State      string
	Climate    []ClimateRecord
	Population int
	Area       int
	Elevation  int
}

type WikipediaSearchQueryResult struct {
	Query struct {
		Pages map[string]struct {
			Pageid int    `json:"pageid"`
			Ns     int    `json:"ns"`
			Title  string `json:"title"`
			Index  int    `json:"index"`
		} `json:"pages"`
	} `json:"query,omitempty"`
}

type WikipediaPageResult struct {
	Parse struct {
		Title  string `json:"title"`
		Pageid int    `json:"pageid"`
		Text   struct {
			All string `json:"*"`
		} `json:"text"`
	} `json:"parse"`
	Error struct {
		Code string `json:"code"`
	} `json:"error,omitempty"`
}

func getPageId(query string) (string, error) {
	req, err := http.NewRequest("GET", utils.WikipediaApiUrl, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("action", "query")
	q.Add("gsrlimit", "1")
	q.Add("gsrsearch", query)
	q.Add("format", "json")
	q.Add("generator", "search")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	var data WikipediaSearchQueryResult
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return maps.Keys(data.Query.Pages)[0], nil
}

func getHtmlByPageId(pageId string) (string, error) {
	req, err := http.NewRequest("GET", utils.WikipediaApiUrl, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("action", "parse")
	q.Add("pageid", pageId)
	q.Add("format", "json")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	var data WikipediaPageResult
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if data.Error.Code != "" {
		return "", errors.New(data.Error.Code)
	}

	return data.Parse.Text.All, nil
}
