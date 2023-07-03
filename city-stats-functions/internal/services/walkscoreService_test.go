package services

import (
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
	walkscorePathFilename           = filepath.Join("..", "..", "testdata", "walkscore_path.json")
	walkscorePageIncompleteFilename = filepath.Join("..", "..", "testdata", "walkscore_page_incomplete.html")
	walkscorePageCompleteFilename   = filepath.Join("..", "..", "testdata", "walkscore_page_complete.html")
)

func TestWalkscoreService_getPath(t *testing.T) {
	client := resty.New()

	defer gock.Off()
	gock.InterceptClient(client.GetClient())
	gock.New(utils.WalkscoreSearchUrl).
		MatchParams(map[string]string{
			"query":         "richmond",
			"skip_entities": "0",
		}).Reply(200).
		File(walkscorePathFilename).Type("json")

	w := WalkscoreService{client}

	path, err := w.getPath("richmond")
	assert.Nil(t, err)
	assert.Equal(t, "/VA/Richmond", path)
}

func TestWalkscoreService_getHtmlByPath(t *testing.T) {
	client := resty.New()

	defer gock.Off()
	gock.InterceptClient(client.GetClient())
	gock.New(utils.WalkscoreBaseUrl + "/VA/Richmond").Reply(200).
		File(walkscorePageIncompleteFilename)

	w := WalkscoreService{client}

	body, err := w.getHtmlByPath("/VA/Richmond")
	assert.Nil(t, err)
	text, err := os.ReadFile(walkscorePageIncompleteFilename)
	assert.Nil(t, err)
	assert.Equal(t, string(text), body)
}

func TestWalkscoreService_parseLocation(t *testing.T) {
	body, err := os.ReadFile(walkscorePageCompleteFilename)
	assert.Nil(t, err)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	assert.Nil(t, err)

	w := WalkscoreService{}

	location := w.parseLocation(doc)

	assert.Equal(t, "Durham", location)
}

func TestWalkscoreService_parseCityScores(t *testing.T) {
	w := WalkscoreService{}

	table := []struct {
		name     string
		filename string
		expected WalkscoreScore
	}{{
		name:     "all scores",
		filename: walkscorePageCompleteFilename,
		expected: WalkscoreScore{30, 28, 38},
	}, {
		name:     "missing transit score",
		filename: walkscorePageIncompleteFilename,
		expected: WalkscoreScore{51, -1, 51},
	}}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			body, err := os.ReadFile(test.filename)
			assert.Nil(t, err)

			doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
			assert.Nil(t, err)

			scores := w.parseCityScores(doc)

			assert.Equal(t, test.expected, scores)
		})
	}
}

func TestWalkscoreService_parseNeighborhoodScores(t *testing.T) {
	w := WalkscoreService{}

	table := []struct {
		name                      string
		filename                  string
		expectedLength            int
		expectedFirstNeighborhood WalkscoreNeighborhood
	}{{
		name:           "all scores",
		filename:       walkscorePageCompleteFilename,
		expectedLength: 27,
		expectedFirstNeighborhood: WalkscoreNeighborhood{
			Name:       "Old West Durham",
			Population: 2_086,
			Score: WalkscoreScore{
				80, 51, 85,
			},
		},
	}, {
		name:           "missing transit score",
		filename:       walkscorePageIncompleteFilename,
		expectedLength: 113,
		expectedFirstNeighborhood: WalkscoreNeighborhood{
			Name:       "Monroe Ward",
			Population: 2_763,
			Score: WalkscoreScore{
				95, -1, 81,
			},
		},
	}}

	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			body, err := os.ReadFile(test.filename)
			assert.Nil(t, err)

			doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
			assert.Nil(t, err)

			neighborhoodScores := w.parseNeighborhoodScores(doc)

			assert.Equal(t, test.expectedLength, len(neighborhoodScores))
			assert.EqualValues(t, test.expectedFirstNeighborhood, neighborhoodScores[0])
		})
	}
}

func TestWalkscoreService_parseBody(t *testing.T) {
	body, err := os.ReadFile(walkscorePageCompleteFilename)
	assert.Nil(t, err)
	w := WalkscoreService{}

	result, err := w.parseBody(string(body))

	assert.Nil(t, err)
	assert.Equal(t, "Durham", result.Location)
	assert.Equal(t, WalkscoreScore{30, 28, 38}, result.CityScore)
	assert.Equal(t, 27, len(result.NeighborhoodScores))
}
