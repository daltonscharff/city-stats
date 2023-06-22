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
	var s WikipediaStats

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return WikipediaStats{}, err
	}

	popIndex := -1
	areaIndex := -1
	doc.Find(".infobox.vcard tr").Each(func(i int, selection *goquery.Selection) {
		th := selection.Find("th").Text()
		td := selection.Find("td").Text()

		switch {
		case strings.ToLower(th) == "state":
			s.State = td

		case s.ElevationFt == 0 && strings.Contains(strings.ToLower(th), "elevation"):
			e := strings.Split(td, "ft")
			elevation := strings.ReplaceAll(strings.TrimSpace(e[0]), ",", "")
			intElevation, err := strconv.Atoi(elevation)
			if err != nil {
				s.ElevationFt = -1
				break
			}
			s.ElevationFt = intElevation

		case popIndex == -1 && strings.Contains(strings.ToLower(th), "population"):
			popIndex = i

		case popIndex > -1 && i == popIndex+1:
			pop := strings.ReplaceAll(td, ",", "")
			intPop, err := strconv.Atoi(pop)
			if err != nil {
				s.Population = -1
				break
			}
			s.Population = intPop

		case areaIndex == -1 && strings.Contains(strings.ToLower(th), "area"):
			areaIndex = i

		case areaIndex > -1 && i == areaIndex+1:
			a := strings.Split(td, "sq")
			area := strings.ReplaceAll(strings.TrimSpace(a[0]), ",", "")
			floatArea, err := strconv.ParseFloat(area, 32)
			if err != nil {
				s.AreaSqMi = -1
				break
			}
			s.AreaSqMi = float32(floatArea)
		}
	})

	s.City = doc.Find("div.mw-parser-output p b").First().Text()

	return s, nil
}

func parseClimateRecord(doc *goquery.Document) ([]WikipediaClimateRecord, error) {
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

	return climateRecords, nil
}
