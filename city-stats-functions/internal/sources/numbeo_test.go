package sources

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestNumbeoFind(t *testing.T) {
	defer gock.Off()

	gock.New("https://www.numbeo.com").
		Get("/cost-of-living/rankings_current.jsp").Persist().
		Reply(200).File(filepath.Join("testdata", "numbeo.html"))

	n := Numbeo{}

	t.Run("Valid location, normal case", func(t *testing.T) {
		row, err := n.Find(("Austin, TX"))
		assert.Nil(t, err)
		assert.Equal(t, row, NumbeoDataRow{"Austin, TX, United States", 76.2, 66.5, 71.5, 69.9, 85.4, 114})
	})

	t.Run("Valid location, lowercase", func(t *testing.T) {
		row, err := n.Find(("dallas"))
		if err != nil {
			fmt.Println(err)
		}
		assert.Nil(t, err)
		assert.Equal(t, row, NumbeoDataRow{"Dallas, TX, United States", 77.4, 53.8, 66.2, 75.7, 78.1, 112})
	})

	t.Run("Invalid location", func(t *testing.T) {
		_, err := n.Find(("hello world"))
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "location not found")
	})
}
