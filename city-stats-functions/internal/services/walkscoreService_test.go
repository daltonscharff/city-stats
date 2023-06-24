package services

import (
	"os"
	"path/filepath"
	"testing"

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
