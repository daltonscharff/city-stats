package services

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/daltonscharff/city-stats/internal/models"
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
	} `json:"suggestions"`
}

type WalkscoreScore struct {
	Walk    int
	Transit int
	Bike    int
}

type WalkscoreNeighborhood struct {
	Name       string
	Population int
	Score      WalkscoreScore
}

type WalkscoreSearchResult struct {
	Location           string
	CityScore          WalkscoreScore
	NeighborhoodScores []WalkscoreNeighborhood
}

func (w WalkscoreService) getPath(query string) (string, error) {
	var data WalkscoreGetPathResponse

	_, err := w.Client.R().SetQueryParams(map[string]string{
		"query":         query,
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

func (w WalkscoreService) LocationSearch(location models.Location) (WalkscoreSearchResult, error) {
	path, err := w.getPath(fmt.Sprintf("%s,%s", location.City, location.StateAbbrev))
	if err != nil {
		return WalkscoreSearchResult{}, err
	}

	body, err := w.getHtmlByPath(path)
	if err != nil {
		return WalkscoreSearchResult{}, err
	}

	return w.parseBody(body)
}

func (w WalkscoreService) parseBody(body string) (WalkscoreSearchResult, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return WalkscoreSearchResult{}, err
	}

	location := w.parseLocation(doc)
	cityScores := w.parseCityScores(doc)
	neighborhoodScores := w.parseNeighborhoodScores(doc)

	return WalkscoreSearchResult{location, cityScores, neighborhoodScores}, nil
}

func (w WalkscoreService) parseLocation(doc *goquery.Document) string {
	location := doc.Find("#title").Text()
	location = strings.ReplaceAll(location, "Living in", "")
	location = strings.TrimSpace(location)
	return location
}

func (w WalkscoreService) parseCityScores(doc *goquery.Document) WalkscoreScore {
	var (
		walk      = -1
		transit   = -1
		bike      = -1
		walkRe    = regexp.MustCompile(`(?i)(\d+) walk score`)
		transitRe = regexp.MustCompile(`(?i)(\d+) transit score`)
		bikeRe    = regexp.MustCompile(`(?i)(\d+) bike score`)
	)

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		alt, _ := s.Attr("alt")
		walkMatch := walkRe.FindAllStringSubmatch(alt, -1)
		transitMatch := transitRe.FindAllStringSubmatch(alt, -1)
		bikeMatch := bikeRe.FindAllStringSubmatch(alt, -1)

		switch {
		case walk == -1 && len(walkMatch) > 0:
			walk, _ = strconv.Atoi(walkMatch[0][1])
		case transit == -1 && len(transitMatch) > 0:
			transit, _ = strconv.Atoi(transitMatch[0][1])
		case bike == -1 && len(bikeMatch) > 0:
			bike, _ = strconv.Atoi(bikeMatch[0][1])
		}
	})

	return WalkscoreScore{walk, transit, bike}
}

func (w WalkscoreService) parseNeighborhoodScores(doc *goquery.Document) []WalkscoreNeighborhood {
	var neighborhoodScores []WalkscoreNeighborhood
	doc.Find("#hoods-list-table tbody tr").Each(func(i int, s *goquery.Selection) {
		var neighborhood WalkscoreNeighborhood
		neighborhood.Name = s.Find(".name").First().Text()

		population, err := strconv.Atoi(strings.ReplaceAll(s.Find(".population").First().Text(), ",", ""))
		if err != nil {
			population = -1
		}
		neighborhood.Population = population

		walkScore, err := strconv.Atoi(strings.TrimSpace(s.Find(".walkscore").First().Text()))
		if err != nil {
			walkScore = -1
		}
		neighborhood.Score.Walk = walkScore

		transitScore, err := strconv.Atoi(strings.TrimSpace(s.Find(".transitscore").First().Text()))
		if err != nil {
			transitScore = -1
		}
		neighborhood.Score.Transit = transitScore

		bikeScore, err := strconv.Atoi(strings.TrimSpace(s.Find(".bikescore").First().Text()))
		if err != nil {
			bikeScore = -1
		}
		neighborhood.Score.Bike = bikeScore

		neighborhoodScores = append(neighborhoodScores, neighborhood)
	})

	return neighborhoodScores
}
