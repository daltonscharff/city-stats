package sources

import (
	"encoding/json"
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
	Batchcomplete string `json:"batchcomplete"`
	Continue      struct {
		Gsroffset int    `json:"gsroffset"`
		Continue  string `json:"continue"`
	} `json:"continue"`
	Query struct {
		Pages map[string]struct {
			Pageid int    `json:"pageid"`
			Ns     int    `json:"ns"`
			Title  string `json:"title"`
			Index  int    `json:"index"`
		} `json:"pages"`
	} `json:"query,omitempty"`
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
