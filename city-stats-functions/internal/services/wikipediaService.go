package services

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/daltonscharff/city-stats/internal/utils"
	"github.com/go-resty/resty/v2"
	"golang.org/x/exp/maps"
)

type WikipediaService struct {
	Client *resty.Client
}

type WikipediaClimateRecord struct {
	Name  string
	Month [12]float32
	Year  float32
}

type WikipediaLocationSearchResult struct {
	City         string
	State        string
	Population   int
	AreaSqMi     float32
	ElevationFt  int
	ClimateTable []WikipediaClimateRecord
}

type wikipediaSearchQueryResponse struct {
	Query struct {
		Pages map[string]struct {
			Pageid int    `json:"pageid"`
			Ns     int    `json:"ns"`
			Title  string `json:"title"`
			Index  int    `json:"index"`
		} `json:"pages"`
	} `json:"query,omitempty"`
}

type wikipediaPageResponse struct {
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

func (w WikipediaService) LocationSearch(location string) (WikipediaLocationSearchResult, error) {
	pageId, err := w.getPageId(location)
	if err != nil {
		return WikipediaLocationSearchResult{}, err
	}

	html, err := w.getHtmlByPageId(pageId)
	if err != nil {
		return WikipediaLocationSearchResult{}, err
	}

	return w.parseLocationData(html)
}

func (w WikipediaService) getPageId(query string) (string, error) {
	var data wikipediaSearchQueryResponse

	res, err := w.Client.R().SetQueryParams(map[string]string{
		"action":    "query",
		"gsrlimit":  "1",
		"gsrsearch": query,
		"format":    "json",
		"generator": "search",
	}).SetResult(&data).Get(utils.WikipediaApiUrl)
	if err != nil {
		return "", err
	}
	if res.StatusCode() >= 300 {
		return "", errors.New(res.Status())
	}
	if len(maps.Keys(data.Query.Pages)) == 0 {
		return "", errors.New("no pageId found")
	}

	return maps.Keys(data.Query.Pages)[0], nil
}

func (w WikipediaService) getHtmlByPageId(pageId string) (string, error) {
	var data wikipediaPageResponse

	res, err := w.Client.R().SetQueryParams(map[string]string{
		"action": "parse",
		"pageid": pageId,
		"format": "json",
	}).SetResult(&data).Get(utils.WikipediaApiUrl)
	if err != nil {
		return "", err
	}
	if res.StatusCode() >= 300 {
		return "", errors.New(res.Status())
	}
	if data.Error.Code != "" {
		return "", errors.New(data.Error.Code)
	}

	return data.Parse.Text.All, nil
}

func (w WikipediaService) parseLocationData(body string) (WikipediaLocationSearchResult, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return WikipediaLocationSearchResult{}, err
	}

	return WikipediaLocationSearchResult{
		w.parseCity(doc),
		w.parseState(doc),
		w.parsePopulation(doc),
		w.parseAreaSqFt(doc),
		w.parseElevationFt(doc),
		w.parseClimateTable(doc),
	}, nil
}

func (w WikipediaService) parseCity(doc *goquery.Document) string {
	return doc.Find("div.mw-parser-output p b").First().Text()
}

func (w WikipediaService) parseState(doc *goquery.Document) string {
	return doc.Find(".infobox.vcard tr").FilterFunction(func(i int, s *goquery.Selection) bool {
		th := s.Find("th").Text()
		return strings.ToLower(th) == "state"
	}).First().Find("td").Text()
}

func (w WikipediaService) parsePopulation(doc *goquery.Document) int {
	populationStr := doc.Find(".infobox.vcard tr").FilterFunction(func(i int, s *goquery.Selection) bool {
		th := s.Find("th").Text()
		return strings.Contains(strings.ToLower(th), "population")
	}).First().Next().Find("td").Text()

	populationStr = strings.ReplaceAll(populationStr, ",", "")
	population, err := strconv.Atoi(populationStr)
	if err != nil {
		population = -1
	}
	return population
}

func (w WikipediaService) parseAreaSqFt(doc *goquery.Document) float32 {
	areaStr := doc.Find(".infobox.vcard tr").FilterFunction(func(i int, s *goquery.Selection) bool {
		th := s.Find("th").Text()
		return strings.Contains(strings.ToLower(th), "area")
	}).First().Next().Find("td").Text()

	areaStr = strings.Split(areaStr, "sq")[0]
	areaStr = strings.ReplaceAll(areaStr, ",", "")
	areaStr = strings.TrimSpace(areaStr)
	areaFloat, err := strconv.ParseFloat(areaStr, 32)
	if err != nil {
		areaFloat = -1

	}
	return float32(areaFloat)
}

func (w WikipediaService) parseElevationFt(doc *goquery.Document) int {
	elevationStr := doc.Find(".infobox.vcard tr").FilterFunction(func(i int, s *goquery.Selection) bool {
		th := s.Find("th").Text()
		return strings.Contains(strings.ToLower(th), "elevation")
	}).First().Find("td").Text()

	elevationStr = strings.Split(elevationStr, "ft")[0]
	elevationStr = strings.ReplaceAll(elevationStr, ",", "")
	elevationStr = strings.TrimSpace(elevationStr)

	elevation, err := strconv.Atoi(elevationStr)
	if err != nil {
		elevation = -1
	}

	return elevation
}

func (w WikipediaService) parseClimateTable(doc *goquery.Document) []WikipediaClimateRecord {
	tableHeader := doc.Find("table.wikitable tbody tr th").FilterFunction(func(_ int, s *goquery.Selection) bool {
		text := strings.ToLower(s.Text())
		return strings.Contains(text, "climate data")
	}).First()

	tableElement := tableHeader.ParentsFiltered("table.wikitable")

	rowElements := tableElement.Find("tr").FilterFunction(func(_ int, s *goquery.Selection) bool {
		return s.Find("td").Length() >= 12
	})

	var climateRecords []WikipediaClimateRecord

	rowElements.Each(func(_ int, s *goquery.Selection) {
		re := regexp.MustCompile(`\(.*\)`)

		name := s.Find("th").Text()
		name = re.ReplaceAllString(name, "")
		name = strings.TrimSpace(name)

		climateRecord := WikipediaClimateRecord{Name: name}

		ss := s.Find("td")
		ss.Each(func(i int, s *goquery.Selection) {
			v := s.Text()
			v = re.ReplaceAllString(v, "")
			v = strings.ReplaceAll(v, ",", "")
			v = strings.ReplaceAll(v, "−", "-")
			v = strings.TrimSpace(v)
			vFloat, err := strconv.ParseFloat(v, 32)
			if err != nil {
				vFloat = -1
			}
			vFloat32 := float32(vFloat)

			if i < 12 {
				climateRecord.Month[i] = vFloat32
			} else {
				climateRecord.Year = vFloat32
			}
		})

		climateRecords = append(climateRecords, climateRecord)
	})

	return climateRecords
}
