package sources

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/daltonscharff/city-stats/internal/utils"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestWikipediaGetPageId(t *testing.T) {
	defer gock.Off()

	gock.New(utils.WikipediaApiUrl).
		MatchParam("action", "query").
		MatchParam("gsrlimit", "1").
		MatchParam("gsrsearch", "Denver,CO").
		MatchParam("format", "json").
		MatchParam("generator", "search").
		Persist().Reply(200).
		File(filepath.Join("testdata", "wikipedia_query_denver.json"))

	id, err := getPageId("Denver,CO")
	assert.Nil(t, err)
	assert.Equal(t, "8522", id)
}

func TestWikipediaGetHtmlByPageId(t *testing.T) {
	defer gock.Off()

	filename := filepath.Join("testdata", "wikipedia_parse_denver.json")

	gock.New(utils.WikipediaApiUrl).
		MatchParam("action", "parse").
		MatchParam("format", "json").
		MatchParam("pageid", "8522").
		Persist().Reply(200).
		File(filename)

	gock.New(utils.WikipediaApiUrl).
		MatchParam("action", "parse").
		MatchParam("format", "json").
		MatchParam("pageid", "-1").
		Persist().Reply(200).
		JSON(map[string](map[string]string){
			"error": {
				"code": "nosuchpageid",
			},
		})

	t.Run("Valid pageid", func(t *testing.T) {
		var data WikipediaPageResult
		text, err := ioutil.ReadFile(filename)
		assert.Nil(t, err)

		json.Unmarshal(text, &data)

		html, err := getHtmlByPageId("8522")
		assert.Nil(t, err)
		assert.Equal(t, data.Parse.Text.All, html)
	})

	t.Run("Invalid pageid", func(t *testing.T) {
		_, err := getHtmlByPageId("-1")
		assert.Equal(t, err.Error(), "nosuchpageid")
	})
}
