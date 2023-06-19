package sources

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumbeoFind(t *testing.T) {
	scrape = func(url string) (string, error) {
		body, err := os.ReadFile(filepath.Join("testdata", "numbeo.html"))
		assert.Nil(t, err)
		return string(body), nil
	}

	n := Numbeo{}

	t.Run("Valid location, normal case", func(t *testing.T) {
		row, err := n.Find(("Austin, TX"))
		assert.Nil(t, err)
		assert.Equal(t, row, NumbeoDataRow{"Austin, TX, United States", 76.2, 66.5, 71.5, 69.9, 85.4, 114})
	})

	t.Run("Valid location, lowercase", func(t *testing.T) {
		row, err := n.Find(("dallas"))
		assert.Nil(t, err)
		assert.Equal(t, row, NumbeoDataRow{"Dallas, TX, United States", 77.4, 53.8, 66.2, 75.7, 78.1, 112})
	})

	t.Run("Invalid location", func(t *testing.T) {
		_, err := n.Find(("hello world"))
		assert.NotNil(t, err)
	})
}
