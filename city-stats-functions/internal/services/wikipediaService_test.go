package services

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/daltonscharff/city-stats/internal/utils"
	"github.com/go-resty/resty/v2"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

var (
	wikipediaParseFilename = filepath.Join("..", "..", "testdata", "wikipedia_parse.json")
	wikipediaQueryFilename = filepath.Join("..", "..", "testdata", "wikipedia_query.json")
)

func TestWikipediaService_getPageId(t *testing.T) {
	client := resty.New()

	defer gock.Off()
	gock.InterceptClient(client.GetClient())
	gock.New(utils.WikipediaApiUrl).
		MatchParams(map[string]string{
			"action":    "query",
			"gsrlimit":  "1",
			"gsrsearch": "Denver,CO",
			"format":    "json",
			"generator": "search",
		}).Reply(200).
		File(wikipediaQueryFilename).Type("json")

	w := WikipediaService{client}

	id, err := w.getPageId("Denver,CO")
	assert.Nil(t, err)
	assert.Equal(t, "8522", id)
}

func TestWikipediaService_getHtmlByPageId(t *testing.T) {
	client := resty.New()

	defer gock.Off()
	gock.InterceptClient(client.GetClient())
	gock.New(utils.WikipediaApiUrl).
		MatchParams(map[string]string{
			"action": "parse",
			"pageid": "8522",
			"format": "json",
		}).
		Persist().Reply(200).
		File(wikipediaParseFilename).Type("json")
	gock.New(utils.WikipediaApiUrl).
		MatchParams(map[string]string{
			"action": "parse",
			"pageid": "-1",
			"format": "json",
		}).
		Persist().Reply(200).
		JSON(map[string](map[string]string){
			"error": {
				"code": "nosuchpageid",
			},
		})

	w := WikipediaService{client}

	t.Run("Valid pageid", func(t *testing.T) {
		var data wikipediaPageResponse
		text, err := os.ReadFile(wikipediaParseFilename)
		assert.Nil(t, err)

		json.Unmarshal(text, &data)

		html, err := w.getHtmlByPageId("8522")
		assert.Nil(t, err)
		assert.Equal(t, data.Parse.Text.All, html)
	})

	t.Run("Invalid pageid", func(t *testing.T) {
		_, err := w.getHtmlByPageId("-1")
		assert.Equal(t, err.Error(), "nosuchpageid")
	})
}

func TestWikipediaService_parseLocationData(t *testing.T) {
	var data wikipediaPageResponse
	text, err := os.ReadFile(wikipediaParseFilename)
	assert.Nil(t, err)
	json.Unmarshal(text, &data)

	w := WikipediaService{}
	s, err := w.parseLocationData(data.Parse.Text.All)
	assert.Nil(t, err)
	assert.IsType(t, WikipediaLocationSearchResult{}, s)
}

func TestWikipediaService_parseCity(t *testing.T) {
	var data wikipediaPageResponse
	text, err := os.ReadFile(wikipediaParseFilename)
	assert.Nil(t, err)
	json.Unmarshal(text, &data)

	w := WikipediaService{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data.Parse.Text.All))
	assert.Nil(t, err)
	assert.Equal(t, "Denver", w.parseCity(doc))
}

func TestWikipediaService_parseState(t *testing.T) {
	var data wikipediaPageResponse
	text, err := os.ReadFile(wikipediaParseFilename)
	assert.Nil(t, err)
	json.Unmarshal(text, &data)

	w := WikipediaService{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data.Parse.Text.All))
	assert.Nil(t, err)
	assert.Equal(t, "Colorado", w.parseState(doc))
}

func TestWikipediaService_parsePopulation(t *testing.T) {
	var data wikipediaPageResponse
	text, err := os.ReadFile(wikipediaParseFilename)
	assert.Nil(t, err)
	json.Unmarshal(text, &data)

	w := WikipediaService{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data.Parse.Text.All))
	assert.Nil(t, err)
	assert.Equal(t, 715_522, w.parsePopulation(doc))
}

