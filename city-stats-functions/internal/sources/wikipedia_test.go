package sources

import (
	"path/filepath"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestWikipediaGetPageId(t *testing.T) {
	defer gock.Off()

	gock.New("https://en.wikipedia.org/w/api.php").
		MatchParam("action", "query").
		MatchParam("gsrlimit", "1").
		MatchParam("gsrsearch", "Denver,CO").
		MatchParam("format", "json").
		MatchParam("generator", "search").
		Persist().Reply(200).
		File(filepath.Join("testdata", "wikipedia_search_denver.json"))

	id, err := getPageId("Denver,CO")
	assert.Nil(t, err)
	assert.Equal(t, "8522", id)
}
