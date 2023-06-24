package services

import (
	"errors"

	"github.com/daltonscharff/city-stats/internal/utils"
	"github.com/go-resty/resty/v2"
)

type WalkscoreService struct {
	Client *resty.Client
}

type WalkscoreGetPathResponse struct {
	Query       string `json:"query"`
	Suggestions []struct {
		Path string `json:"path"`
		Name string `json:"name"`
	} `json:"suggestions`
}

type WalkscoreScore struct {
	Walk    int8
	Transit int8
	Bike    int8
}

type WalkscoreNeighborhood struct {
	Name       string
	Population int32
	Score      WalkscoreScore
}

type WalkscoreSearchResult struct {
	Location      string
	AverageScore  WalkscoreScore
	Neighborhoods []WalkscoreNeighborhood
}

func (w WalkscoreService) getPath(location string) (string, error) {
	var data WalkscoreGetPathResponse

	_, err := w.Client.R().SetQueryParams(map[string]string{
		"query":         location,
		"skip_entities": "0",
	}).SetResult(&data).Get(utils.WalkscoreSearchUrl)

	if err != nil {
		return "", err
	}
	if len(data.Suggestions) == 0 {
		return "", errors.New("no suggestions found")
	}

	return data.Suggestions[0].Path, nil
}

func (w WalkscoreService) getHtmlByPath(path string) (string, error) {
	res, err := w.Client.R().Get(utils.WalkscoreBaseUrl + path)
	if err != nil {
		return "", err
	}

	return string(res.Body()), nil
}