func TestWikipediaService_parseAreaSqFt(t *testing.T) {
	var data wikipediaPageResponse
	text, err := os.ReadFile(wikipediaParseFilename)
	assert.Nil(t, err)
	json.Unmarshal(text, &data)

	w := WikipediaService{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data.Parse.Text.All))
	assert.Nil(t, err)
	assert.Equal(t, float32(154.726), w.parseAreaSqFt(doc))
}

func TestWikipediaService_parseElevationFt(t *testing.T) {
	var data wikipediaPageResponse
	text, err := os.ReadFile(wikipediaParseFilename)
	assert.Nil(t, err)
	json.Unmarshal(text, &data)

	w := WikipediaService{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data.Parse.Text.All))
	assert.Nil(t, err)
	assert.Equal(t, 5_276, w.parseElevationFt(doc))
}

func TestWikipediaService_parseClimateTable(t *testing.T) {
	var data wikipediaPageResponse
	text, err := os.ReadFile(wikipediaParseFilename)
	assert.Nil(t, err)
	json.Unmarshal(text, &data)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data.Parse.Text.All))
	assert.Nil(t, err)

	w := WikipediaService{}
	climateRecords := w.parseClimateTable(doc)
	fixture := []WikipediaClimateRecord{{"Record high °F", [12]float32{76, 80, 84, 90, 95, 105, 105, 105, 101, 90, 81, 79}, 105}, {"Mean maximum °F", [12]float32{65, 67.1, 74.7, 80.8, 88.3, 96.5, 99.6, 96.9, 92.9, 84.1, 73.6, 65.3}, 100.6}, {"Average high °F", [12]float32{44.6, 45.7, 55.7, 61.7, 71.2, 83.4, 89.9, 87.5, 79.6, 65.3, 52.9, 44}, 65.1}, {"Daily mean °F", [12]float32{31.7, 32.7, 41.6, 47.8, 57.4, 68.2, 75.1, 72.9, 64.8, 51.1, 39.4, 31.2}, 51.2}, {"Average low °F", [12]float32{18.7, 19.7, 27.5, 33.9, 43.6, 53, 60.2, 58.3, 50, 37, 26, 18.4}, 37.2}, {"Mean minimum °F", [12]float32{-3.8, -1.5, 9.5, 19.8, 30.2, 41.9, 51.4, 48.8, 35.9, 19.6, 5.4, -3.4}, -11}, {"Record low °F", [12]float32{-29, -25, -11, -2, 19, 30, 42, 40, 17, -2, -18, -25}, -29}, {"Average precipitation inches", [12]float32{0.38, 0.41, 0.86, 1.68, 2.16, 1.94, 2.14, 1.58, 1.35, 0.99, 0.64, 0.35}, 14.48}, {"Average snowfall inches", [12]float32{6.4, 7.6, 8.8, 6.2, 1.4, 0, 0, 0, 0.8, 3.9, 7.3, 6.6}, 49}, {"Average precipitation days", [12]float32{4.4, 5.5, 6.2, 9, 10.4, 8.1, 8.3, 7.5, 6, 5.3, 4.6, 4.4}, 79.7}, {"Average snowy days", [12]float32{5, 5.3, 4.8, 4.1, 0.8, 0, 0, 0, 0.4, 1.8, 4.6, 4.6}, 31.4}, {"Average relative humidity", [12]float32{55.2, 55.8, 53.7, 49.6, 51.7, 49.3, 47.8, 49.3, 50.1, 49.2, 56.3, 56.6}, 52}, {"Average dew point °F", [12]float32{12.7, 16.2, 19.9, 26.2, 35.8, 43.5, 48.4, 47.7, 39.6, 28.6, 21, 14.2}, 29.5}, {"Mean monthly sunshine hours", [12]float32{215.3, 211.1, 255.6, 276.2, 290, 315.3, 325, 306.4, 272.3, 249.2, 194.3, 195.9}, 3106.6}, {"Percent possible sunshine", [12]float32{72, 70, 69, 69, 65, 70, 71, 72, 73, 72, 65, 67}, 70}, {"Average ultraviolet index", [12]float32{2, 3, 5, 7, 9, 11, 11, 10, 7, 5, 3, 2}, 6}}
	assert.EqualValues(t, fixture, climateRecords)
}
