package sources

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/daltonscharff/city-stats/internal/utils"
	"golang.org/x/exp/maps"
)

type WikipediaClimateRecord struct {
	Name  string
	Month [12]float32
	Year  float32
}

type WikipediaStats struct {
	City         string
	State        string
	Population   int
	AreaSqMi     float32
	ElevationFt  int
	ClimateTable []WikipediaClimateRecord
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

func parseHtml(body string) (WikipediaStats, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return WikipediaStats{}, err
	}

	return WikipediaStats{
		parseCity(doc),
		parseState(doc),
		parsePopulation(doc),
		parseAreaSqFt(doc),
		parseElevationFt(doc),
		parseClimateTable(doc),
	}, nil
}

func parseCity(doc *goquery.Document) string {
	return doc.Find("div.mw-parser-output p b").First().Text()
}

func parseState(doc *goquery.Document) string {
	return doc.Find(".infobox.vcard tr").FilterFunction(func(i int, s *goquery.Selection) bool {
		th := s.Find("th").Text()
		return strings.ToLower(th) == "state"
	}).First().Find("td").Text()
}

func parsePopulation(doc *goquery.Document) int {
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

func parseAreaSqFt(doc *goquery.Document) float32 {
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

func parseElevationFt(doc *goquery.Document) int {
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

func parseClimateTable(doc *goquery.Document) []WikipediaClimateRecord {
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
