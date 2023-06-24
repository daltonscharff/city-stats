package services

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/daltonscharff/city-stats/internal/utils"
	"github.com/go-resty/resty/v2"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

var numbeoColFilename = filepath.Join("..", "..", "testdata", "numbeo_col.html")

func TestNumbeoServiceLocationSearch(t *testing.T) {
	client := resty.New()
	defer gock.Off()
	gock.InterceptClient(client.GetClient())
	gock.New(utils.NumbeoUrl).
		Get("/cost-of-living/rankings_current.jsp").Persist().
		Reply(200).File(numbeoColFilename)

	n := NumbeoService{client}

	t.Run("Valid location, normal case", func(t *testing.T) {
		row, err := n.LocationSearch(("Austin, TX"))
		assert.Nil(t, err)
		assert.Equal(t, NumbeoCostOfLivingRecord{"Austin, TX, United States", 76.2, 66.5, 71.5, 69.9, 85.4, 114}, row)
	})

	t.Run("Valid location, lowercase", func(t *testing.T) {
		row, err := n.LocationSearch(("dallas"))
		if err != nil {
			fmt.Println(err)
		}
		assert.Nil(t, err)
		assert.Equal(t, NumbeoCostOfLivingRecord{"Dallas, TX, United States", 77.4, 53.8, 66.2, 75.7, 78.1, 112}, row)
	})

	t.Run("Invalid location", func(t *testing.T) {
		_, err := n.LocationSearch(("hello world"))
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "location not found")
	})
}
